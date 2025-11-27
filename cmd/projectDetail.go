// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var projectDetailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Returning the project details",
	Long:  `The details of the project`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := helper.DiscoApiGet(helper.GetProjectAPIURL(conf.DefaultApiVersion, ""))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
