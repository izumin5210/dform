package system

import (
	"fmt"
	"io"
	"strings"

	"github.com/spf13/viper"

	"github.com/izumin5210/dform/util/log"
)

const (
	// viper keys
	keyDgraphURL  = "dgraph-url"
	keySchemaPath = "dgraph-schema-path"

	// default values
	defaultSchemaName = "dgraph.schema"
)

// Config stores general setting params and provides accessors for them.
type Config struct {
	Name, Version, Revision string
	ConfigFilePath          string
	SchemaFilePath          string
	Debug, Verbose          bool
	InReader                io.Reader
	OutWriter, ErrWriter    io.Writer
	viper                   *viper.Viper
}

// NewConfig creates new Config object.
func NewConfig(
	name, version, revision string,
	inReader io.Reader,
	outWriter, errWriter io.Writer,
) *Config {
	return &Config{
		Name:      name,
		Version:   version,
		Revision:  revision,
		InReader:  inReader,
		OutWriter: outWriter,
		ErrWriter: errWriter,
		viper:     viper.New(),
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
		log.Error("failed to read config", "error", err)
	}
}

// GetDefaultConfigName returns a default config file name
func (c *Config) GetDefaultConfigName() string {
	return fmt.Sprintf(".%s", c.Name)
}

// GetDefaultSchemaName returns a default schema file path
func (c *Config) GetDefaultSchemaName() string {
	return defaultSchemaName
}

// GetDgraphURL returns the URL for the target Dgraph.
func (c *Config) GetDgraphURL() string {
	return c.viper.GetString(keyDgraphURL)
}

// GetSchemaPath returns the path for the schema file.
func (c *Config) GetSchemaPath() string {
	if len(c.ConfigFilePath) > 0 {
		return c.ConfigFilePath
	} else if c.viper.IsSet(keySchemaPath) {
		return c.viper.GetString(keySchemaPath)
	}
	return c.GetDefaultSchemaName()
}
