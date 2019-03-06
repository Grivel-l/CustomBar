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

    workspaces, err = getWorkspaces()
    if (err != nil) {
        fmt.Fprintf(os.Stderr, err.Error())
        return
    }
    xutil, err = xgbutil.NewConn()
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

func getWorkspaces() ([]string, error) {
    var err         error
    var xutil       *xgbutil.XUtil
    var desktops    []string

    xutil, err = xgbutil.NewConn()
    if (err != nil) {
        return desktops, err
    }
    return ewmh.DesktopNamesGet(xutil)
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

func initWorkspaces(config BarConfig) (error) {
    var i           int
    var err         error
    var workspaces  []string

    workspaces, err = getWorkspaces()
    if (err != nil) {
        return err
    }
    for i = 0; i < len(workspaces); i++ {
        createWorkspaceWidget(workspaces[i])
    }
    return nil
}
