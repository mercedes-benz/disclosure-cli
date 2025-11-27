// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var sbomPRStatusCmd = &cobra.Command{
	Use:   "prstatus [sbomId]",
	Short: "Policy rule status of sbom",
	Long:  `Policy rule status of sbom`,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		sbomId := ""

		if len(args) > 0 {
			sbomId = args[0]
		} else {
			helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString("[sbomId] is missing in input"))
			os.Exit(1)
		}

		msg := helper.DiscoApiGet(helper.GetProjectVersionAPIURL(conf.DefaultApiVersion, projectVersion, "sboms/"+sbomId+"/status"))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
