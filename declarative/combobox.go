// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"errors"
)

import (
	"github.com/lxn/walk"
)

type ComboBox struct {
	AssignTo              **walk.ComboBox
	Name                  string
	Enabled               Property
	Visible               Property
	Font                  Font
	ToolTipText           Property
	MinSize               Size
	MaxSize               Size
	StretchFactor         int
	Row                   int
	RowSpan               int
	Column                int
	ColumnSpan            int
	ContextMenuActions    []*walk.Action
	OnKeyDown             walk.KeyEventHandler
	OnMouseDown           walk.MouseEventHandler
	OnMouseMove           walk.MouseEventHandler
	OnMouseUp             walk.MouseEventHandler
	OnSizeChanged         walk.EventHandler
	Format                string
	Precision             int
	Model                 walk.ListModel
	Value                 Property
	CurrentIndex          Property
	OnCurrentIndexChanged walk.EventHandler
}

func (cb ComboBox) Create(builder *Builder) error {
	w, err := walk.NewComboBox(builder.Parent())
	if err != nil {
		return err
	}

	return builder.InitWidget(cb, w, func() error {
		_, valueIsBind := cb.Value.(Bind)
		_, valueIsBindTo := cb.Value.(BindTo)
		valueBound := valueIsBind || valueIsBindTo
		if _, ok := cb.Model.(walk.BindingValueProvider); !ok && valueBound {
			return errors.New("declarative.ComboBox: Data binding is only supported using a model that implements walk.BindingValueProvider.")
		}

		w.SetFormat(cb.Format)
		w.SetPrecision(cb.Precision)

		if err := w.SetModel(cb.Model); err != nil {
			return err
		}

		if cb.OnCurrentIndexChanged != nil {
			w.CurrentIndexChanged().Attach(cb.OnCurrentIndexChanged)
		}

		if cb.AssignTo != nil {
			*cb.AssignTo = w
		}

		return nil
	})
}

func (w ComboBox) WidgetInfo() (name string, disabled, hidden bool, font *Font, toolTipText string, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, contextMenuActions []*walk.Action, OnKeyDown walk.KeyEventHandler, OnMouseDown walk.MouseEventHandler, OnMouseMove walk.MouseEventHandler, OnMouseUp walk.MouseEventHandler, OnSizeChanged walk.EventHandler) {
	return w.Name, false, false, &w.Font, "", w.MinSize, w.MaxSize, w.StretchFactor, w.Row, w.RowSpan, w.Column, w.ColumnSpan, w.ContextMenuActions, w.OnKeyDown, w.OnMouseDown, w.OnMouseMove, w.OnMouseUp, w.OnSizeChanged
}
