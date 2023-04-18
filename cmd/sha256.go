// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/mercedes-benz/disclosure-cli/pkg/helper"
	"github.com/spf13/cobra"
)

func EncodeSha256ToString(hash hash.Hash) string {
	return hex.EncodeToString(hash.Sum(nil))
}

var sha256Cmd = &cobra.Command{
	Use:   "sha256 [filename]",
	Short: "Generates a sha256 hash",
	Long:  `Generates a sha256 hash from given filename `,
	Run: func(cmd *cobra.Command, args []string) {
		filename := ""
		if len(args) > 0 {
			filename = args[0]
		} else {
			fmt.Println("Missing filename")
			os.Exit(1)
		}
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error while opening the file. Did you provide the correct file name?")
		}
		defer f.Close()

		writer := sha256.New()
		io.Copy(writer, f)

		helper.WriteMessageToOut(cmd, ""+EncodeSha256ToString(writer))
	},
}
