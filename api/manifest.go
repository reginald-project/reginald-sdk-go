// Copyright 2025 Antti Kivi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

// A Manifest is the program representation of a plugin manifest.
type Manifest struct {
	Name        string `json:"name"`
	Domain      string `json:"domain"`
	Description string `json:"description"`
	Executable  string `json:"executable"`
}

// A Command is the program representation of a plugin command that is defined
// in the manifest. The plugin commands should mainly be reserved for utility
// commands and all of the real state changes should be done through the task
// config. When the user exeuctes the command, it is specified as a subcommand
// for the plugin domain on the command line.
type Command struct {
	// Name is the name of the command as it should be written by the user in
	// the terminal.
	Name string `json:"name"`

	// Usage is the one-line usage of the command that is shown to the user in
	// the help message. Usage should not include the plugin domain.
	Usage string `json:"usage"`

	// Description is the description of the command that is shown to the user
	// in the help message.
	Description string `json:"description"`
}

// A Flag is a command-line flag the is defined in the manifest for a plugin
// command. A Flag is always associated with a ConfigEntry in the plugin as when
// the user uses the flag, the value of the flag is set to the value of
// the ConfigEntry and passed to the plugin via that.
type Flag struct {
	// Name is the long name of the flag as it should be written by the user in
	// the terminal. If the name of the flag were "example", the user would
	// write "--example" in the terminal.
	//
	// If the flag name is empty but the flag is associated with a ConfigEntry
	// in the plugin or in a plugin command, the key of the ConfigEntry is used
	// as the name of the flag.
	Name string `json:"name"`

	// Shorthand is the short one-letter name of the flag, used in the form of
	// "-e". The shorthand can be omitted if the flag shouldn't have one.
	Shorthand string `json:"shorthand"`

	// Description is the description of the flag that is shown to the user in
	// the help message.
	Description string `json:"description"`
	Default     string `json:"default"`
	Required    bool   `json:"required"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
	Type  string `json:"type"`
}

// A ConfigEntry is a configuration entry that is defined in the manifest. It
// represents a config value that is supported by the plugin or by a command in
// the plugin. For each ConfigEntry, Reginald will add an entry to the config
// file and automatically create a flag and check the respective environment
// variable for the value of the ConfigEntry. These behaviors can be overridden
// with the fields provided in the ConfigEntry.
type ConfigEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Default     string `json:"default"`
}
