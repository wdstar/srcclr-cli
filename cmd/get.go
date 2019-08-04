/*
Copyright © 2019 wdstar

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
	"log"
	"time"

	"github.com/spf13/cobra"
	"github.com/wdstar/srcclr-cli/client/workspaces"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get (workspaces|ws|issues)",
	//DisableFlagsInUseLine: true,
	Short: "Get entities (workspaces, issues)",
	Long: `Get command retrieves entity info. via SourceClear API.
Available entity types are workspaces, ws (alias of workspaces) and issues.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entity := args[0]
		switch entity {
		case "workspaces", "ws":
			params := workspaces.NewGetWorkspacesUsingGETParams()
			params.SetTimeout(10 * time.Second)
			// In the case the swagger specs has security definitions
			resp, err := srcclrClient.Workspaces.GetWorkspacesUsingGET(params, hmacAuth)
			// In the case the swagger specs has NO security definitions
			// See the `init()` function's setup of the `cmd/root.go` too.
			//resp, err := srcclrClient.Workspaces.GetWorkspacesUsingGET(params)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%#v\n", resp.Payload)
		default:
			log.Fatalf("invalid entity type: %s", entity)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
