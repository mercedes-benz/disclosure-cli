// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var versionDetailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Returning the project version details",
	Long:  `The details of the project version`,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		if len(projectVersion) <= 0 {
			fmt.Println("Missing project version")
			os.Exit(1)
		}
		msg := helper.DiscoApiGet(helper.GetProjectVersionAPIURL(conf.DefaultApiVersion, projectVersion, ""))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
