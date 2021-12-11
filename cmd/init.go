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

	git "github.com/libgit2/git2go/v33"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise a local geck repository with no upstream.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		initPath := cmd.Flag("path").Value.String()
		_, err := git.InitRepository(initPath+"/.git", true)
		if err != nil {
			fmt.Println("Could not initialise git repository")
			fmt.Println(err)
		} else {
			fmt.Println("git repository initialised at " + initPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initCmd.Flags().StringP("path", "p", "./", "The local directory in which to initialise a geck repository")
}
