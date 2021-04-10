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
	"fmt"
	"log"
	"strings"

	"github.com/google-github-actions/actions-docs/internal/fileio"
	"github.com/google-github-actions/actions-docs/pkg/parser"
)

//Main ff
func Main() {
	input := parseFlags()
	dat := fileio.ReadFile(input.actionMetadataPath)
	config, err := parser.ParseActionMetadata(dat)
	if err != nil {
		log.Fatalf("Error parsing metadata yaml: %v", err)
	}
	var inputTable, outputTable string
	if config.Inputs != nil {
		inputTable = convertInputToTable(config.Inputs)
	}
	if config.Outputs != nil {
		outputTable = convertOutputToTable(config.Outputs)
	}

	doc := fileio.ReadFile(input.readmePath)
	docArr := docByLine(doc)
	start, stop, err := findDocInsertPoint(docArr)
	if err != nil {
		log.Fatalf("Error inserting into readme: %v", err)
	}
	finalDoc := (insertIntoDoc(strings.Join(docArr[:start+1], "\n"), strings.Join(docArr[stop:], "\n"), inputTable, outputTable))
	if input.dryRun {
		fmt.Println(finalDoc)
	} else {
		fileio.WriteFile(finalDoc, input.readmePath)
	}

}
