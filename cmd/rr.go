package cmd

import (
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/domain"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var reviewRemarkListCmd = &cobra.Command{
	Use:   "rrs",
	Short: "Version review remarks",
	Long:  `Get the versions review remarks`,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion

		msg := helper.DiscoApiGet(helper.GetProjectVersionAPIURL(conf.DefaultApiVersion, projectVersion, "reviewremarks"))
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}

var reviewRemarkCommentCmd = &cobra.Command{
	Use:   "rrComment [reviewRemarkId] [comment]",
	Short: "Version review remark comment",
	Long:  `Comment on a review remark`,
	Run: func(cmd *cobra.Command, args []string) {
		projectVersion := conf.Config.ProjectVersion
		data := domain.RequestCommentRR{}
		rrId := ""

		if len(args) > 1 {
			rrId = args[0]
			data.Content = args[1]
		} else {
			helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString("[reviewRemarkId] or [comment] is missing in input"))
			os.Exit(1)
		}

		msg := helper.DiscoApiPost(helper.GetProjectVersionAPIURL(conf.DefaultApiVersion, projectVersion, "reviewremarks/"+rrId), data)
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
