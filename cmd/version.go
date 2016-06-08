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
	"github.com/sapcc/lyra-cli/locales"
	"github.com/sapcc/lyra-cli/version"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: locales.CmdShortDescription("version"),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// DO NOT REMOVE. SHOULD OVERRIDE THE ROOT PersistentPreRunE
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println(version.String())
		return nil
	},
}

func init() {
	RootCmd.AddCommand(VersionCmd)
	initVersionCmdFlags()
}

func initVersionCmdFlags() {
}
