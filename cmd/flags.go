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

import "flag"

type actionInputYml struct {
	actionMetadataPath string
	readmePath         string
	dryRun             bool
}

func parseFlags() actionInputYml {
	actionMetadataPath := flag.String("action-metadata", "action.yml", "Path to action metadata file (action.yml)")
	readmePath := flag.String("readme", "README.md", "Path to README.md")
	dryRun := flag.Bool("dry-run", false, "Print to stdout instead of overwriting readme")
	flag.Parse()
	return actionInputYml{*actionMetadataPath, *readmePath, *dryRun}
}
