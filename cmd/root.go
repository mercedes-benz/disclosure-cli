// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   getCliName(),
		Short: "A client for the disclosure public api",
		Long:  `A client for the disclosure public api, to manage your projects.`,
	}
)

func getCliName() string {
	if strings.Contains(os.Args[0], "/main") {
		return "disclosure-portal-cli"
	}
	return os.Args[0]
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "configFile", "c", "", "Location of the config file")
	rootCmd.PersistentFlags().StringVarP(&conf.Config.ProjectUUID, "uuid", "u", conf.Config.ProjectUUID, "Uuid of the project")
	rootCmd.PersistentFlags().StringVarP(&conf.Config.ProjectToken, "token", "t", conf.Config.ProjectToken, "Disco project token")
	rootCmd.PersistentFlags().StringVarP(&conf.Config.ProjectVersion, "version", "v", conf.Config.ProjectVersion, "Version of the project")
	rootCmd.PersistentFlags().StringVarP(&conf.Config.Host, "host", "H", conf.Config.Host, "Host of the disclosure portal")

	projectCmd.AddCommand(projectDetailsCmd)
	projectCmd.AddCommand(projectStatus)
	projectCmd.AddCommand(policyRulesCmd)
	projectCmd.AddCommand(schemaCmd)
	projectCmd.AddCommand(onDemandCheckSBOM)
	rootCmd.AddCommand(projectCmd)

	versionCmd.AddCommand(versionListCmd)
	versionCmd.AddCommand(createVersionCmd)
	versionCmd.AddCommand(versionDetailsCmd)
	versionCmd.AddCommand(statusCCSCmd)
	versionCmd.AddCommand(ccsAddCmd)
	versionCmd.AddCommand(sbomsCmd)
	versionCmd.AddCommand(sbomDetailsCmd)
	versionCmd.AddCommand(sbomUploadCmd)
	versionCmd.AddCommand(sbomNoticeCmd)
	versionCmd.AddCommand(sbomStatusCmd)
	rootCmd.AddCommand(versionCmd)

	sbomCmd.AddCommand(sbomTagCmd)
	rootCmd.AddCommand(sbomCmd)

	rootCmd.AddCommand(sha256Cmd)

}

func initConfig() {
	if cfgFile != "" {
		fmt.Println("Path to config file: " + cfgFile)
		conf.LoadConfig(cfgFile)
	} else {
		conf.LoadConfig("./config.yml")
	}

	if !helper.CheckConfigIsValid() {
		rootCmd.MarkPersistentFlagRequired("uuid")
		rootCmd.MarkPersistentFlagRequired("token")
		rootCmd.MarkPersistentFlagRequired("host")
	}
}
