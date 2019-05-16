package main

// #include "./events.h"
import "C"

import (
    "os"
    "fmt"
    "unsafe"
    "strings"
    "github.com/therecipe/qt/core"
    "github.com/BurntSushi/xgbutil"
    "github.com/therecipe/qt/widgets"
    "github.com/BurntSushi/xgbutil/ewmh"
    "./structs"
)

func createWorkspaceWidget(name string, xutil *xgbutil.XUtil, config structs.WorkspacesConfig) {
    var i           int
    var err         error
    var workspaces  []string
    var filter      *core.QObject

    texts[name] = widgets.NewQLabel(nil, 0)
    texts[name].SetText(name)
    texts[name].SetMinimumWidth(40)
    texts[name].SetAlignment(core.Qt__AlignHCenter | core.Qt__AlignVCenter)
    texts[name].SetStyleSheet("color: white")
    texts[name].SetEnabled(true)
    if (config.Click) {
        filter = core.NewQObject(nil)
        filter.ConnectEventFilter(func (watched *core.QObject, event *core.QEvent) bool {
            if (event.Type() == core.QEvent__MouseButtonPress) {
                workspaces, err = ewmh.DesktopNamesGet(xutil)
                if (err != nil) {
                    fmt.Fprintf(os.Stderr, "Error: Couldn't get workspaces\n")
                }
                for i = 0; workspaces[i] != name; i++ {
                    ;
                }
                C.sendClientMessage(C.CString("_NET_CURRENT_DESKTOP"), C.int(i))
            }
            return false
        })
        texts[name].InstallEventFilter(filter)
    }
}

func getStylesheet(color string) (string) {
    var builder strings.Builder

    builder.WriteString("color: white; background-color: ")
    builder.WriteString(color)
    return builder.String()
}

//export updateWorkspace
func updateWorkspace(widgetP unsafe.Pointer, xutilP unsafe.Pointer, signalsP unsafe.Pointer, appP unsafe.Pointer, configP unsafe.Pointer) {
    var i           int
    var current     uint
    var err         error
    var stylesheet  string
    var workspaces  []string
    var signals     *Signals
    var xutil       *xgbutil.XUtil
    var widget      *widgets.QWidget
    var loop        *core.QEventLoop
    var config      structs.WorkspacesConfig
    var app         *widgets.QApplication

    app = *(**widgets.QApplication)(appP)
    signals = *(**Signals)(signalsP)
    xutil = *(**xgbutil.XUtil)(xutilP)
    widget = *(**widgets.QWidget)(widgetP)
    workspaces, err = ewmh.DesktopNamesGet(xutil)
    config = *(*structs.WorkspacesConfig)(configP)
    stylesheet = getStylesheet(config.CurrentColor)
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
            signals.AddWorkspace(app, loop, workspaces, widget, i, int(current), stylesheet)
        } else {
            signals.AddWidget(app, widget, loop, workspaces[i], stylesheet, xutil, config)
        }
        loop.Exec(core.QEventLoop__AllEvents)
    }
    signals.AddWorkspace(app, loop, workspaces, widget, -1, -1, stylesheet)
    loop.Exec(core.QEventLoop__AllEvents)
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

func initWorkspaces(config structs.WorkspacesConfig, xutil *xgbutil.XUtil) (error) {
    var i           int
    var current     uint
    var err         error
    var workspaces  []string
    var builder     strings.Builder

    workspaces, err = ewmh.DesktopNamesGet(xutil)
    if (err != nil) {
        return err
    }
    for i = 0; i < len(workspaces); i++ {
        createWorkspaceWidget(workspaces[i], xutil, config)
    }
    current, err = ewmh.CurrentDesktopGet(xutil)
    if (err != nil) {
        return err
    }
    builder.WriteString("color: white; background-color: ")
    builder.WriteString(config.CurrentColor)
    texts[workspaces[current]].SetStyleSheet(builder.String())
    return nil
}
