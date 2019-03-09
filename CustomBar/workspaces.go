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
    texts[name].SetAlignment(core.Qt__AlignLeft)
    texts[name].SetStyleSheet("color: white; background-color: black")
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
    layout := widget.Layout().ItemAt(0).Layout()
    for (!layout.IsEmpty()) {
        fmt.Printf("Removing...\n")
        layout.RemoveItem(layout.TakeAt(0))
    }
    if (texts[workspaces[current]] == nil) {
        fmt.Printf("Creating new widget\n")
        for i = 0; i < len(workspaces); i++ {
            if (texts[workspaces[i]] != nil) {
                layout.AddWidget(texts[workspaces[i]])
                texts[workspaces[i]].SetStyleSheet("color: white; background-color: black")
            } else {
                loop = core.NewQEventLoop(nil)
                signals.AddWidget(app, widget, loop, workspaces[i])
                loop.Exec(core.QEventLoop__AllEvents)
                texts[workspaces[i]].SetStyleSheet("color: white; background-color: green")
            }
        }
    } else {
        fmt.Printf("All widgets already exist\n")
        for i = 0; i < len(workspaces); i++ {
            layout.AddWidget(texts[workspaces[i]])
            if (uint(i) == current) {
                fmt.Printf("%v to green\n", workspaces[i])
                texts[workspaces[i]].SetStyleSheet("color: white; background-color: green")
            } else {
                fmt.Printf("%v to black\n", workspaces[i])
                texts[workspaces[i]].SetStyleSheet("color: white; background-color: black")
            }
        }
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
    current, err = ewmh.NumberOfDesktopsGet(xutil)
    fmt.Printf("Nbr of desktops: %v\n", current)
    current, err = ewmh.CurrentDesktopGet(xutil)
    if (err != nil) {
        return err
    }
    texts[workspaces[current]].SetStyleSheet("color: white; background-color: green")
    return nil
}
