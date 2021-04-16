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
	"flag"
	"os"
	"testing"
)

func Test_parseFlags(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want actionInputYml
	}{
		{"no args", []string{"cmd"}, actionInputYml{"action.yml", "README.md", false}},
		{"dry-run", []string{"cmd", "--dry-run"}, actionInputYml{"action.yml", "README.md", true}},
		{"explicit metadata path", []string{"cmd", "--action-metadata", "/foo/action.yml"}, actionInputYml{"/foo/action.yml", "README.md", false}},
		{"explicit readme path", []string{"cmd", "--readme", "/bar/readme.md"}, actionInputYml{"action.yml", "/bar/readme.md", false}},
		{"explicit readme and metadata path", []string{"cmd", "--action-metadata", "/foo/action.yml", "--readme", "/bar/readme.md"}, actionInputYml{"/foo/action.yml", "/bar/readme.md", false}},
		{"explicit paths and dry-run", []string{"cmd", "--action-metadata", "/foo/action.yml", "--readme", "/bar/readme.md", "--dry-run"}, actionInputYml{"/foo/action.yml", "/bar/readme.md", true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualOsArgs := os.Args
			defer func() {
				os.Args = actualOsArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()

			os.Args = tt.args
			got := parseFlags()
			if got != tt.want {
				t.Errorf("parseFlags() = %v, want %v", got, tt.want)
			}
		})
	}
}
