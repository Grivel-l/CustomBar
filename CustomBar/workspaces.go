package main

import "C"

import (
    "os"
    "fmt"
    "unsafe"
    "github.com/therecipe/qt/core"
    "github.com/BurntSushi/xgbutil"
    "github.com/therecipe/qt/widgets"
    "github.com/BurntSushi/xgbutil/ewmh"
)

func createWorkspaceWidget(name string) {
    texts[name] = widgets.NewQLabel(nil, 0)
    texts[name].SetText(name)
    texts[name].SetMinimumWidth(40)
    texts[name].SetAlignment(core.Qt__AlignHCenter | core.Qt__AlignVCenter)
    texts[name].SetStyleSheet("color: white")
}

//export updateWorkspace
func updateWorkspace(widgetP unsafe.Pointer, xutilP unsafe.Pointer, signalsP unsafe.Pointer, appP unsafe.Pointer) {
    var i           int
    var current     uint
    var err         error
    var workspaces  []string
    var signals     *Signals
    var xutil       *xgbutil.XUtil
    var widget      *widgets.QWidget
    var loop        *core.QEventLoop
    var app         *widgets.QApplication

    app = *(**widgets.QApplication)(appP)
    signals = *(**Signals)(signalsP)
    xutil = *(**xgbutil.XUtil)(xutilP)
    widget = *(**widgets.QWidget)(widgetP)
    workspaces, err = ewmh.DesktopNamesGet(xutil)
    if (err != nil) {
        fmt.Fprintf(os.Stderr, err.Error())
        return
    }
    current, err = ewmh.CurrentDesktopGet(xutil)
    if (err != nil) {
        fmt.Fprintf(os.Stderr, err.Error())
        return
    }
    for (!widget.Layout().ItemAt(0).Layout().IsEmpty()) {
        signals.HideFirstChild(widget)
    }
    for i = 0; i < len(workspaces); i++ {
        loop = core.NewQEventLoop(nil)
        if (texts[workspaces[i]] != nil) {
            signals.AddWorkspace(app, loop, workspaces, widget, i, int(current))
        } else {
            signals.AddWidget(app, widget, loop, workspaces[i])
        }
        loop.Exec(core.QEventLoop__AllEvents)
    }
}

func getWorkspacesNbr() (uint, error) {
    var err    error
    var xutil  *xgbutil.XUtil

    xutil, err = xgbutil.NewConn()
    if (err != nil) {
        return 0, err
    }
    return ewmh.NumberOfDesktopsGet(xutil)
}

func initWorkspaces(config BarConfig, xutil *xgbutil.XUtil) (error) {
    var i           int
    var current     uint
    var err         error
    var workspaces  []string

    workspaces, err = ewmh.DesktopNamesGet(xutil)
    if (err != nil) {
        return err
    }
    for i = 0; i < len(workspaces); i++ {
        createWorkspaceWidget(workspaces[i])
    }
    current, err = ewmh.CurrentDesktopGet(xutil)
    if (err != nil) {
        return err
    }
    texts[workspaces[current]].SetStyleSheet("color: white; background-color: green")
    return nil
}
