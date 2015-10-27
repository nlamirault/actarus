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
	"log"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"

	"github.com/nlamirault/actarus/command"
)

// KeyboardHandler handle events from keyboard
func KeyboardHandler(event chan interface{}, repl *gtk.Entry) {
	for {
		e := <-event
		log.Printf("Event : %#v\n", e)
		switch ev := e.(type) {
		case *gdk.EventKey:
			log.Println("[DEBUG] key-press-event: ",
				ev.Keyval)
			// log.Printf("Shift : %v\n", gdk.SHIFT_MASK)
			// log.Printf("Mod1 : %v\n", gdk.MOD1_MASK)
			// log.Printf("Mod2 : %v\n", gdk.MOD2_MASK)
			// log.Printf("Mod3 : %v\n", gdk.MOD3_MASK)
			// log.Printf("Mod4 : %v\n", gdk.MOD4_MASK)
			// log.Printf("Mod5 : %v\n", gdk.MOD5_MASK)
			// log.Printf("Control : %v\n", gdk.CONTROL_MASK)
			// log.Printf("Control : %v\n", gdk.MODIFIER_MASK)
			// log.Printf("Modifier  %d\n", ev.State)
			switch ev.Keyval {
			case gdk.KEY_colon:
				if !repl.IsFocus() {
					repl.SetVisible(true)
					repl.GrabFocus()
					repl.SetText(":")
					repl.SetPosition(1)
				}
				break
			case gdk.KEY_Escape:
				repl.SetVisible(false)
				break
			case gdk.KEY_Return:
				text := repl.GetText()
				log.Printf("Repl text : %s", text)
				if len(text) > 0 {
					command.Run(text, "")
				}
				repl.SetText("")
				break
			case gdk.KEY_q:
				if int(ev.State) == int(gdk.CONTROL_MASK) {
					gtk.MainQuit()
				}
			}
			break
		default:
			log.Printf("[DEBUG] event: %v\n", ev)
		}
	}
}
