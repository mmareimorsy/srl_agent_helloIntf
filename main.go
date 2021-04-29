// helloIntf v0.3
// Simple NDK agent to count number of interfaces that are in an admin-up state
// Tested against SRL 21.3
// Original Idea from Rob Renisson for a similar example
// Mohamed M. Morsy

package main

import (
	"context"
	"encoding/json"
	pb "helloIntf/nokia.com/srlinux/sdk/protos"
	"io"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"helloIntf/ndkagent"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// handleConfigNotification handles a config notification & update state as well for configured items
func handleConfigNotification(notification *pb.Notification, agent *ndkagent.SrlAgent) {
	configPath := notification.GetConfig().Key.JsPath
	agent.Logger.Debug("Got config ", configPath)
	if configPath == ".commit.end" { // Skip commit.end notification for now
		agent.Logger.Debug("Skipping commit.end notification")
	} else if strings.Contains(configPath, agent.Name) { // handle relevant commits that belong to the agent
		agent.Logger.Debug("Relevant commit detected, working on it...")
		operation := notification.GetConfig().Op           // Commit should be either Create, Delete or Replace
		content := notification.GetConfig().Data.GetJson() // Json content of the notification
		agent.Logger.Debug("Json data ", content)
		// if commit is asking for delete skip for now
		if operation.String() == "Delete" {
			agent.Logger.Info("Delete detected, To Be Done...")
		} else {
			agent.Logger.Debug("Create commit detected, Need to handle it")
			agent.Logger.Info(content)
			err := agent.YangModel.AgentConfig.PopulateConfig(content)
			if err != nil {
				agent.Logger.Info("Failed to populate new config", err)
			}
			runningConfig := agent.YangModel.AgentConfig.TopContainer.GetConfig()
			action := agent.YangModel.AgentConfig.TopContainer.GetAction()
			debug := agent.YangModel.AgentConfig.TopContainer.GetDebug()
			agent.Logger.Info("action seen is ", action)
			agent.YangModel.AgentState.TopContainer.UpdateAction(runningConfig)
			if agent.YangModel.AgentConfig.TopContainer.GetAction() == 0 {
				agent.Logger.Info("Turning on agent...")
			} else {
				agent.Logger.Info("Turning off agent...")
				jsPath := "." + agent.Name
				agent.DelTelemetry(jsPath)
			}
			agent.Logger.Info("Value of verbose is ", debug)
			agent.YangModel.AgentState.TopContainer.UpdateDebug(runningConfig)
			if debug == 0 {
				agent.Logger.SetLevel(log.DebugLevel)
				agent.Logger.Info("Turning on debugging mode...")
			} else {
				agent.Logger.SetLevel(log.InfoLevel)
				agent.Logger.Info("Turning off debugging mode...")
			}
			agent.YangModel.AgentState.TopContainer.PopulateState(agent)
		}
	} else {
		agent.Logger.Debug("Recieved irrelavant config, that shouldn't happen...")
	}
}

// handleIntfNotification handles state change of interfaces as they are streamed & update the interfacesList inside interfaces container
func handleIntfNotification(notification *pb.Notification, agent *ndkagent.SrlAgent) {
	interfaceName := notification.GetIntf().Key.IfName
	interfaceAdminState := notification.GetIntf().Data.AdminIsUp
	agent.Logger.Info("Recieved notification for interface ", interfaceName, " with admin state ", interfaceAdminState)
	// check if this interface was already known & with what state
	prevState, found := agent.YangModel.AgentState.InterfacesContainer.InterfacesList[interfaceName]
	if found {
		// no change in state seen
		if prevState.Interface.GetState() == interfaceAdminState {
			agent.Logger.Info("Interface ", interfaceName, " didn't change its state")
			// moved from up to down state so should decrement count
		} else if prevState.Interface.GetState() == 1 && interfaceAdminState == 0 {
			agent.Logger.Info("Interface ", interfaceName, " moved from up to down")
			prevState.Interface.SetState(interfaceAdminState)
			agent.YangModel.AgentState.TopContainer.DecrementIntfCounter()
			prevState.PopulateState(interfaceName, agent)
			// moved from down to up so should incremenet count
		} else {
			agent.Logger.Info("Interface ", interfaceName, " moved from down to up")
			prevState.Interface.SetState(interfaceAdminState)
			agent.YangModel.AgentState.TopContainer.IncrementIntfCounter()
			prevState.PopulateState(interfaceName, agent)
		}
	} else {
		// new interface not seen before so update count if it is admin-up
		if interfaceAdminState == 1 {
			agent.Logger.Info("Interface ", interfaceName, " is a new one in up state")
			newIntfElement := new(ndkagent.InterfaceElement)
			newIntfElement.Interface = new(ndkagent.InterfaceState)
			newIntfElement.Interface.SetState(interfaceAdminState)
			agent.YangModel.AgentState.InterfacesContainer.InterfacesList[interfaceName] = newIntfElement
			newIntfElement.PopulateState(interfaceName, agent)
			agent.YangModel.AgentState.TopContainer.IncrementIntfCounter()
		}
	}
	agent.YangModel.AgentState.TopContainer.PopulateState(agent)
}

// handleNotification handles streamed notifications based on their type
func handleNotification(notification *pb.Notification, agent *ndkagent.SrlAgent) {
	prettyNot, err := json.MarshalIndent(notification, "", "    ")
	if err != nil {
		agent.Logger.Debug("Failed to Marshal notification")
	}
	agent.Logger.Debug(string(prettyNot))
	agent.Logger.Debug(notification)
	if notification.GetConfig() != nil {
		handleConfigNotification(notification, agent)
	} else if notification.GetIntf() != nil {
		agent.Logger.Debug("Recieved intf notification")
		handleIntfNotification(notification, agent)
	} else {
		agent.Logger.Debug("Skipping other stuff notifications...")
	}
}

// start is where the agent registers with the system, subscribe to notifications & initializes config & state structs of the agent
func start(agent *ndkagent.SrlAgent, wg *sync.WaitGroup) {
	response, err := agent.Stub.AgentRegister(agent.Ctx, &pb.AgentRegistrationRequest{})
	if err != nil {
		agent.Logger.Debug("Failed to send agent resgistration request")
		agent.Logger.Debug(err)
		os.Exit(-1)
	}
	agent.Logger.Info("Agent", agent.Name, "resgistration status is", response.Status)
	if response.Status.String() == "kSdkMgrFailed" {
		agent.Logger.Debug("Failed to register ", agent.Name)
	}
	agent.AppID = response.GetAppId() // Getting app ID
	agent.Logger.Info("Agent ", agent.Name, " ID has been set to ", agent.AppID)

	err = agent.SubscribeNotifications()
	if err != nil {
		agent.Logger.Debug("Failed to subscribe to notification stream...")
	}
	agent.Logger.Info("Subscribed to desired notifications")
	defer wg.Done()

	// Init config & state components
	agent.YangModel.AgentConfig = ndkagent.InitConfig()
	agent.YangModel.AgentState = ndkagent.InitState()

	for {
		streamResponse, err := agent.StreamClient.Recv()
		if err == io.EOF {
			agent.Logger.Debug("Server declares no more notifications to be streamed")
			return
		}
		if err != nil {
			agent.Logger.Debug("Failed to recieve stream response... ", err)
			return
		}
		notifications := streamResponse.GetNotification()
		agent.Logger.Debug("Received ", len(notifications), " new notifications")
		for _, notification := range notifications {
			handleNotification(notification, agent)
		}
	}
}

func main() {
	// Create a new instance of srlAgent
	agent := new(ndkagent.SrlAgent)
	agent.Name = "helloIntf"

	// Handling of Kill signal from OS via routine
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGTERM, syscall.SIGQUIT)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go agent.ExitGracefully(osSignals, wg)

	//setup stdout logging
	agent.LogSetup()

	// Setting up grpc channel
	agent.Logger.Debug("Establishing grpc channel...")
	var err error
	agent.Channel, err = grpc.Dial("127.0.0.1:50053", grpc.WithInsecure())
	if err != nil {
		agent.Logger.Info("Failed to establish grpc channel")
	}
	defer agent.Channel.Close()

	// Creating the stubs
	agent.Stub = pb.NewSdkMgrServiceClient(agent.Channel)
	agent.Timeout = 5
	agent.Ctx, agent.Cancel = context.WithCancel(context.Background())
	defer agent.Cancel()
	agent.Ctx = metadata.AppendToOutgoingContext(agent.Ctx, "agent_name", agent.Name)
	agent.NotificationStub = pb.NewSdkNotificationServiceClient(agent.Channel)
	agent.TelemetryStub = pb.NewSdkMgrTelemetryServiceClient(agent.Channel)

	wg.Add(1)
	go start(agent, wg)
	wg.Add(1)
	go agent.KeepAlive(wg, 10*time.Second)
	wg.Wait()
	agent.Logger.Info("Terminating all here...")
}
