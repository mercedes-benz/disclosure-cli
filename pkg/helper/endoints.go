// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package helper

import (
	"fmt"
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
)

func GetProjectAPIURL(appendix string) string {
	if len(conf.Config.ProjectUUID) > 0 {
		return conf.Config.Host + "/v1/projects/" + conf.Config.ProjectUUID + appendix
	} else {
		fmt.Println("Missing flag u - uuid of the project")
		os.Exit(1)
	}
	return ""
}

func GetProjectVersionAPIURL(versionName, appendix string) string {
	if len(versionName) > 0 {
		url := conf.Config.Host + "/v1/projects/" + conf.Config.ProjectUUID + "/versions/" + versionName
		if len(appendix) > 0 {
			url = url + "/" + appendix
		}
		return url
	} else {
		fmt.Println("Missing flag u - versionName of the project")
		os.Exit(1)
	}
	return ""
}
