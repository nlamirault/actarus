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

package events

import (
	"log"

	"github.com/mattn/go-gtk/gdk"
)

// MotionHandler handle events from mouse
func MotionHandler(event chan interface{}) {
	for {
		e := <-event
		switch ev := e.(type) {
		case *gdk.EventMotion:
			log.Println("[DEBUG] motion-notify-event:",
				int(ev.X), int(ev.Y))
			break
		default:
			log.Printf("[DEBUG] event: %v\n", ev)
		}
	}
}
