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

package ui

import (
	"github.com/mattn/go-gtk/gtk"
)

func dialog(parent *gtk.Window, msg string, t gtk.MessageType) *gtk.MessageDialog {
	messagedialog := gtk.NewMessageDialog(
		parent,
		gtk.DIALOG_MODAL,
		t,
		gtk.BUTTONS_OK,
		msg)
	messagedialog.Response(func() {
		messagedialog.Destroy()
	})
	return messagedialog
}

// ErrorDialog display a window for error messages
func ErrorDialog(parent *gtk.Window, msg string) {
	dialog := dialog(parent, msg, gtk.MESSAGE_ERROR)
	dialog.Run()
}

// InfoDialog display a window for error messages
func InfoDialog(parent *gtk.Window, msg string) {
	dialog := dialog(parent, msg, gtk.MESSAGE_INFO)
	dialog.Run()
}
