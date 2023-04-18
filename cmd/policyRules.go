// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var policyRulesCmd = &cobra.Command{
	Use:   "policyrules",
	Short: "Returning the project policy rules",
	Long:  `The policy rules of the project`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := helper.DiscoApiGet(helper.GetProjectAPIURL("/policyrules"))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
