// Copyright © 2018 Infostellar, Inc.
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

package groundstation

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/infostellarinc/stellarcli/cmd/flag"
	"github.com/infostellarinc/stellarcli/cmd/util"
	"github.com/infostellarinc/stellarcli/pkg/groundstation/uw"
)

var (
	addUWUse   = util.Normalize("add-uw [Ground Station ID] [Start Time] [End Time]")
	addUWShort = util.Normalize("Adds unavailability windows on a ground station.")
	addUWLong  = util.Normalize(
		`Adds unavailability windows on a ground station. Unavailability windows between the given time range
		are returned.`)
)

// Create add-uw command.
func NewAddUWCommand() *cobra.Command {
	outputFormatFlags := flag.NewOutputFormatFlags()
	flags := flag.NewFlagSet(outputFormatFlags)

	command := &cobra.Command{
		Use:   addUWUse,
		Short: addUWShort,
		Long:  addUWLong,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 3 {
				return fmt.Errorf("accepts 3 arg(s), received %d", len(args))
			}

			if err := flags.ValidateAll(); err != nil {
				return err
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			p := outputFormatFlags.ToPrinter()

			startTime, err := util.ParseDateTime(args[1])
			if err != nil {
				log.Fatal(err)
			}

			endTime, err := util.ParseDateTime(args[2])
			if err != nil {
				log.Fatal(err)
			}

			o := &uw.AddUWOptions{
				Printer:   p,
				ID:        args[0],
				StartTime: startTime,
				EndTime:   endTime,
			}

			uw.AddUW(o)
		},
	}

	// Add flags to the command.
	flags.AddAllFlags(command)

	return command
}
