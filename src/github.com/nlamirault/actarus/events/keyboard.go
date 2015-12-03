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

package events

import (
	// "fmt"
	"log"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"

	"github.com/nlamirault/actarus/command"
	"github.com/nlamirault/actarus/keyhandler"
	"github.com/nlamirault/actarus/ui"
)

// KeyboardHandler handle events from keyboard
func KeyboardHandler(event chan *keyhandler.KeyPressEvent, window *gtk.Window,
	repl *gtk.Entry, notebook *gtk.Notebook) {
	for {
		kpe := <-event
		log.Printf("[DEBUG] KeyPressEvent : %v", kpe)
		gdk.ThreadsEnter()
		switch kpe.KeyVal {
		case gdk.KEY_Escape:
			repl.SetVisible(false)
			break
		case gdk.KEY_colon:
			if !repl.IsFocus() {
				repl.SetVisible(true)
				repl.GrabFocus()
				repl.SetText(":")
				repl.SetPosition(1)
			}
			break
		case gdk.KEY_Return:
			if repl.IsFocus() {
				text := repl.GetText()
				log.Printf("Repl text : %s", text)
				if len(text) > 0 {
					command.Run(text, window, "")
				}
				repl.SetText("")
			}
			break
		// case gdk.KEY_w:
		// 	if kpe.GetModifier() == keyhandler.CTRL {
		// 		log.Printf("[DEBUG] nb : %d", notebook.GetNPages())
		// 		notebook.RemovePage(notebook.GetCurrentPage())
		// 		log.Printf("[DEBUG] nb : %d", notebook.GetNPages())
		// 	}
		// 	break
		case gdk.KEY_t:
			if kpe.GetModifier() == keyhandler.CTRL {
				log.Printf("[DEBUG] New tab")
				log.Printf("[DEBUG] nb : %d", notebook.GetNPages())
				log.Printf("[DEBUG] current : %d",
					notebook.GetCurrentPage())
				tab := ui.NewBrowser("")
				page := gtk.NewFrame("")
				//fmt.Sprintf("%d", notebook.GetNPages()+1))
				notebook.AppendPage(page, gtk.NewLabel("New tab"))
				page.Add(tab.VBox)
				log.Printf("[DEBUG] nb : %d", notebook.GetNPages())
				notebook.ShowAll()
			}
			break
		case gdk.KEY_q:
			if kpe.GetModifier() == keyhandler.CTRL {
				gtk.MainQuit()
			}
			break
		}
		gdk.ThreadsLeave()
	}
}
