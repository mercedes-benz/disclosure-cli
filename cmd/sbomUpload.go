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

var sbomUploadCmd = &cobra.Command{
	Use:   "sbomUpload [FileName] [Tag]",
	Short: "Uploads SBOM file to a project version",
	Long:  `Uploads SBOM file to a project version`,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		fileName := ""
		tag := ""
		if len(projectVersion) <= 0 {
			fmt.Println("Missing project version")
			os.Exit(1)
		}
		if len(args) > 0 {
			fileName = args[0]
			if len(args) > 1 {
				tag = args[1]
			}
		} else {
			fmt.Println("Missing filename of SBOM upload")
			os.Exit(1)
		}
		msg := helper.SbomUploadFormData(helper.GetProjectVersionAPIURL(projectVersion, "sboms"), fileName, tag)
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
