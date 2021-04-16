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

package parser

import (
	"reflect"
	"testing"

	"github.com/google-github-actions/actions-docs/internal/fileio"
)

func TestParseActionMetadata(t *testing.T) {
	type args struct {
		dat []byte
	}
	tests := []struct {
		name    string
		args    args
		want    ActionDef
		wantErr bool
	}{
		{"simple", args{fileio.ReadFile("testdata/simple.yml")}, ActionDef{
			Name:        "foo",
			Author:      "bar",
			Description: "baz",
		}, false},
		{"simpleNoName", args{fileio.ReadFile("testdata/simpleNoName.yml")}, ActionDef{}, true},
		{"simpleWithInputs", args{fileio.ReadFile("testdata/simpleWithInputs.yml")}, ActionDef{
			Name:        "foo",
			Author:      "bar",
			Description: "baz",
			Inputs: map[string]ActionInput{"input1": {
				Description: "foo",
				Required:    false,
				Default:     "bar",
			},
				"input2": {
					Description: "baz",
					Default:     "qux",
				},
			}}, false},
		{"simpleMissingInputDescription", args{fileio.ReadFile("testdata/simpleMissingInputDescription.yml")}, ActionDef{}, true},
		{"simpleWithInputsOutputs", args{fileio.ReadFile("testdata/simpleWithInputsOutputs.yml")}, ActionDef{
			Name:        "foo",
			Author:      "bar",
			Description: "baz",
			Inputs: map[string]ActionInput{"input1": {
				Description: "foo",
				Required:    false,
				Default:     "bar",
			},
				"input2": {
					Description: "baz",
					Default:     "qux",
				},
			},
			Outputs: map[string]ActionOutput{"output1": {
				Description: "foo",
			}},
		}, false},
		{"simpleMissingOututDescription", args{fileio.ReadFile("testdata/simpleMissingOutputDescription.yml")}, ActionDef{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseActionMetadata(tt.args.dat)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseActionMetadata() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseActionMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
