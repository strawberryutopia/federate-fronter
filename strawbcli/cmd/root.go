package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool
var jsonLogs bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "strawbcli",
	Short: "Demo app to test federate-fronter",
	Long:  `A simple CLI which I can use to test the federate-fronter package`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		if jsonLogs {
			log.SetFormatter(&log.JSONFormatter{})
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.strawbcli.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
	rootCmd.PersistentFlags().BoolVarP(&jsonLogs, "json", "j", false, "Log in json format")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".strawbcli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".strawbcli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// TODO: bind in config from PWD, a like from
	// https://github.com/spf13/viper/issues/181#issuecomment-678528505
}
