// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"github.com/nlamirault/actarus/version"
)

// Command represent a command from user
type Command struct {
	Name        string
	Action      func(*gtk.Window, string) (*Output, error)
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

func actionAbout(parent *gtk.Window, args string) (*Output, error) {
	log.Printf("About command")
	// dialog := ui.NewAboutDialog()
	// dialog.Dialog.Run()
	ui.NewAboutDialog(parent)
	return &Output{
		Content: fmt.Sprintf("Version :%s", version.Version),
	}, nil
}

func actionHelp(parent *gtk.Window, args string) (*Output, error) {
	ui.ErrorDialog(parent, "Not available")
	return nil, nil
}

func actionQuit(parent *gtk.Window, args string) (*Output, error) {
	log.Printf("Quit Actarus : %s\n", args)
	gtk.MainQuit()
	return nil, nil
}

func Run(cmd string, parent *gtk.Window, args string) error {
	for _, command := range commands {
		log.Printf("[DEBUG] Command: %s %v\n", cmd, command)
		if cmd == fmt.Sprintf(":%s", command.Name) {
			_, err := command.Action(parent, args)
			return err
		}
	}
	return nil
}
