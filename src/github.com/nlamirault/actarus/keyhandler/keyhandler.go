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

package keyhandler

import (
	"github.com/mattn/go-gtk/gdk"
)

const (
	// NONE is none key
	NONE Modifier = iota

	// CTRL is Control key
	CTRL
	// FN is Function key
	FN
	// HYPER is Hyper key
	HYPER
	// META is Meta key
	META
	// SUPER is Super key
	SUPER
)

// Modifier represents a key modifier
type Modifier int

// KeyPressEvent represents a keypress consisting of the particular key
// (KeyVal) and possibly a modifier (0 if no modifier is given).
type KeyPressEvent struct {
	KeyVal   int
	Modifier gdk.ModifierType
}

// GetKeyValue returnes the KeyVal value
func (kpe KeyPressEvent) GetKeyValue() int {
	return kpe.KeyVal
}

// GetModifier returns the Modifier used
func (kpe KeyPressEvent) GetModifier() Modifier {
	mod := kpe.Modifier
	switch {
	case mod&gdk.CONTROL_MASK != 0:
		return CTRL
	}
	return NONE
}

// Equals compare two KeyPressEvents
func (kpe KeyPressEvent) Equals(k2 KeyPressEvent) bool {
	return kpe.GetKeyValue() == k2.GetKeyValue() &&
		kpe.GetModifier() == k2.GetModifier()
}
