package ndkagent

import "encoding/json"

// config manipulation based on yang model
type Config struct {
	TopContainer *HelloIntfConfig
}

type HelloIntfConfig struct {
	Action *string `json:"action"`
	Debug  *string `json:"debug"`
}

// InitConfig initializes a config model struct
func InitConfig() *Config {
	newConfig := new(Config)
	newTopContainer := new(HelloIntfConfig)
	newConfig.TopContainer = newTopContainer
	return newConfig
}

// PopulateConfig is basically unmarshal from json into config struct
func (c *Config) PopulateConfig(json_config string) error {
	err := c.TopContainer.PopulateConfig(json_config)
	if err != nil {
		return err
	}
	return nil
}

// PopulateConfig populates only the top container level of the config
func (c *HelloIntfConfig) PopulateConfig(json_config string) error {
	err := json.Unmarshal([]byte(json_config), c)
	return err
}

// GetConfig returns the current config as it has been populated so far
func (c *HelloIntfConfig) GetConfig() *HelloIntfConfig {
	return c
}

// GetAction returns the configured action under HelloIntf module
func (c *HelloIntfConfig) GetAction() uint32 {
	if *c.Action == "ACTION_enable" {
		return 0
	} else {
		return 1
	}
}

// GetDebug returns the configured Debug under HelloIntf module
func (c *HelloIntfConfig) GetDebug() uint32 {
	if *c.Debug == "DEBUG_enable" {
		return 0
	} else {
		return 1
	}
}
