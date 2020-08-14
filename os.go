/*
   Copyright 2020 - Jose Gonzalez Krause

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"os/exec"
	"strings"
)

// OS command execution object
type OS struct{}

// Exec executes a given command
func (o *OS) Exec(c string) (out []byte, err error) {
	cmdSlice := strings.Split(c, " ")
	var command string
	var args []string
	if len(cmdSlice) > 1 {
		command = cmdSlice[0]
		args = cmdSlice[1:]
	} else {
		command = cmdSlice[0]
	}

	cmd := exec.Command(command, args...)
	out, err = cmd.CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, err
}
