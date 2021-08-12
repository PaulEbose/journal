/*
Copyright © 2021 Paul Ebose <paulebose@gmail.com>

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

	"github.com/paulebose/journal/internal"
	"github.com/spf13/cobra"
)

// versionRun executes when `version` command is run.
func versionRun(cmd *cobra.Command, args []string) {
	fmt.Printf("Journal v%s \n", internal.Version)
}

// versionCmd represents the `version` command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show installed version of Journal",
	Run:   versionRun,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
