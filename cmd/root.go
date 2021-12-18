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

// Execute the given Geck command
func Execute() {
	cobra.CheckErr(geckSingleton.rootCmd.Execute())
}

type Geck struct {
	rootCmd *cobra.Command
}

var geckSingleton = newGeck()

// Initialise an instance of the Geck struct
// that contains the root Cobra command of the Geck CLI interface
func newGeck() Geck {
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

	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./", "The path to the geck config file")

	cobra.OnInitialize(initConfig(cfgFile))

	rootCmd.AddCommand(newCloneCommand())
	rootCmd.AddCommand(newAddCommand())
	rootCmd.AddCommand(newRemoveCommand())

	return rootCmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig(cfgFile string) func() {
	return func() {
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
}
