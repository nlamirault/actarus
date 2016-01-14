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
	//"fmt"

	"github.com/mattn/go-gtk/gtk"

	"github.com/nlamirault/actarus/version"
)

var (
	authors []string
)

func init() {
	authors = []string{
		"Nicolas Lamirault <nicolas.lamirault@gmail.com>",
	}
}

// AboutDialog defines a dialog window
type AboutDialog struct {
	Dialog *gtk.AboutDialog
}

// NewAboutDialog launch an About dialog window
func NewAboutDialog(parent *gtk.Window) {
	dialog := gtk.NewAboutDialog()
	dialog.SetName("Actarus")
	dialog.SetProgramName("Actarus")
	dialog.SetVersion(version.Version)
	dialog.SetAuthors(authors)
	dialog.SetCopyright("Copyright (c) 2015, 2016 Nicolas Lamirault")
	dialog.SetLicense(`
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	  http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
	`)
	dialog.SetWrapLicense(true)
	dialog.Run()
	dialog.Destroy()

	// 	dialog := gtk.NewMessageDialog(
	// 		parent,
	// 		gtk.DIALOG_MODAL,
	// 		gtk.MESSAGE_INFO,
	// 		gtk.BUTTONS_OK,
	// 		`Actarus
	// Copyright (c) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>`)
	// 	dialog.Response(func() {
	// 		dialog.Destroy()
	// 	})
	dialog.Run()
}
