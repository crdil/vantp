package cmd

import (
	"fmt"
	"os"
	"projects/cobratest/globals"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigFile to use for request
var ConfigFile string

// Verbose response
var Verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vantp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	if err := rootCmd.Execute(); err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	fmt.Println(`Should make a GET request`)
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// TODO: add fix for default command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// TODO: working on contenttype, and form requst

func init() {
	cobra.OnInitialize(initConfig)
	globals.ContentType = "json"
	rootCmd.PersistentFlags().StringVar(&ConfigFile, "config", "", "config file (default is $HOME/.cobratest.yaml)")
	rootCmd.PersistentFlags().StringVar(&globals.ContentType, "type", "json", "Set content type for request and response, values 'json', 'form', 'text', 'html TODO' and 'file TODO' is supported")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringArrayVar(&globals.Headers, "header", []string{}, "Headers to use for request, usage: Token:123")
	rootCmd.PersistentFlags().StringVar(&globals.BasicAuth, "auth", "", "Basic authentication with request")
	rootCmd.PersistentFlags().IntVarP(&globals.Timeout, "timeout", "t", 2000, "Set the request timeout")
	rootCmd.PersistentFlags().BoolVar(&globals.NoSSLVerify, "no-verify", false, "Disable SSL verification")
	// rootCmd.PersistentFlags().StringVarP(&globals.Session, "session", "s", "", "Session to use during the request")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(ConfigFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobratest" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobratest")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
