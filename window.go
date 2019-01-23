package main

import (
    "github.com/BurntSushi/xgbutil"
    "github.com/BurntSushi/xgbutil/xwindow"
)

type Window struct {
    win *xwindow.Window
}

func initX() (*xgbutil.XUtil, error) {
    X, err := xgbutil.NewConn()
    if (err != nil) {
        return nil, err
    }
    return X, nil
}

func createWindow(X *xgbutil.XUtil) (*Window, error) {
    var err error
    window := new(Window)
    window.win, err = xwindow.Create(X, X.RootWin())
    if (err != nil) {
        return nil, err
    }
    window.win.MoveResize(0, 0, 1000, 1000)
    window.win.Map()
    return window, nil
}

