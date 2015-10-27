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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"github.com/mattn/go-webkit/webkit"

	"github.com/nlamirault/actarus/events"
	"github.com/nlamirault/actarus/logging"
	"github.com/nlamirault/actarus/version"
)

const (
	application      = "Actarus"
	defaultWinWidth  = 1024
	defaultWinHeight = 768

	homePage = "https://github.com/nlamirault"
)

var (
	port  string
	debug bool
)

func init() {
	// parse flags
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.StringVar(&port, "port", "7070", "port to use")
	flag.Parse()
}

func getApplicationTitle() string {
	return fmt.Sprintf("%s - v%s", application, version.Version)
}

func setupProxy() {
	// Handle proxy
	proxy := os.Getenv("HTTP_PROXY")
	if len(proxy) > 0 {
		soupURI := webkit.SoupUri(proxy)
		webkit.GetDefaultSession().Set("proxy-uri", soupURI)
		soupURI.Free()
	}
}

func runGUI() {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle(getApplicationTitle())
	window.Connect("destroy", gtk.MainQuit)

	keyboardEvent := make(chan interface{})
	window.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		keyboardEvent <- *(**gdk.EventKey)(unsafe.Pointer(&arg))
	})

	// motionEvent := make(chan interface{})
	// window.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
	// 	arg := ctx.Args(0)
	// 	motionEvent <- *(**gdk.EventMotion)(unsafe.Pointer(&arg))
	// })
	// go events.MotionHandler(motionEvent)

	setupProxy()

	vbox := gtk.NewVBox(false, 1)

	urlBarEntry := gtk.NewEntry()
	urlBarEntry.SetText(homePage)
	vbox.PackStart(urlBarEntry, false, false, 0)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)

	webview := webkit.NewWebView()
	webview.Connect("load-committed", func() {
		urlBarEntry.SetText(webview.GetUri())
	})
	swin.Add(webview)

	vbox.Add(swin)

	urlBarEntry.Connect("activate", func() {
		webview.LoadUri(urlBarEntry.GetText())
	})

	statusbar := gtk.NewStatusbar()
	contextID := statusbar.GetContextId("actarus")
	statusbar.Push(contextID, "Welcome to Actarus.")
	vbox.PackStart(statusbar, false, false, 0)

	replEntry := gtk.NewEntry()
	replEntry.Hide()
	vbox.PackEnd(replEntry, false, false, 0)

	window.Add(vbox)

	// window.SetEvents(
	// 	int(gdk.POINTER_MOTION_MASK |
	// 		gdk.POINTER_MOTION_HINT_MASK |
	// 		gdk.BUTTON_PRESS_MASK))
	window.SetSizeRequest(defaultWinWidth, defaultWinHeight)
	window.ShowAll()

	urlBarEntry.Emit("activate")
	replEntry.GrabFocus()
	replEntry.SetVisible(false)

	// Handlers
	go events.KeyboardHandler(keyboardEvent, replEntry)

	gtk.Main()
}

func main() {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	log.Printf("[INFO] Start Actarus")
	runtime.GOMAXPROCS(runtime.NumCPU())
	//runProfiler(fmt.Sprintf("localhost:%s", port))
	runGUI()
}
