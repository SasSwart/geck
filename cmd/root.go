/*
Copyright © 2021 Sas Swart sasswartcdk@protonmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

type Geck struct {
	rootCmd *cobra.Command
}

func NewGeck() Geck {
	cobra.OnInitialize(initConfig)
	return Geck{
		rootCmd: newRootCommand(),
	}
}

func newRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "geck",
		Short: "An easy to use git based dotfile manager",
		Long: `A configuration management tool for workstations that tracks the files and resources
	that you care about; so that you don't have to write yaml and deploy it before seeing results.`,
	}

	rootCmd.AddCommand(newCloneCommand())
	rootCmd.AddCommand(newAddCommand())
	rootCmd.AddCommand(newRemoveCommand())

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./", "The path to the geck config file")

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (g Geck) Execute() {
	cobra.CheckErr(g.rootCmd.Execute())
}

var cfgFile, path string

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".geck" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".geck")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
