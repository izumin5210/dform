package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/izumin5210/dform/app/component"
	"github.com/izumin5210/dform/util/log"
)

// New creates a new command object
func New(app component.App) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   app.Config().Name,
		Short: "CLI tool to manage Dgraph schema",
		Long:  "CLI tool to manage Dgraph schema",
		PersistentPreRunE: func(*cobra.Command, []string) error {
			// initialize logger
			var lv zap.AtomicLevel
			var logging bool
			if app.Config().Debug {
				lv = zap.NewAtomicLevelAt(zapcore.DebugLevel)
				logging = true
			} else if app.Config().Verbose {
				lv = zap.NewAtomicLevelAt(zapcore.InfoLevel)
				logging = true
			}
			if logging {
				config := zap.Config{
					Level:       lv,
					Development: false,
					Encoding:    "json",
				}
				logger, err := config.Build()
				if err != nil {
					return err
				}
				log.SetLogger(logger.Sugar())
			}
			return nil
		},
		PersistentPostRun: func(*cobra.Command, []string) {
			log.Close()
		},
	}

	cobra.OnInitialize(app.Config().Init)

	rootCmd.PersistentFlags().StringVar(
		&(app.Config().ConfigFilePath),
		"config",
		"",
		fmt.Sprintf("config file (default is $PWD/%s.toml)", app.Config().GetDefaultConfigName()),
	)
	rootCmd.PersistentFlags().BoolVar(
		&(app.Config().Debug),
		"debug",
		false,
		fmt.Sprintf("Debug level output"),
	)
	rootCmd.PersistentFlags().BoolVar(
		&(app.Config().Verbose),
		"verbose",
		false,
		fmt.Sprintf("Verbose level output"),
	)

	rootCmd.AddCommand(
		newDiffCommand(app),
		newExportCommand(app),
		newVersionCommand(app),
	)

	return rootCmd
}
