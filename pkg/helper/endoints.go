// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package helper

import (
	"fmt"
	"os"

	"github.com/mercedes-benz/disclosure-cli/conf"
)

// BuildAPIURL builds the complete API URL with version
// version: "v1" or "v2"
// path: the endpoint path (e.g., "/projects/uuid")
func BuildAPIURL(version, path string) string {
	baseHost := conf.Config.Host
	return fmt.Sprintf("%s/%s%s", baseHost, version, path)
}

func GetProjectAPIURL(version, appendix string) string {
	if len(conf.Config.ProjectUUID) > 0 {
		return BuildAPIURL(version, "/projects/"+conf.Config.ProjectUUID+appendix)
	} else {
		fmt.Println("Missing flag u - uuid of the project")
		os.Exit(1)
	}
	return ""
}

func GetGroupAPIURL(version, appendix string) string {
	if len(conf.Config.ProjectUUID) > 0 {
		return BuildAPIURL(version, "/groups/"+conf.Config.ProjectUUID+appendix)
	} else {
		fmt.Println("Missing flag u - uuid of the project")
		os.Exit(1)
	}
	return ""
}

func GetProjectVersionAPIURL(version, versionName, appendix string) string {
	if len(versionName) > 0 {
		path := "/projects/" + conf.Config.ProjectUUID + "/versions/" + versionName
		if len(appendix) > 0 {
			path = path + "/" + appendix
		}
		return BuildAPIURL(version, path)
	} else {
		fmt.Println("Missing flag u - versionName of the project")
		os.Exit(1)
	}
	return ""
}
