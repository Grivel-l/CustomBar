package main

import (
    "strings"
    "github.com/therecipe/qt/core"
    "github.com/BurntSushi/xgbutil"
    "github.com/therecipe/qt/widgets"
    "./structs"
)

// Signals export
type Signals struct {
    core.QObject
        _ func() `constructor:"init"`
        _ func(app *widgets.QApplication, widget *widgets.QWidget, loop *core.QEventLoop, name string, stylesheet string, xutil *xgbutil.XUtil, config structs.WorkspacesConfig) `slot:"addWidget"`
        _ func(app *widgets.QApplication, loop *core.QEventLoop, workspaces []string, widget *widgets.QWidget, i int, current int, stylesheet string) `slot:"addWorkspace"`
        _ func(widget *widgets.QWidget) `slot:"hideFirstChild"`
        _ func(name string, volume string) `slot:"updateWidget"`
        _ func(order string) `slot:"updateOrder"`
}

func (s *Signals) init() {
    s.ConnectAddWidget(func(app *widgets.QApplication, widget *widgets.QWidget, loop *core.QEventLoop, name string, stylesheet string, xutil *xgbutil.XUtil, config structs.WorkspacesConfig) {
        createWorkspaceWidget(name, xutil, config)
        widget.Layout().ItemAt(0).Layout().AddWidget(texts[name])
        texts[name].SetStyleSheet(stylesheet)
        app.SendEvent(loop, core.NewQEvent(core.QEvent__Quit))
    })
    s.ConnectAddWorkspace(func(app *widgets.QApplication, loop *core.QEventLoop, workspaces []string, widget *widgets.QWidget, i int, current int, stylesheet string) {
        var filler  *widgets.QWidget

        if (i == -1 && current == -1) {
            filler = widgets.NewQWidget(nil, 0)
            filler.SetSizePolicy(widgets.NewQSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__DefaultType))
            widget.Layout().ItemAt(0).Layout().AddWidget(filler)
            app.SendEvent(loop, core.NewQEvent(core.QEvent__Quit))
            return
        }
        widget.Layout().ItemAt(0).Layout().AddWidget(texts[workspaces[i]])
        if (i == current) {
            texts[workspaces[i]].SetStyleSheet(stylesheet)
        } else {
            texts[workspaces[i]].SetStyleSheet("color: white")
        }
        texts[workspaces[i]].Show()
        app.SendEvent(loop, core.NewQEvent(core.QEvent__Quit))
    })
    s.ConnectHideFirstChild(func(widget *widgets.QWidget) {
        var item    *widgets.QWidget
        var layout  *widgets.QLayout

        layout = widget.Layout().ItemAt(0).Layout()
        item = layout.ItemAt(0).Widget()
        item.Hide()
        layout.RemoveWidget(item)
    })
    s.ConnectUpdateWidget(func(name string, text string) {
        texts[name].SetText(text)
    })
    s.ConnectUpdateOrder(func(order string) {
      var builder strings.Builder

      builder.WriteString("Order #")
      builder.WriteString(order)
      texts["olkb"].SetText(builder.String())
    })
}

