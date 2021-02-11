package model

type TestConfiguration struct {
	Timeout            	uint   `yaml:"timeout" json:"timeout" default:"8080"`
	AppName        		string `yaml:"appname" json:"appname"`
}

func NewTest() *TestConfiguration  {
	return &TestConfiguration{}
}

func (properties *TestConfiguration) GetConfigurationPrefix() string {
	return "myconfig"
}