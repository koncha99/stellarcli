//
// Copyright © 2019 Infostellar, Inc.
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

package flag

import "github.com/spf13/cobra"

type DebugFlag struct {
	IsDebug bool
}

// Add flags to the command.
func (f *DebugFlag) AddFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&f.IsDebug, "debug", "", false, "Output debug information. (default false)")
}

// Validate flag values.
func (f *DebugFlag) Validate() error {
	return nil
}

// Create a new DebugFlag.
func NewDebugFlag() *DebugFlag {
	return &DebugFlag{}
}
