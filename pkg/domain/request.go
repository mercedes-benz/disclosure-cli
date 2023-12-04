// SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
//
// SPDX-License-Identifier: MIT

package domain

type RequestCreateVersion struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RequestCssAdd struct {
	URL     string `json:"url"`
	Comment string `json:"comment"`
}

type RequestCreateTag struct {
	Tag string `json:"tag"`
}
