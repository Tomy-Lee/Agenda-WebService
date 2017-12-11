// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"entity"
	"os"
	"utils"

	errors "github.com/go-errors/errors"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	serverURL string
	logger    = utils.NewLogger()
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "Agenda-Go",
	Short: "Agenda implemented in golang",
	Long:  `Agenda implemented in golang`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer utils.CloseLogFile()
	if _, present := os.LookupEnv("DEBUG"); present == false {
		defer func() {
			if err := recover(); err != nil {
				logger.Fatalln("[execute]", err)
			}
			return
		}()
	}
	if err := RootCmd.Execute(); err != nil {
		return
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.agenda-go.yaml)")

	RootCmd.PersistentFlags().StringVar(&serverURL, "server", "", "set server url (default to http://localhost:8080)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.Fatalln("Failed to find home directory", errors.Wrap(err, 0))
		}

		// Search config in home directory with name ".agenda-go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".agenda-go")
	}

	if _, present := os.LookupEnv("MOCK"); present {
		serverURL = "https://private-ef9a4e-mensu.apiary-mock.com"
	}

	if serverURL != "" {
		if err := entity.SetServer(serverURL); err != nil {
			logger.Fatalln("Invalid server url", serverURL)
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()

	// set current working directory
	cwd := viper.GetString("cwd")
	if len(cwd) > 0 {
		err := os.Chdir(cwd)
		if err != nil {
			logger.Fatalf("Failed to change current working directory: %v", err)
		}
	}

	// set log file
	logPath := viper.GetString("log")
	utils.InitLogFile(logPath)

	if err == nil {
		logger.Println("[init] using configs from config file:", viper.ConfigFileUsed())
	} else {
		logger.Println("[init] using default configs")
	}

	entity.Init()
}
