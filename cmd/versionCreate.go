// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"
	"os"

	"github.com/mercedes-benz/disclosure-cli/pkg/domain"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var createVersionCmd = &cobra.Command{
	Use:   "create [name] [description]",
	Short: "Create version",
	Long:  `Creates a project version by name and description`,
	Run: func(cmd *cobra.Command, args []string) {
		data := domain.RequestCreateVersion{}
		if len(args) > 0 {
			data.Name = args[0]
		} else {
			fmt.Println("Version name is missing")
			os.Exit(1)
		}
		if len(args) > 1 {
			data.Description = args[1]
		}

		msg := helper.DiscoApiPost(helper.GetProjectAPIURL("/versions"), data)
		helper.WriteMessageToOut(cmd, ""+helper.PrettyJSONString(msg))
	},
}
