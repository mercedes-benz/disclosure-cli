// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/domain"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var ccsAddCmd = &cobra.Command{
	Use:   "ccsAdd [url] [comment]",
	Short: "Add reference (url) to complete coresponding source code",
	Long:  `Add link to the corresponding source code of a project version `,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		if len(projectVersion) <= 0 {
			fmt.Println("Missing project version")
			os.Exit(1)
		}

		data := domain.RequestCssAdd{}
		if len(args) > 0 {
			data.URL = args[0]
		} else {
			fmt.Println("The link to your ccs is missing")
			os.Exit(1)
		}
		if len(args) > 1 {
			data.Comment = args[1]
		}

		msg := helper.DiscoApiPost(helper.GetProjectVersionAPIURL(projectVersion, "ccs"), data)
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
