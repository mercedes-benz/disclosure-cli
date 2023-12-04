// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/domain"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var sbomTagCmd = &cobra.Command{
	Use:   "tag [sbomId] [tag]",
	Short: "Add tag to a sbom",
	Long:  `Add tag to a sbom `,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		data := domain.RequestCreateTag{}
		sbomId := ""

		if len(args) > 1 {
			sbomId = args[0]
			data.Tag = args[1]
		} else {
			helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString("[sbomId] or [tag] is missing in input"))
			os.Exit(1)
		}

		msg := helper.DiscoApiPut(helper.GetProjectVersionAPIURL(projectVersion, "sboms/"+sbomId+"/tag"), data)
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
