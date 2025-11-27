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

var sbomNoticeCmd = &cobra.Command{
	Use:   "sbomNotice [sbomId] [format]",
	Short: "Get third party notice information for a SBOM as html / json / text ",
	Long:  `Get third party notice information for a SBOM by [id] or [latest] and [format] (html / json / text)`,
	Run: func(cmd *cobra.Command, args []string) {
		availableNoticeFormats := []string{"json", "html", "text"}
		projectVersion := conf.Config.ProjectVersion
		sbomId := ""
		noticeFormat := ""

		if len(projectVersion) <= 0 {
			fmt.Println("Missing project version")
			os.Exit(1)
		}

		containsNoticeFormat, format := helper.Contains(args, availableNoticeFormats)
		if containsNoticeFormat {
			noticeFormat = format
		} else {
			noticeFormat = "json"
		}

		if len(args) > 0 && args[0] != format {
			sbomId = args[0]
		} else if len(args) > 1 && args[1] != format {
			sbomId = args[1]
		} else {
			sbomId = "latest"
		}

		msg := helper.DiscoApiGet(helper.GetProjectVersionAPIURL(conf.DefaultApiVersion, projectVersion, "sboms/"+sbomId+"/notice/"+noticeFormat))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
