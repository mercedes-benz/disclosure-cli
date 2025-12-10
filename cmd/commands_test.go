// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/mercedes-benz/disclosure-cli/conf"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

var (
	configPath       = "./../conf.yml"
	spdxFile         = "./../sbom.json"
	ccsReferenceLink = "https://github.com/mercedes-benz"
	sbomId           = "9de2a2a5-637d-44cd-a2ef-e4618bdabd13"
)

func init() {
	err := Execute()
	if err != nil {
		return
	}
}

func ExecuteCommandC(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)

	_, err = root.ExecuteC()

	return buf.String(), err
}

func TestHelpCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Project: get project details",
			cmd:  []string{"-h"},
			want: []string{"Usage:", "Available Commands:", "Flags:"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestProjectDetailsCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Project: get project details",
			cmd:  []string{"project", "details", "-c", configPath},
			want: []string{"name", "uuid", "created", "updated", "schema", "description"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestProjectPolicyRulesCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Project: get policyrules",
			cmd:  []string{"project", "policyrules", "-c", configPath},
			want: []string{"name", "description", "licenses", "type", "created", "updated"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestProjectSbomCheckCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Project: post sbomCheck",
			cmd:  []string{"project", "sbomCheck", spdxFile, "-c", configPath},
			want: []string{"disclaimer", "scanRemarks", "licenseRemarks", "generalRemarks", "components", "spdxId", "license", "name", "version", "status", "remark", "type", "licenseMatched", "description"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestProjectSchemaCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Project: get schema",
			cmd:  []string{"project", "schema", "-c", configPath},
			want: []string{"schema", "id", "title", "type", "properties", "required"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestProjectStatusCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Project: get schema",
			cmd:  []string{"project", "status", "-c", configPath},
			want: []string{"status"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestVersionCssCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Version: get css",
			cmd:  []string{"version", "ccs", "-c", configPath},
			want: []string{"url", "comment", "created", "origin", "uploader"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestVersionCssAddCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Version: post cssAdd",
			cmd:  []string{"version", "ccsAdd", ccsReferenceLink, "-c", configPath},
			want: []string{"success", "message"},
		}, {
			name: "Version: post cssAdd",
			cmd:  []string{"version", "ccsAdd", ccsReferenceLink, "Source code repository", "-c", configPath},
			want: []string{"success", "message"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestVersionDetailsCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Version: get details",
			cmd:  []string{"version", "details", "-c", configPath},
			want: []string{"name", "description", "status"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestVersionListCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Version: get list",
			cmd:  []string{"version", "list", "-c", configPath},
			want: []string{conf.Config.ProjectVersion},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestVersionSbomsCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Version: get sboms",
			cmd:  []string{"version", "sboms", "-c", configPath},
			want: []string{"name", "updated", "valid", "id"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestVersionSbomDetailsCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Version: get sbom details",
			cmd:  []string{"version", "sbomDetails", "-c", configPath},
			want: []string{"name", "id", "version", "creators", "created", "uploaded", "status"},
		}, {
			name: "Version: get sbom details",
			cmd:  []string{"version", "sbomDetails", sbomId, "-c", configPath},
			want: []string{"name", "id", "version", "creators", "created", "uploaded", "status"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestVersionSbomNoticeCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Version: get sbom notice in json format",
			cmd:  []string{"version", "sbomNotice", "-c", configPath},
			want: []string{"components", "name", "version", "licenseName", "licenseID", "copyright", "id", "text", "{", "}"},
		}, {
			name: "Version: get sbom notice in json format",
			cmd:  []string{"version", "sbomNotice", sbomId, "-c", configPath},
			want: []string{"components", "name", "version", "licenseName", "licenseID", "copyright", "id", "text", "{", "}"},
		}, {
			name: "Version: get sbom notice with sbomId in json format",
			cmd:  []string{"version", "sbomNotice", sbomId, "json", "-c", configPath},
			want: []string{"components", "name", "version", "licenseName", "licenseID", "copyright", "id", "text", "{", "}"},
		}, {
			name: "Version: get sbom notice with sbomId in html format",
			cmd:  []string{"version", "sbomNotice", sbomId, "html", "-c", configPath},
			want: []string{"<br>", "<table"},
		}, {
			name: "Version: get sbom notice with sbomId in text format",
			cmd:  []string{"version", "sbomNotice", sbomId, "text", "-c", configPath},
			want: []string{"- Components -", "- Copyright Texts -", "- Licenses -"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestVersionSbomStatusCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Version: get sbom status",
			cmd:  []string{"version", "sbomStatus", "-c", configPath},
			want: []string{"disclaimer", "scanRemarks", "licenseRemarks", "generalRemarks", "components", "spdxId", "license", "name", "version", "status", "remark", "type", "licenseMatched", "description"},
		}, {
			name: "Version: get sbom status",
			cmd:  []string{"version", "sbomStatus", sbomId, "-c", configPath},
			want: []string{"disclaimer", "scanRemarks", "licenseRemarks", "generalRemarks", "components", "spdxId", "license", "name", "version", "status", "remark", "type", "licenseMatched", "description"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}

func TestSbomTagCmd(t *testing.T) {
	tests := []struct {
		name string
		cmd  []string
		want []string
	}{
		{
			name: "Sbom: Add Tag to sbom",
			cmd:  []string{"sbom", "tag", sbomId, "2.2", "-c", configPath},
			want: []string{"success", "true", "message", "Spdx tag updated"},
		},
	}

	for _, tt := range tests {
		msg, err := ExecuteCommandC(rootCmd, tt.cmd...)
		assert.NoError(t, err)

		for _, want := range tt.want {
			assert.Contains(t, msg, want, fmt.Sprintf("got: %s; want: %s ", msg, want))
		}
	}
}
