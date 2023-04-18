// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var sbomDetailsCmd = &cobra.Command{
	Use:   "sbomDetails [sbomId]",
	Short: "Details of SBOM",
	Long:  `The sbom status of the project version `,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		sbomId := ""

		if len(args) > 0 {
			sbomId = args[0]
		} else {
			sbomId = "latest"
		}

		msg := helper.DiscoApiGet(helper.GetProjectVersionAPIURL(projectVersion, "sboms/"+sbomId))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
