package main

import (
    "github.com/BurntSushi/xgbutil"
    "github.com/BurntSushi/xgbutil/xwindow"
    "github.com/BurntSushi/xgb/xproto"
)

type Window struct {
    win *xwindow.Window
}

func initX() (*xgbutil.XUtil, error) {
    var err error
    var X   *xgbutil.XUtil

    X, err = xgbutil.NewConn()
    if (err != nil) {
        return nil, err
    }
    return X, nil
}

func createWindow(X *xgbutil.XUtil) (Window, error) {
    var err     error
    var window  Window

    window.win, err = xwindow.Generate(X)
    if (err != nil) {
        return window, err
    }
    window.win.Create(X.RootWin(), 1, 1, 1920, 40, xproto.CwBackPixel, 0x0)
    return window, nil
}

