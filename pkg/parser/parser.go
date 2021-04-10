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
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"
)

// ActionDef holds the action metadata definition
type ActionDef struct {
	Name        string                  `yaml:"name"`
	Author      string                  `yaml:"author"`
	Description string                  `yaml:"description"`
	Inputs      map[string]ActionInput  `yaml:"inputs"`
	Outputs     map[string]ActionOutput `yaml:"outputs"`
}

// ActionInput holds a single action input
type ActionInput struct {
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

// ActionOutput holds a single action output
type ActionOutput struct {
	Description string `yaml:"description"`
}

// ParseActionMetadata parses an action metadata yaml
func ParseActionMetadata(dat []byte) (ActionDef, error) {
	metadata := &ActionDef{}
	err := yaml.Unmarshal(dat, &metadata)
	if err != nil {
		return ActionDef{}, err
	}
	return *metadata, nil
}

// UnmarshalYAML implements yaml.Unmarshaler interface
func (a *ActionDef) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type tmpDefType ActionDef
	var tmpDef tmpDefType
	if err := unmarshal(&tmpDef); err != nil {
		return err
	}
	if tmpDef.Name == "" {
		return errors.New("Action metadata yaml: `name` key is required")
	}
	if tmpDef.Description == "" {
		return errors.New("Action metadata yaml: `description` key is required")
	}
	for key, val := range tmpDef.Inputs {
		if val.Description == "" {
			return fmt.Errorf("Action input %s: `description` key is required", key)
		}
	}
	for key, val := range tmpDef.Outputs {
		if val.Description == "" {
			return fmt.Errorf("Action output %s: `description` key is required", key)
		}
	}
	a.Name = tmpDef.Name
	a.Author = tmpDef.Author
	a.Description = tmpDef.Description
	a.Inputs = tmpDef.Inputs
	a.Outputs = tmpDef.Outputs
	return nil
}
