package cmd

import (
	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var projectChildrenCmd = &cobra.Command{
	Use:   "children",
	Short: "Returning the project children",
	Long:  `The details of the project children`,
	Run: func(cmd *cobra.Command, args []string) {
		msg := helper.DiscoApiGet(helper.GetGroupAPIURL(conf.DefaultApiVersion, "/children"))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
