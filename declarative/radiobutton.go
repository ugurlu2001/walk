// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package declarative

import (
	"github.com/lxn/walk"
)

type RadioButton struct {
	// Window

	Background       Brush
	ContextMenuItems []MenuItem
	Enabled          Property
	Font             Font
	MaxSize          Size
	MinSize          Size
	Name             string
	OnKeyDown        walk.KeyEventHandler
	OnKeyPress       walk.KeyEventHandler
	OnKeyUp          walk.KeyEventHandler
	OnMouseDown      walk.MouseEventHandler
	OnMouseMove      walk.MouseEventHandler
	OnMouseUp        walk.MouseEventHandler
	OnSizeChanged    walk.EventHandler
	Persistent       bool
	ToolTipText      Property
	Visible          Property

	// Widget

	AlwaysConsumeSpace bool
	Column             int
	ColumnSpan         int
	Row                int
	RowSpan            int
	StretchFactor      int

	// Button

	OnClicked walk.EventHandler
	Text      Property

	// RadioButton

	AssignTo **walk.RadioButton
	Value    interface{}
}

func (rb RadioButton) Create(builder *Builder) error {
	w, err := walk.NewRadioButton(builder.Parent())
	if err != nil {
		return err
	}

	return builder.InitWidget(rb, w, func() error {
		w.SetValue(rb.Value)

		if rb.OnClicked != nil {
			w.Clicked().Attach(rb.OnClicked)
		}

		if rb.AssignTo != nil {
			*rb.AssignTo = w
		}

		return nil
	})
}
