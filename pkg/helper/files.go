// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package helper

import (
	"fmt"
	"io"
	"os"
)

func ReadFileFromLocalFileSystem(filePath string) (io.ReadCloser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ReadFileFully(filePath string) ([]byte, error) {
	file, err := ReadFileFromLocalFileSystem(filePath)
	if err != nil {
		return nil, err
	}
	result, err := ReadAllAndClose(file)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func ReadTextFile(filePath string) (*string, error) {
	bytes, err := ReadFileFully(filePath)
	if err != nil {
		return nil, err
	}
	resultString := string(bytes)
	return &resultString, nil
}

func ReadAllAndClose(source io.ReadCloser) ([]byte, error) {
	defer source.Close()
	content, err := io.ReadAll(source)
	if err != nil {
		fmt.Printf("Could not read bytes %s", err)
		return nil, err
	}
	return content, nil
}
