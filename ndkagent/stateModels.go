package ndkagent

import "encoding/json"

// state manipulation based on yang model
type State struct {
	TopContainer        *HelloIntfState
	InterfacesContainer *InterfacesContainerState
}
type HelloIntfState struct {
	Debug        uint32 `json:"debug"`
	Action       uint32 `json:"action"`
	AdminUpCount uint32 `json:"admin-up-count"`
}
type InterfacesContainerState struct {
	InterfacesList map[string]*InterfaceElement
}
type InterfaceState struct {
	State string `json:"state"`
}
type InterfaceElement struct {
	Interface *InterfaceState `json:"interface"`
}

// InitState initializes state struct based on yang model
func InitState() *State {
	newState := new(State)
	newHelloIntfState := new(HelloIntfState)
	newState.TopContainer = newHelloIntfState
	newInterfacesContainer := new(InterfacesContainerState)
	newInterfacesContainer.InterfacesList = make(map[string]*InterfaceElement)
	newState.InterfacesContainer = newInterfacesContainer
	return newState
}

// IncrementIntfCounter increments the admin-up-count in state
func (s *HelloIntfState) IncrementIntfCounter() {
	s.AdminUpCount += 1
}

// DecrementIntfCounter decrements the admin-up-count in state
func (s *HelloIntfState) DecrementIntfCounter() {
	s.AdminUpCount -= 1
}

// UpdateAction reflects to state the Action of the agent as configured
func (s *HelloIntfState) UpdateAction(c *HelloIntfConfig) {
	s.Action = c.GetAction()
}

// UpdateDebug reflects to state the Debug flag as configured
func (s *HelloIntfState) UpdateDebug(c *HelloIntfConfig) {
	s.Debug = c.GetDebug()
}

// PopulateState populates the state of the top container
func (s *HelloIntfState) PopulateState(agent *SrlAgent) {
	agent.Logger.Info("Updating state of the agent on the system via telemetry new method")
	jsPath := "." + agent.Name
	jsData, err := json.Marshal(s)
	if err != nil {
		agent.Logger.Debug("Failed in marshalling of JSON data")
	}
	agent.Logger.Debug("JSON Marshal result is ", string(jsData))
	agent.AddTelemetry(jsPath, string(jsData))
}

// PopulateState populates the state of a key in interfaces list in interfaces container, it creates key if not existing or updates it if existing already
func (s *InterfaceElement) PopulateState(key string, agent *SrlAgent) {
	newKey := "{.name==\"" + key + "\"}"
	jsPath := "." + agent.Name + ".interfaces.interface" + newKey
	jsData, err := json.Marshal(s)
	if err != nil {
		agent.Logger.Debug("Failed in marshalling of JSON data")
	}
	agent.Logger.Debug("JSON Marshal result is ", string(jsData))
	agent.AddTelemetry(jsPath, string(jsData))
}

// GetState returns the state of a specific interface
func (i *InterfaceState) GetState() uint32 {
	if i.State == "admin-up" {
		return 1
	} else {
		return 0
	}
}

// SetState sets the state of a specific interface
func (i *InterfaceState) SetState(adminstate uint32) {
	if adminstate == 1 {
		i.State = "admin-up"
	} else {
		i.State = "admin-down"
	}
}
