package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	keyDgraphURL = "dgraph-url"
)

type root struct {
	*cobra.Command
	viper                   *viper.Viper
	name, version, revision string
	cfgFile                 string
}

// New creates a new command object
func New(name, version, revision string) *cobra.Command {
	rootCmd := &root{
		Command: &cobra.Command{
			Use:   name,
			Short: "CLI tool to manage Dgraph schema",
			Long:  "CLI tool to manage Dgraph schema",
		},
		viper:    viper.New(),
		name:     name,
		version:  version,
		revision: revision,
	}

	cobra.OnInitialize(rootCmd.initConfig)

	rootCmd.PersistentFlags().StringVar(
		&rootCmd.cfgFile,
		"config",
		"",
		fmt.Sprintf("config file (default is $PWD/%s.toml)", rootCmd.defaultConfigName()),
	)
	rootCmd.AddCommand(
		rootCmd.newExportCommand(),
		rootCmd.newVersionCommand(),
	)

	return rootCmd.Command
}

func (r *root) initConfig() {
	if r.cfgFile != "" {
		r.viper.SetConfigFile(r.cfgFile)
	} else {
		r.viper.AddConfigPath(".")
		r.viper.SetConfigName(r.defaultConfigName())
	}

	r.viper.BindEnv(keyDgraphURL)
	r.viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if err := r.viper.ReadInConfig(); err != nil {
		log.Println(fmt.Errorf("failed to read config: %v", err))
	}
}

func (r *root) defaultConfigName() string {
	return fmt.Sprintf(".%s", r.name)
}

func (r *root) getDgraphURL() string {
	return r.viper.GetString(keyDgraphURL)
}