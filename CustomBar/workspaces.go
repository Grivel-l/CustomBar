package main

import (
    "strings"
    "github.com/therecipe/qt/core"
    "github.com/BurntSushi/xgbutil"
    "github.com/therecipe/qt/widgets"
    "github.com/BurntSushi/xgbutil/ewmh"
)

func getWorkspaces() (uint, error) {
    var err     error
    var xutil   *xgbutil.XUtil

    xutil, err = xgbutil.NewConn()
    if (err != nil) {
        return 0, err
    }
    return ewmh.NumberOfDesktopsGet(xutil)
}

func initWorkspaces(config BarConfig) (error) {
    var i       uint
    var nbr     uint
    var err     error
    var name    strings.Builder

    nbr, err = getWorkspaces()
    if (err != nil) {
        return err
    }
    for i = 0; i < nbr; i += 1 {
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
        texts[name.String()].SetStyleSheet("color: white; background-color: yellow")
        texts[name.String()].SetText("yo")
        name.Reset()
    }
    return nil
}
