package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/izumin5210/dform/app/di"
	"github.com/izumin5210/dform/util/log"
)

// New creates a new command object
func New(component di.RootComponent) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           component.Config().Name,
		Short:         "CLI tool to manage Dgraph schema",
		Long:          "CLI tool to manage Dgraph schema",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(*cobra.Command, []string) error {
			// initialize logger
			var zapCfg zap.Config
			if component.Config().Debug {
				zapCfg = zap.NewProductionConfig()
				zapCfg.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
				zapCfg.InitialFields = map[string]interface{}{
					"version":         component.Config().Version,
					"revision":        component.Config().Revision,
					"runtime_version": runtime.Version(),
					"goos":            runtime.GOOS,
					"goarch":          runtime.GOARCH,
				}
			} else if component.Config().Verbose {
				zapCfg = zap.NewDevelopmentConfig()
				zapCfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
				zapCfg.DisableStacktrace = true
				zapCfg.DisableCaller = true
				zapCfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
					enc.AppendString(t.Local().Format("2006-01-02 15:04:05 MST"))
				}
				zapCfg.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
			}
			if len(zapCfg.Encoding) != 0 {
				logger, err := zapCfg.Build()
				if err != nil {
					return err
				}
				log.SetLogger(logger)
			}
			return nil
		},
	}

	cobra.OnInitialize(component.Config().Init)

	rootCmd.PersistentFlags().StringVar(
		&(component.Config().ConfigFilePath),
		"config",
		"",
		fmt.Sprintf("config file (default is $PWD/%s.toml)", component.Config().GetDefaultConfigName()),
	)
	rootCmd.PersistentFlags().BoolVar(
		&(component.Config().Debug),
		"debug",
		false,
		fmt.Sprintf("Debug level output"),
	)
	rootCmd.PersistentFlags().BoolVarP(
		&(component.Config().Verbose),
		"verbose",
		"v",
		false,
		fmt.Sprintf("Verbose level output"),
	)

	rootCmd.AddCommand(
		newDiffCommand(component),
		newExportCommand(component),
		newVersionCommand(component),
	)

	return rootCmd
}
