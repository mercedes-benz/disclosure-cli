// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var sbomStatusCmd = &cobra.Command{
	Use:   "sbomStatus [sbomId]",
	Short: "Status information of SBOM",
	Long:  `Get sbom status information by sbomId or latest `,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		sbomId := ""

		if len(args) > 0 {
			sbomId = args[0]
		} else {
			sbomId = "latest"
		}

		msg := helper.DiscoApiGet(helper.GetProjectVersionAPIURL(conf.DefaultApiVersion, projectVersion, "sboms/"+sbomId+"/check"))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
