package main

import (
    "strings"
    "github.com/therecipe/qt/core"
    "github.com/BurntSushi/xgbutil"
    "github.com/therecipe/qt/widgets"
    "github.com/BurntSushi/xgbutil/ewmh"
)

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
    var desktops    []string
    var name        strings.Builder

    desktops, err = getWorkspaces()
    if (err != nil) {
        return err
    }
    for i = 0; i < len(desktops); i++ {
        _, err = name.WriteString("workspace")
        if (err != nil) {
            return err
        }
        err = name.WriteByte(byte(i + 48))
        if (err != nil) {
            return err
        }
        texts[name.String()] = widgets.NewQLabel(nil, 0)
        texts[name.String()].SetAlignment(core.Qt__AlignLeft)
        texts[name.String()].SetStyleSheet("color: white; background-color: black")
        texts[name.String()].SetText(desktops[i])
        name.Reset()
    }
    return nil
}
