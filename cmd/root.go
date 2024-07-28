package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/gkwa/myaction/internal/logger"
)

var (
	cfgFile      string
	verbose      bool
	logFormat    string
	cliLogger    logr.Logger
	gmailctlPath string
)

var rootCmd = &cobra.Command{
	Use:   "myaction",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cliLogger.IsZero() {
			cliLogger = logger.NewConsoleLogger(verbose, logFormat == "json")
		}

		ctx := logr.NewContext(context.Background(), cliLogger)
		cmd.SetContext(ctx)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myaction.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "", "json or text (default is text)")
	rootCmd.PersistentFlags().StringVar(&gmailctlPath, "gmailctl-path", "/usr/local/bin/gmailctl", "path to gmailctl executable")

	if err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		fmt.Printf("Error binding verbose flag: %v\n", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("log-format", rootCmd.PersistentFlags().Lookup("log-format")); err != nil {
		fmt.Printf("Error binding log-format flag: %v\n", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("gmailctl-path", rootCmd.PersistentFlags().Lookup("gmailctl-path")); err != nil {
		fmt.Printf("Error binding gmailctl-path flag: %v\n", err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".myaction")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	logFormat = viper.GetString("log-format")
	verbose = viper.GetBool("verbose")
	gmailctlPath = viper.GetString("gmailctl-path")
}

func LoggerFrom(ctx context.Context, keysAndValues ...interface{}) logr.Logger {
	if cliLogger.IsZero() {
		cliLogger = logger.NewConsoleLogger(verbose, logFormat == "json")
	}
	newLogger := cliLogger
	if ctx != nil {
		if l, err := logr.FromContext(ctx); err == nil {
			newLogger = l
		}
	}
	return newLogger.WithValues(keysAndValues...)
}
