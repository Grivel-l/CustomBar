package main

import (
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
)

// Signals export
type Signals struct {
    core.QObject
        _ func() `constructor:"init"`
        _ func(app *widgets.QApplication, widget *widgets.QWidget, loop *core.QEventLoop, name string) `slot:"addWidget"`
        _ func(app *widgets.QApplication, loop *core.QEventLoop, workspaces []string, widget *widgets.QWidget, i int, current int) `slot:"addWorkspace"`
        _ func(widget *widgets.QWidget) `slot:"hideFirstChild"`
}

func (s *Signals) init() {
    s.ConnectAddWidget(func(app *widgets.QApplication, widget *widgets.QWidget, loop *core.QEventLoop, name string) {
        createWorkspaceWidget(name)
        widget.Layout().ItemAt(0).Layout().AddWidget(texts[name])
        texts[name].SetStyleSheet("color: white; background-color: green")
        app.SendEvent(loop, core.NewQEvent(core.QEvent__Quit))
    })
    s.ConnectAddWorkspace(func(app *widgets.QApplication, loop *core.QEventLoop, workspaces []string, widget *widgets.QWidget, i int, current int) {
        widget.Layout().ItemAt(0).Layout().AddWidget(texts[workspaces[i]])
        if (i == current) {
            texts[workspaces[i]].SetStyleSheet("color: white; background-color: green")
        } else {
            texts[workspaces[i]].SetStyleSheet("color: white; background-color: black")
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
}

