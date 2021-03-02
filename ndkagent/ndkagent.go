package ndkagent

import (
	"context"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"

	pb "helloIntf/nokia.com/srlinux/sdk/protos"
)

// SrlAgent struct used to ease moving around the agent needed components around different functions in the code
type SrlAgent struct {
	AppID            uint32
	Name             string
	Timeout          time.Duration
	Channel          *grpc.ClientConn
	Stub             pb.SdkMgrServiceClient
	Ctx              context.Context
	Cancel           context.CancelFunc
	Logger           *log.Logger
	NotificationStub pb.SdkNotificationServiceClient
	StreamID         uint64
	StreamClient     pb.SdkNotificationService_NotificationStreamClient
	TelemetryStub    pb.SdkMgrTelemetryServiceClient
}

// AddTelemetry will add or update state of agent on IDB
// agent: This is a pointer to the the current agent in use
// jsPath: string for the root key of the agent, usually that's the .agent_name
// jsData: properly formatted JSON data to populate the updated state
func AddTelemetry(agent *SrlAgent, jsPath string, jsData string) {
	telemetryKey := &pb.TelemetryKey{
		JsPath: jsPath,
	}
	telemetryData := &pb.TelemetryData{
		JsonContent: jsData,
	}
	telemetryInfo := &pb.TelemetryInfo{
		Key:  telemetryKey,
		Data: telemetryData,
	}
	updateRequest := &pb.TelemetryUpdateRequest{
		State: []*pb.TelemetryInfo{telemetryInfo},
	}
	agent.Logger.Debug("Actual telemetry update is ", updateRequest)
	response, err := agent.TelemetryStub.TelemetryAddOrUpdate(agent.Ctx, updateRequest)
	if err != nil {
		agent.Logger.Debug("Failted to add Telemetry")
	}
	agent.Logger.Debug("Response for adding telemetry ", response.Status)
}

// DelTelemetry will delete the current state of the agent, used when we disable the agent in CLI
// agent: This is a pointer to the the current agent in use
// jsPath: string for the root key of the agent, usually that's the .agent_name
func DelTelemetry(agent *SrlAgent, jsPath string) {
	telemetryKey := &pb.TelemetryKey{
		JsPath: jsPath,
	}
	deleteRequest := &pb.TelemetryDeleteRequest{
		Key: []*pb.TelemetryKey{telemetryKey},
	}
	response, err := agent.TelemetryStub.TelemetryDelete(agent.Ctx, deleteRequest)
	if err != nil {
		agent.Logger.Debug("Failted to delete Telemetry")
	}
	agent.Logger.Debug("Response for deleting telemetry ", response.Status)
}

// Subscribe subscribes to notification topics as needed
// agent: pointer to the current agent in use
// topic: name of the topic to subscribe to, has to be one of the supported topics on the system
func Subscribe(agent *SrlAgent, topic string) {
	agent.Logger.Debug("Subscribing for ", topic)
	op := pb.NotificationRegisterRequest_AddSubscription
	response := &pb.NotificationRegisterResponse{}
	var err error
	if topic == "intf" {
		response, err = agent.Stub.NotificationRegister(agent.Ctx, &pb.NotificationRegisterRequest{
			StreamId:          agent.StreamID,
			Op:                op,
			SubscriptionTypes: &pb.NotificationRegisterRequest_Intf{},
		})
	}
	if topic == "nw_inst" {
		response, err = agent.Stub.NotificationRegister(agent.Ctx, &pb.NotificationRegisterRequest{
			StreamId:          agent.StreamID,
			Op:                op,
			SubscriptionTypes: &pb.NotificationRegisterRequest_NwInst{},
		})
	}
	if topic == "lldp" {
		response, err = agent.Stub.NotificationRegister(agent.Ctx, &pb.NotificationRegisterRequest{
			StreamId:          agent.StreamID,
			Op:                op,
			SubscriptionTypes: &pb.NotificationRegisterRequest_LldpNeighbor{},
		})
	}
	if topic == "route" {
		response, err = agent.Stub.NotificationRegister(agent.Ctx, &pb.NotificationRegisterRequest{
			StreamId:          agent.StreamID,
			Op:                op,
			SubscriptionTypes: &pb.NotificationRegisterRequest_Route{},
		})
	}
	if topic == "cfg" {
		response, err = agent.Stub.NotificationRegister(agent.Ctx, &pb.NotificationRegisterRequest{
			StreamId:          agent.StreamID,
			Op:                op,
			SubscriptionTypes: &pb.NotificationRegisterRequest_Config{},
		})
	}

	if err != nil {
		agent.Logger.Debug("Failed to subscribe for ", topic)
	}
	agent.Logger.Debug("Response for Notification register for ", topic, " is ", response.Status)
}

