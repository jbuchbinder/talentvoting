package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var (
	Config           *AppConfig
	cachedConfigPath string
)

type AppConfig struct {
	Debug    bool `yaml:"debug"`
	Port     int  `yaml:"port"`
	Database struct {
		Name string `yaml:"name"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
		Host string `yaml:"host"`
	} `yaml:"database"`
}

func (c *AppConfig) SetDefaults() {
	c.Port = 8888
	c.Database.Name = "talent"
	c.Database.User = "talent"
	c.Database.Pass = "talent"
	c.Database.Host = "localhost"
}

func LoadConfigWithDefaults(configPath string) (*AppConfig, error) {
	cachedConfigPath = configPath
	c := &AppConfig{}
	c.SetDefaults()
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal([]byte(data), c)

	return c, err
}

func ConfigReload() error {
	c := &AppConfig{}
	c.SetDefaults()
	data, err := ioutil.ReadFile(cachedConfigPath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal([]byte(data), c)
	if err == nil {
		Config = c
	}
	return err
}
