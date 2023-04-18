// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/spf13/cobra"
)

func WriteMessageToOut(myCommand *cobra.Command, message string) {
	_, err := myCommand.OutOrStdout().Write([]byte(message + "\n"))
	if err != nil {
		_ = fmt.Errorf("can't write to sdtout of command %s", message)
		return
	}
}

func PrettyJSONString(str string) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return str
	}
	return prettyJSON.String()
}

func CheckConfigIsValid() bool {
	if len(conf.Config.Host) <= 0 {
		return false
	}
	if len(conf.Config.ProjectUUID) <= 0 {
		return false
	}
	if len(conf.Config.ProjectToken) <= 0 {
		return false
	}

	return true
}

// Check if values from an array is in an array and returns value if found
func Contains(args []string, values []string) (bool, string) {
	for _, arg := range args {
		for _, value := range values {
			if strings.ToLower(arg) == value {
				return true, value
			}
		}
	}
	return false, ""
}
