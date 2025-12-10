package cmd

import (
	"fmt"
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var versionDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a project version",
	Long:  `Delete a project version`,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		if len(projectVersion) <= 0 {
			fmt.Println("Missing project version")
			os.Exit(1)
		}
		msg := helper.DiscoApiDelete(helper.GetProjectVersionAPIURL(conf.DefaultApiVersion, projectVersion, ""))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