// SubscribeNotifications create notifications stub & register for notifications
// agent: pointer to the current agent in use
func SubscribeNotifications(agent *SrlAgent) error {
	notificationResponse, err := agent.Stub.NotificationRegister(agent.Ctx, &pb.NotificationRegisterRequest{})
	if err != nil {
		agent.Logger.Debug("Failed to resgister for notifications...")
		// os.Exit(-1)
	}
	agent.Logger.Debug("Response for notifications register ", notificationResponse.Status)
	// get stream ID
	agent.StreamID = notificationResponse.StreamId
	agent.Logger.Debug("Stream ID for notifications is ", agent.StreamID)

	// do the actual subscribtion here
	Subscribe(agent, "intf")
	Subscribe(agent, "nw_inst")
	Subscribe(agent, "lldp")
	Subscribe(agent, "route")
	Subscribe(agent, "cfg")

	// get a stream client for streaming notifications
	agent.StreamClient, err = agent.NotificationStub.NotificationStream(agent.Ctx, &pb.NotificationStreamRequest{
		StreamId: agent.StreamID,
	})
	if err != nil {
		agent.Logger.Debug("Failed to subscribe to notification stream...")
	}
	agent.Logger.Debug("agent has created a stream client")

	return err
}

// ExitGracefully go routine to handle kill signal
// osSignals: a channel of data type os.Signal
// agent: pointer to the current agent in use
func ExitGracefully(osSignals chan os.Signal, agent *SrlAgent, wg *sync.WaitGroup) {
	recievedSignal := <-osSignals
	agent.Logger.Info("Kill signal detected ", recievedSignal.String())
	agent.Ctx.Done()
	response, err := agent.Stub.AgentUnRegister(agent.Ctx, &pb.AgentRegistrationRequest{})
	if err != nil {
		agent.Logger.Info("Exiting immediatley, no cleanup done...")
		os.Exit(-1)
	}
	agent.Logger.Info("Recieved response for un-register request ", response.Status)
	wg.Done()
}

// KeepAlive used to send keepalive requests for agent liveliness
// In case of sdk_mgr restarting the keepalives should detect that & restart the agent after 3 missed keepalives
// agent: instance of SrlAgent
// wg: waitGroup for the thread running the keepalives
// keepAliveInterval: time.Duration for how frequent the keepalives should be (multiplier of 3 is hardcoded)
func KeepAlive(agent *SrlAgent, wg *sync.WaitGroup, keepAliveInterval time.Duration) {
	defer wg.Done()
	missedCounter := 0
	for {
		time.Sleep(keepAliveInterval)
		keepAliveRequest := pb.KeepAliveRequest{}
		response, err := agent.Stub.KeepAlive(agent.Ctx, &keepAliveRequest)
		if err != nil {
			agent.Logger.Debug("Failed to send keep alive request")
		}
		agent.Logger.Debug("Keepalive sent at ", time.Now(), " got ", response.Status)
		if response.Status.String() != "kSdkMgrSuccess" {
			agent.Logger.Info("Keepalive missed")
			missedCounter++
			if missedCounter >= 3 {
				agent.Logger.Info("3 keepalives missed in a row")
				wg.Done()
				break
			}
		} else {
			missedCounter = 0
		}
	}
}

// LogSetup setting up of logging to stdout directory
func LogSetup(agent *SrlAgent) {
	// Handle stdout for logging
	hostName, _ := os.Hostname()
	stdoutDir := "/var/log/srlinux/stdout"
	logFileName := stdoutDir + "/" + hostName + agent.Name + ".log"
	_, err := os.Stat(stdoutDir)
	if os.IsNotExist(err) {
		os.MkdirAll(stdoutDir, 0760)
	}
	logWriter, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// defer logWriter.Close()
	// agent.Logger = log.New(logWriter, header, log.LstdFlags)
	agent.Logger = log.New()
	agent.Logger.SetOutput(logWriter)
	agent.Logger.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	agent.Logger.SetLevel(log.InfoLevel)
	agent.Logger.Info("Starting agent ", agent.Name, "...")
}
