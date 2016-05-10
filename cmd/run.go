// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
)

type Run struct {
	AutomationId string `json:"automation_id"`
	Selector     string `json:"selector"`
}

var (
	run = Run{}
)

// automationCmd represents the automation command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Automation run service.",
	Long:  `A longer description for automation.`,
}

func init() {
	RootCmd.AddCommand(RunCmd)
}