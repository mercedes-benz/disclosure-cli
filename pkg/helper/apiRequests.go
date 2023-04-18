// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"

	"github.com/mercedes-benz/disclosure-cli/conf"
)

func DiscoApiMultipartPost(url string, completeFilename string) string {
	client := &http.Client{}
	r, w := io.Pipe()
	m := multipart.NewWriter(w)
	go func() {
		defer w.Close()
		defer m.Close()

		_, fileName := path.Split(completeFilename)
		part, err := m.CreateFormFile("file", fileName)
		if err != nil {
			fmt.Println("Error on create form request data with file " + completeFilename)
			os.Exit(1)
		}
		file, err := os.Open(completeFilename)
		if err != nil {
			fmt.Println("Error on open file " + completeFilename)
			os.Exit(1)
		}
		defer file.Close()

		if _, err = io.Copy(part, file); err != nil {
			return
		}
	}()

	req, err := http.NewRequest("POST", url, r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("content-type", m.FormDataContentType())
	req.Header.Set("Authorization", "DISCO"+" "+conf.Config.ProjectToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on requesting url " + url)
		fmt.Println(err)
		os.Exit(1)
	}

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
    	fmt.Println("Operation failed with status", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error on read response of url " + url)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Response from Host " + url)
	return string(body)
}

func DiscoApiPost(url string, v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "DISCO"+" "+conf.Config.ProjectToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on requesting url " + url)
		os.Exit(1)
	}

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
    	fmt.Println("Operation failed with status", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Response from Host " + url)

	return string(body)
}

func DiscoApiGet(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "DISCO"+" "+conf.Config.ProjectToken)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error on requesting url " + url)
		fmt.Println(err)
		os.Exit(1)
	}

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	if !statusOK {
    	fmt.Println("Operation failed with status", resp.Status)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Response from Host " + url)

	return string(body)
}
