package main

import "C"

import (
    "os"
    "fmt"
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
func updateWorkspace() {
    var i           int
    var current     uint
    var err         error
    var workspaces  []string
    var xutil       *xgbutil.XUtil

    xutil, err = xgbutil.NewConn()
    if (err != nil) {
        fmt.Fprintf(os.Stderr, err.Error())
        return
    }
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
    for i = 0; i < len(workspaces); i++ {
        if (uint(i) == current) {
            if (texts[workspaces[i]] == nil) {
                createWorkspaceWidget(workspaces[i])
            }
            texts[workspaces[i]].SetStyleSheet("color: white; background-color: green")
        } else {
            texts[workspaces[i]].SetStyleSheet("color: white; background-color: black")
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
    current, err = ewmh.CurrentDesktopGet(xutil)
    if (err != nil) {
        return err
    }
    texts[workspaces[current]].SetStyleSheet("color: white; background-color: green")
    return nil
}
