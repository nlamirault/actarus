// Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"fmt"
	"log"

	"github.com/mattn/go-gtk/gtk"

	"github.com/nlamirault/actarus/ui"
)

// Command represent a command from user
type Command struct {
	Name        string
	Action      func(*gtk.Window, string)
	Arg         string
	Description string
}

// Output defines the result of a command
type Output struct {
	Content string
}

var commands []Command

func init() {
	commands = []Command{
		{
			Name:        "about",
			Action:      actionAbout,
			Description: "About Actarus",
		},
		{
			Name:        "help",
			Action:      actionHelp,
			Description: "Show this help",
		},
		{
			Name:        "quit",
			Action:      actionQuit,
			Description: "Quit the session",
		},
	}
}

func actionAbout(parent *gtk.Window, args string) {
	log.Printf("About command")
	ui.NewAboutDialog(parent)
}

func actionHelp(parent *gtk.Window, args string) {
	ui.ErrorDialog(parent, "Not available")
}

func actionQuit(parent *gtk.Window, args string) {
	log.Printf("Quit Actarus : %s\n", args)
	gtk.MainQuit()
}

// Run search for a command and execute it
func Run(cmd string, parent *gtk.Window, args string) {
	for _, command := range commands {
		log.Printf("[DEBUG] Command: %s %v\n", cmd, command)
		if cmd == fmt.Sprintf(":%s", command.Name) {
			command.Action(parent, args)
		}
	}
}
