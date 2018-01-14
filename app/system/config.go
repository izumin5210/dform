package system

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

const (
	keyDgraphURL = "dgraph-url"
)

// Config stores general setting params and provides accessors for them.
type Config struct {
	Name, Version, Revision string
	ConfigFilePath          string
	viper                   *viper.Viper
}

// NewConfig creates new Config object.
func NewConfig(name, version, revision string) *Config {
	return &Config{
		Name:     name,
		Version:  version,
		Revision: revision,
		viper:    viper.New(),
	}
}

// Init reads settings from files and environment variables.
func (c *Config) Init() {
	if c.ConfigFilePath != "" {
		c.viper.SetConfigFile(c.ConfigFilePath)
	} else {
		c.viper.AddConfigPath(".")
		c.viper.SetConfigName(c.GetDefaultConfigName())
	}

	c.viper.BindEnv(keyDgraphURL)
	c.viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if err := c.viper.ReadInConfig(); err != nil {
		log.Println(fmt.Errorf("failed to read config: %v", err))
	}
}

// GetDefaultConfigName returns a default config file name
func (c *Config) GetDefaultConfigName() string {
	return fmt.Sprintf(".%s", c.Name)
}

// GetDgraphURL returns the URL for the target Dgraph.
func (c *Config) GetDgraphURL() string {
	return c.viper.GetString(keyDgraphURL)
}
