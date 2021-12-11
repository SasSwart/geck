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

	"github.com/sasswart/geck/git"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Initialise a geck repository by cloning an upstream repository",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		repo := cmd.Flag("repo").Value.String()
		path := cmd.Flag("path").Value.String()
		branch := cmd.Flag("branch").Value.String()

		_, err := git.Clone(repo, path, branch)

		if err != nil {
			fmt.Println("Could not clone git repository")
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)

	cloneCmd.Flags().StringP("repo", "r", "", "The remote git repository to clone")
	cloneCmd.Flags().StringP("path", "p", "./", "The local path to which to clone the remote git repository")
	cloneCmd.Flags().StringP("branch", "b", "main", "Which branch of the remote repository to clone")
}
