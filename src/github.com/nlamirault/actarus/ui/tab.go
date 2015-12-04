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

package ui

import (
	"log"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-webkit/webkit"
)

// Browser defines
type Browser struct {
	VBox     *gtk.VBox
	WebView  *webkit.WebView
	URLEntry *gtk.Entry
	Link     string
}

// NewBrowser returns a new browser component
func NewBrowser(uri string) *Browser {
	vbox := gtk.NewVBox(false, 1)

	urlBarEntry := gtk.NewEntry()
	urlBarEntry.SetText(uri)
	urlBarEntry.GrabFocus()
	vbox.PackStart(urlBarEntry, false, false, 0)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)

	webview := webkit.NewWebView()
	swin.Add(webview)
	vbox.Add(swin)

	// urlBarEntry.Connect("activate", func() {
	// 	webview.LoadUri(urlBarEntry.GetText())
	// })
	//return vbox
	browser := &Browser{
		VBox:     vbox,
		WebView:  webview,
		URLEntry: urlBarEntry,
		Link:     "",
	}
	browser.connectSignals()
	urlBarEntry.Emit("activate")
	return browser
}

func (b *Browser) connectSignals() {
	b.WebView.Connect("load-committed", func() {
		b.URLEntry.SetText(b.WebView.GetUri())
	})
	b.WebView.Connect("hovering-over-link", func(ctx *glib.CallbackContext) {
		uri := ctx.Args(1).ToString()
		log.Printf("[DEBUG] URI: %s", uri)
	})
	b.URLEntry.Connect("activate", func() {
		b.WebView.LoadUri(b.URLEntry.GetText())
	})
}
