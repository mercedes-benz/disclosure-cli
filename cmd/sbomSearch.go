// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"os"

	"github.com/mercedes-benz/disclosure-cli/pkg/domain"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var sbomSearchCmd = &cobra.Command{
	Use:   "search [tag]",
	Short: "Search for a sbom in a project",
	Long:  `Search for a sbom in a project by tag `,
	Run: func(cmd *cobra.Command, args []string) {
		data := domain.RequestSbomSearch{}

		if len(args) > 0 {
			data.Tag = args[0]
		} else {
			helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString("[tag] is missing in input"))
			os.Exit(1)
		}

		msg := helper.DiscoApiPost(helper.GetProjectAPIURL("/search"), data)
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
