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
	"testing"

	"github.com/google-github-actions/actions-docs/internal/fileio"
	"github.com/google-github-actions/actions-docs/pkg/parser"
)

func Test_convertInputToTable(t *testing.T) {
	type args struct {
		inputs map[string]parser.ActionInput
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simpleInput", args{map[string]parser.ActionInput{"foo": {Description: "bar"}}}, string(fileio.ReadFile("testdata/simpleInput.md"))},
		{"simpleInputWithRequired", args{map[string]parser.ActionInput{"foo": {Required: true, Description: "bar"}}}, string(fileio.ReadFile("testdata/simpleInputWithRequired.md"))},
		{"simpleInputWithDefault", args{map[string]parser.ActionInput{"bar": {Required: true, Description: "bar", Default: "baz"}}}, string(fileio.ReadFile("testdata/simpleInputWithDefault.md"))},
		{"multipleInputs", args{map[string]parser.ActionInput{"foo": {Required: true, Description: "foo1", Default: "def1"}, "bar": {Required: false, Description: "bar1"}, "baz": {Description: "baz1"}}}, string(fileio.ReadFile("testdata/multipleInputs.md"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertInputToTable(tt.args.inputs); got != tt.want {
				t.Errorf("convertInputToTable() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}

func Test_convertOutputToTable(t *testing.T) {
	type args struct {
		outputs map[string]parser.ActionOutput
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"simpleOutput", args{map[string]parser.ActionOutput{"foo": {Description: "bar"}}}, string(fileio.ReadFile("testdata/simpleOutput.md"))},
		{"multipleOutputs", args{map[string]parser.ActionOutput{"foo": {Description: "foo1"}, "bar": {Description: "bar1"}, "baz": {Description: "baz1"}}}, string(fileio.ReadFile("testdata/multipleOutputs.md"))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertOutputToTable(tt.args.outputs); got != tt.want {

				t.Errorf("convertOutputToTable() = \n%v\n, want \n%v", got, tt.want)
			}
		})
	}
}
