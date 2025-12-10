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

var onDemandCheckSBOM = &cobra.Command{
	Use:   "sbomCheck [fileName]",
	Short: "On demand check for SBOM files",
	Long:  `Uploads a SBOM file and run a check against the project settings`,
	Run: func(cmd *cobra.Command, args []string) {
		fileName := ""
		if len(args) > 0 {
			fileName = args[0]
		} else {
			fmt.Println("Missing filename of SBOM upload")
			os.Exit(1)
		}
		msg := helper.SbomUploadFormData(helper.GetProjectAPIURL(conf.DefaultApiVersion, "/sbomcheck"), fileName, "")
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
