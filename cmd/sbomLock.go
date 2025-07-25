package cmd

import (
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var sbomLockCmd = &cobra.Command{
	Use:   "lock [sbomId]",
	Short: "Lock sbom",
	Long:  `Lock sbom`,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		sbomId := ""

		if len(args) > 0 {
			sbomId = args[0]
		} else {
			helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString("[sbomId] is missing in input"))
			os.Exit(1)
		}

		msg := helper.DiscoApiPut(helper.GetProjectVersionAPIURL(projectVersion, "sboms/"+sbomId+"/lock"), struct{}{})
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
