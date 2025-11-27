// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var versionListCmd = &cobra.Command{
	Use:   "list",
	Short: "Returning the project version list",
	Long:  `The version list of the project`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := helper.DiscoApiGet(helper.GetProjectAPIURL(conf.DefaultApiVersion, "/versions"))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
