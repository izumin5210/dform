package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultConfigName = ".dform"
)

var (
	cfgFile = ""
	rootCmd = &cobra.Command{
		Use:   "dform",
		Short: "CLI tool to manage Dgraph schema",
		Long:  "CLI tool to manage Dgraph schema",
	}
)

// Execute executes the command.
func Execute() int {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return -1
	}
	return 0
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default is $PWD/%s.toml)", defaultConfigName))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(defaultConfigName)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("failed to read config: %v", err))
	}
}
