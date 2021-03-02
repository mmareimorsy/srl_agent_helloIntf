// helloIntf v0.3
// Simple NDK agent to count number of interfaces that are in an admin-up state
// Tested against SRL 20.6.2
// Original Idea from Rob Renisson for a similar example
// Mohamed M. Morsy (mohamed.m.morsy@nokia.com)

package main

import (
	"context"
	"encoding/json"
	pb "helloIntf/nokia.com/srlinux/sdk/protos"
	"io"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"helloIntf/ndkagent"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Global variables used in this version which is not ideal
// Potentially this could go out of sync from the state in IDB
// Ideally initial state should be driven from querying current state in IDB
var interfaces = make(map[string]uint32)
var totalUp int = 0
var agentState bool = false

// struct used to follow the YANG model for ease of JSON Unmarshal
type configModel struct {
	Action string
	Debug  string
}

func handleConfigNotification(notification *pb.Notification, agent *ndkagent.SrlAgent) {
	// handles config notifications
	// notification: This is a notification message as received from the system
	// agent: This is a pointer to the the current agent in use
	configPath := notification.GetConfig().Key.JsPath
	agent.Logger.Debug("Got config ", configPath)
	// Skip commit.end notification for now
	if configPath == ".commit.end" {
		agent.Logger.Debug("Skipping commit.end notification")
		// handle relevant commits that belong to the agent
	} else if strings.Contains(configPath, agent.Name) {
		agent.Logger.Debug("Relevant commit detected, working on it...")
		// Commit should be either Create, Delete or Replace
		operation := notification.GetConfig().Op
		// Json content of the notification
		content := notification.GetConfig().Data.GetJson()
		agent.Logger.Debug("Json data ", content)
		// if commit is asking for delete skip for now
		// Otherwise go ahead & handle the commit
		if operation.String() == "Delete" {
			agent.Logger.Info("Delete detected, To Be Done...")
		} else {
			agent.Logger.Debug("Create commit detected, Need to handle it")
			agent.Logger.Debug(content)
			commitContent := new(configModel)
			json.Unmarshal([]byte(content), commitContent)
			action := commitContent.Action
			verbose := commitContent.Debug
			agent.Logger.Info("action seen is ", action)
			if action == "ACTION_enable" {
				agent.Logger.Info("Turning on agent...")
				agentState = true
				updateCount(agent)
				// if disable then triggers deleting state from IDB
			} else {
				agent.Logger.Info("Turning off agent...")
				jsPath := "." + agent.Name
				ndkagent.DelTelemetry(agent, jsPath)
				agentState = false
			}
			agent.Logger.Info("Value of verbose is ", verbose)
			if verbose == "DEBUG_enable" {
				agent.Logger.SetLevel(log.DebugLevel)
				agent.Logger.Info("Turning on debugging mode...")
			} else {
				agent.Logger.SetLevel(log.InfoLevel)
				agent.Logger.Info("Turning off debugging mode...")
			}
		}
	} else {
		agent.Logger.Debug("Recieved irrelavant config, that shouldn't happen...")
	}
}

func updateCount(agent *ndkagent.SrlAgent) {
	// This builds the state that should be updated
	// calls addTelemetry to actually talk to the IDB
	// agent: This is a pointer to the the current agent in use
	agent.Logger.Info("Updating state of the agent on the system via telemetry")
	jsPath := "." + agent.Name
	stateContent := map[string]map[string]string{}
	stateContent["admin_up_count"] = map[string]string{}
	stateContent["admin_up_count"]["value"] = strconv.Itoa(totalUp)
	jsData, err := json.Marshal(stateContent)
	if err != nil {
		agent.Logger.Debug("Failed in marshalling of JSON data")
	}
	agent.Logger.Debug("JSON Marshal result is ", string(jsData))
	ndkagent.AddTelemetry(agent, jsPath, string(jsData))
}

func handleIntfNotification(notification *pb.Notification, agent *ndkagent.SrlAgent) {
	// handle interface notifications
	// notification: notification message as receieved from the system
	// agent: This is a pointer to the the current agent in use
	interfaceName := notification.GetIntf().Key.IfName
	interfaceAdminState := notification.GetIntf().Data.AdminIsUp
	agent.Logger.Info("Recieved notification for interface ", interfaceName, " with admin state ", interfaceAdminState)
	// check if this interface was already known & with what state
	prevState, found := interfaces[interfaceName]
	if found {
		// no change in state seen
		if prevState == interfaceAdminState {
			agent.Logger.Info("Interface ", interfaceName, " didn't change its state")
			// moved from up to down state so should decrement count
		} else if prevState == 1 && interfaceAdminState == 0 {
			agent.Logger.Info("Interface ", interfaceName, " moved from up to down")
			totalUp--
			interfaces[interfaceName] = interfaceAdminState
			// moved from down to up so should incremenet count
		} else {
			agent.Logger.Info("Interface ", interfaceName, " moved from down to up")
			totalUp++
			interfaces[interfaceName] = interfaceAdminState
		}
	} else {
		// new interface not seen before so update count if it is admin-up
		if interfaceAdminState == 1 {
			agent.Logger.Info("Interface ", interfaceName, " is a new one in up state")
			totalUp++
			interfaces[interfaceName] = interfaceAdminState
		}
	}
	// update telemetry if the agent is configured as enable
	if agentState {
		updateCount(agent)
	}
}
func handleNotification(notification *pb.Notification, agent *ndkagent.SrlAgent) {
	// Handle of config vs intf vs anything else
	// notification: notification message as recieved from the system
	// agent: pointer to the current agent in use
	// Pretify JSON notification for logging only
	prettyNot, err := json.MarshalIndent(notification, "", "    ")
	if err != nil {
		agent.Logger.Debug("Failed to Marshal notification")
	}
	agent.Logger.Debug(string(prettyNot))
	if notification.GetConfig() != nil {
		handleConfigNotification(notification, agent)
	} else if notification.GetIntf() != nil {
		agent.Logger.Debug("Recieved intf notification")
		handleIntfNotification(notification, agent)
	} else {
		agent.Logger.Debug("Skipping other stuff notifications...")
	}
}

func start(agent *ndkagent.SrlAgent, wg *sync.WaitGroup) {
	// Agent initial resgistration & trigger to notification subscription
	// agent: pointer to the current agent in use
	response, err := agent.Stub.AgentRegister(agent.Ctx, &pb.AgentRegistrationRequest{})
	if err != nil {
		agent.Logger.Debug("Failed to send agent resgistration request")
		agent.Logger.Debug(err)
		os.Exit(-1)
	}
	agent.Logger.Info("Agent", agent.Name, "resgistration status is", response.Status)
	if response.Status.String() == "kSdkMgrFailed" {
		agent.Logger.Debug("Failed to register ", agent.Name)
		// os.Exit(-1)
	}
	// Getting app ID
	agent.AppID = response.GetAppId()
	agent.Logger.Info("Agent ", agent.Name, " ID has been set to ", agent.AppID)

	err = ndkagent.SubscribeNotifications(agent)
	if err != nil {
		agent.Logger.Debug("Failed to subscribe to notification stream...")
	}
	defer wg.Done()
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
	// updateCount(agent)
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
	go ndkagent.ExitGracefully(osSignals, agent, wg)

	//setup stdout logging
	ndkagent.LogSetup(agent)

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
	go ndkagent.KeepAlive(agent, wg, 10*time.Second)
	wg.Wait()
	agent.Logger.Info("Terminating all here...")
}
