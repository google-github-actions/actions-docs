// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/google-github-actions/actions-docs/pkg/parser"

	"github.com/olekukonko/tablewriter"
)

func convertInputToTable(inputs map[string]parser.ActionInput) string {
	buf := &bytes.Buffer{}
	table := convertToTable([]string{"input id", "description", "required", "default"}, buf)
	keys := getSortedKeys(inputs)
	for _, key := range keys {
		table.Append([]string{fmt.Sprintf("`%s`", key), strings.Replace(inputs[key].Description, "\n", " ", -1), strconv.FormatBool(inputs[key].Required), inputs[key].Default})
	}
	table.Render()
	return buf.String()
}

func convertOutputToTable(outputs map[string]parser.ActionOutput) string {
	buf := &bytes.Buffer{}
	table := convertToTable([]string{"output id", "description"}, buf)
	keys := getSortedKeys(outputs)
	for _, key := range keys {
		table.Append([]string{fmt.Sprintf("`%s`", key), outputs[key].Description})
	}
	table.Render()
	return buf.String()
}

func convertToTable(headings []string, buf *bytes.Buffer) tablewriter.Table {
	table := tablewriter.NewWriter(buf)
	table.SetHeader(headings)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)
	return *table
}

// sorts keys for action input/output to keep md table ordering stable
func getSortedKeys(kv interface{}) []string {
	if kv == nil {
		return nil
	}

	if reflect.TypeOf(kv).Kind() != reflect.Map {
		return nil
	}
	reflectKeys := reflect.ValueOf(kv).MapKeys()
	keys := make([]string, 0, len(reflectKeys))
	for _, key := range reflectKeys {
		keys = append(keys, key.String())
	}
	sort.Strings(keys)
	return keys
}
