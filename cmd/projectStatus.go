// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var projectStatus = &cobra.Command{
	Use:   "status",
	Short: "Returning the project status",
	Long:  `The status of the project`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := helper.DiscoApiGet(helper.GetProjectAPIURL("/status"))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
