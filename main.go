// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"

	"github.com/mercedes-benz/disclosure-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("%s. ", err)
		os.Exit(1)
	}
}
