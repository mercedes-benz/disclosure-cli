// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package conf

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/jinzhu/configor"
)

const (
	DefaultApiVersion = "v1"
	CLIVersion        = "v0.7"
)

var Config = struct {
	ProjectToken   string `default:""`
	ProjectUUID    string `default:""`
	ProjectVersion string `default:""`
	Host           string `default:""`
}{}

func UserAgent() string {
	return "disclosure-cli / " + CLIVersion
}

func LoadConfig(configFileLocation string) {
	if _, err := os.Stat(configFileLocation); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No config file available in location. Proceeding to check environment variables.")
		// ConfigFileLocation does not exist
	} else {
		err := configor.Load(&Config, configFileLocation)
		if err != nil {
			fmt.Println("Error on create form request data with file " + configFileLocation)
			log.Fatalln(err)
			return
		}
	}
	checkEnvironmentVariables()
}

func checkEnvironmentVariables() {
	Config.ProjectToken = getEnvVariable("INPUT_TOKEN", Config.ProjectToken)
	Config.ProjectUUID = getEnvVariable("INPUT_PROJECT_UUID", Config.ProjectUUID)
	Config.ProjectVersion = getEnvVariable("INPUT_PROJECT_VERSION", Config.ProjectVersion)
	Config.Host = getEnvVariable("INPUT_HOST", Config.Host)
}

func rightPad2Len(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

func dumpStructToSystemOut(parentTitle string, data interface{}) interface{} {
	if reflect.ValueOf(data).Kind() == reflect.Struct {
		v := reflect.ValueOf(data)
		typeOfS := v.Type()
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).Kind() == reflect.Struct {
				dumpStructToSystemOut(typeOfS.Field(i).Name, v.Field(i).Interface())
			} else {
				title := rightPad2Len(strings.ToUpper(parentTitle)+"_"+strings.ToUpper(typeOfS.Field(i).Name), " ", 33)
				valueStr := fmt.Sprintf("%v", v.Field(i).Interface())
				titleLower := strings.ToLower(title)
				if len(valueStr) <= 0 {
					valueStr = "[WARNING is empty]"
				} else if strings.Index(titleLower, "token") > -1 ||
					strings.Index(titleLower, "secret") > -1 ||
					strings.Index(titleLower, "pass") > -1 {
					if strings.Index(titleLower, "preventtokenhijacking") == -1 {
						valueStr = "***"
					}
				}
				fmt.Printf("\n%s = %v", title, valueStr)

			}

		}
	}

	return data
}

func LogConfiguration() {
	fmt.Printf("\n[CONFIGURATION]")
	dumpStructToSystemOut("", Config)
	fmt.Printf("\n\n")
}

func getEnvVariable(envKey string, defaultValue string) string {
	val, exists := os.LookupEnv(envKey)
	if exists {
		return val
	}
	return defaultValue
}

func EnsureApiVerison() {
	Config.Host = strings.TrimSuffix(Config.Host, "/")
	if strings.HasSuffix(Config.Host, "/public") || strings.HasSuffix(Config.Host, "/disco") {
		Config.Host += "/" + DefaultApiVersion
		return
	}
	if !strings.HasSuffix(Config.Host, "/"+DefaultApiVersion) {
		fmt.Println("WARNING: non-default api version set, unexpected behavior is possible")
	}

}
