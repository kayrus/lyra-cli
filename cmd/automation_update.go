// Copyright © 2016 Arturo Reuschenbach <a.reuschenbach.puncernau@sap.com>
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
	"errors"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var AutomationUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates an exsiting automation",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// setup rest client
		err := setupRestClient()
		if err != nil {
			return err
		}

		// setup automation update
		err = setupAutomationUpdate()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	AutomationCmd.AddCommand(AutomationUpdateCmd)
	//flags
	AutomationUpdateCmd.PersistentFlags().StringVarP(&automationId, "id", "", "", "Id from the automation that should be updated.")
}

func setupAutomationUpdate() error {
	// check required automation id
	if len(automationId) == 0 {
		return errors.New("No automation id given.")
	}
	return nil
}