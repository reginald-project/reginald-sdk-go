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

// The supported value types for a KeyValue.
const (
	BoolValue   ValueType = "bool"
	IntValue    ValueType = "int"
	StringValue ValueType = "string"
)

// ValueType is used as the type indicator of a KeyValue.
type ValueType string

// A Manifest is the program representation of a plugin manifest.
type Manifest struct {
	// Name is the human-readable name of the plugin. It must be unique.
	Name string `json:"name"`

	// Domain is the domain of the plugin that is used to identify the plugin in
	// commands and tasks. It must be unique.
	//
	// All of the commands defined by the plugin are subcommands of the domain.
	// The domain is also used as a prefix for the tasks provided by the plugin.
	Domain string `json:"domain"`

	// Description is the description of the plugin that is shown to the user in
	// the help message.
	Description string `json:"description"`

	// Executable is the name of the executable file of the plugin in
	// the plugin's directory.
	Executable string `json:"executable"`

	// Config is a list of ConfigEntries that are used to define
	// the configuration of the plugin.
	Config []ConfigEntry `json:"config,omitempty"`

	// Commands is a list of Commands that this plugin provides.
	Commands []Command `json:"commands,omitempty"`

	// Tasks is a list of Tasks that this plugin provides.
	Tasks []Task `json:"tasks,omitempty"`
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

	// Aliases is a list of aliases for the command that can be used instead of
	// Name to run this command.
	Aliases []string `json:"aliases,omitempty"`

	// Config is a list of ConfigEntries that are used to define
	// the configuration of the command.
	Config []ConfigEntry `json:"config,omitempty"`
}

// A Task is the program representation of a plugin task that is defined in
// the manifest.
type Task struct {
	// Type is the name of the task type that is used to identify this task in
	// the config file. The plugin's domain will be used as a prefix for
	// the task type when a task is defined in the config file.
	Type string `json:"type"`

	// Description is the description of the task that is shown to the user in
	// the help message.
	Description string `json:"description"`

	// Config is a list of KeyValues that are used to define the configuration
	// of the task.
	Config []KeyValue `json:"config,omitempty"`
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

	// TODO: Add inverse flag for booleans.
}

// A KeyValue is a key-value pair that is used to define a config value in the
// manifest. Depending on the context, it is used either as a part of
// a ConfigEntry or as is.
type KeyValue struct {
	// Key is the key of the KeyValue as it would be written in, for example,
	// the config file.
	Key string `json:"key"`

	// Value is the current value of KeyValue as the type it should be defined
	// as. When the plugin provides the initial KeyValue, for example, as a part
	// of the manifest, Value should contain the default value of the KeyValue.
	// When KeyValue is used to send data from Reginald to the plugin, Value
	// contains the current value of the KeyValue.
	Value any `json:"value"`

	// Type is a string representation of the type of the value that this
	// KeyValue holds.
	Type ValueType `json:"type"`
}

// A ConfigEntry is a configuration entry that is defined in the manifest. It
// represents a config value that is supported by the plugin or by a command in
// the plugin. For each ConfigEntry, Reginald will add an entry to the config
// file and automatically create a flag and check the respective environment
// variable for the value of the ConfigEntry. These behaviors can be overridden
// with the fields provided in the ConfigEntry.
type ConfigEntry struct {
	KeyValue

	// Flag contains the information on the possible command-line flag that is
	// associated with this ConfigEntry. Flag must be nil if the ConfigEntry has
	// no associated flag. Otherwise, its type must match [Flag].
	Flag *Flag `json:"flag,omitempty"`

	// EnvOverride optionally defines a string to use in the environment
	// variable name instead of the automatic name of the variable that will be
	// composed using the Key in the embedded [KeyValue]. It is appended after
	// the prefix `REGINALD_` but if EnvOverride is used to set the name of
	// the environment variable, the name of the plugin or the name of
	// the command is not added to variable name automatically.
	EnvOverride string `json:"envOverride,omitempty"`

	// FlagOnly tells Reginald whether this ConfigEntry should only be
	// controlled by a command-line flag. If this is set to true, Reginald won't
	// read the value of this ConfigEntry from the config file or from
	// environment variables.
	FlagOnly bool `json:"flagOnly,omitempty"`
}
