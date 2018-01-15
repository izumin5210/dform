package cmd

import (
	"fmt"
	"runtime"

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
			if app.Config().Debug {
				zapCfg := zap.NewProductionConfig()
				zapCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
				logger, err := zapCfg.Build()
				if err != nil {
					return err
				}
				logger = logger.With(
					zap.String("version", app.Config().Version),
					zap.String("revision", app.Config().Revision),
					zap.String("runtime_version", runtime.Version()),
					zap.String("goos", runtime.GOOS),
					zap.String("goarch", runtime.GOARCH),
				)
				log.SetLogger(logger)
			} else if app.Config().Verbose {
				zapCfg := zap.NewDevelopmentConfig()
				zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
				zapCfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
				logger, err := zapCfg.Build()
				if err != nil {
					return err
				}
				log.SetLogger(logger)
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
