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

	"github.com/spf13/cobra"
)

func newAddFileCommand() *cobra.Command {
	var path string

	addFileCmd := &cobra.Command{
		Use:   "file",
		Short: "Add a file to the system state watchlist",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Adding " + path + " to the system state watchlist")
		},
	}

	addFileCmd.Flags().StringVar(&path, "path", "", "The path to the file that you would like to track")

	return addFileCmd
}

func newRemoveFileCommand() *cobra.Command {
	var path string

	removeFileCmd := &cobra.Command{
		Use:   "file",
		Short: "Remove a file from the system state watchlist",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Removing " + path + " from the system state watchlist")
		},
	}
	removeFileCmd.Flags().StringVar(&path, "path", "", "The path to the file that you would like to track")

	return removeFileCmd
}
