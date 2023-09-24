// Copyright 2023 Olivier Mengué
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Command env is an implementation of the POSIX env command in pure Go.
//
// Usage: env [-i] [name=value ...] [utility [argument ...]]
//
// See https://pubs.opengroup.org/onlinepubs/9699919799/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	args := os.Args[1:]
	env := os.Environ()

	for len(args) > 0 {
		arg := args[0]
		if len(arg) == 0 || arg[0] != '-' {
			break
		}

		if len(arg) == 1 || (len(arg) == 2 && arg[1] == 'i') {
			env = nil // Clear env
		} else {
			fmt.Fprintln(os.Stderr, os.Args[0]+": illegal option --"+arg[1:2])
			fmt.Fprintln(os.Stderr, "usage: env [-i] [name=value ...] [utility [argument ...]]")
			os.Exit(1)
		}

		args = args[1:]
	}

	for len(args) > 0 {
		arg := args[0]
		p := strings.IndexByte(arg, '=')
		if p == -1 {
			break
		}

		key := arg[:p+1]
		found := false

		for i, v := range env {
			if len(v) <= p {
				continue
			}
			if v[:p+1] == key {
				env[i] = arg
				found = true
				break
			}
		}
		if !found {
			env = append(env, arg)
		}

		args = args[1:]
	}

	if len(args) == 0 {
		for _, v := range env {
			os.Stdout.Write([]byte(v))
			fmt.Fprintln(os.Stdout)
		}
	} else {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Env = env
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if errExit, ok := err.(*exec.ExitError); ok {
			os.Exit(errExit.ExitCode())
		} else if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
