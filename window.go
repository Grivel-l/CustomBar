package main

import (
    "image"
    "github.com/BurntSushi/xgbutil"
    "github.com/BurntSushi/xgbutil/xwindow"
    "github.com/BurntSushi/xgb/xproto"
    "github.com/BurntSushi/xgbutil/xgraphics"
)

type Window struct {
    win *xwindow.Window
    img xgraphics.Image
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

func createWindow(X *xgbutil.XUtil) (*Window, error) {
    var err     error
    var window  *Window

    window = new(Window)
    window.win, err = xwindow.Create(X, X.RootWin())
    if (err != nil) {
        return nil, err
    }
    window.win.MoveResize(0, 0, 1000, 1000)
    window.win.Map()
    return window, nil
}

func colorWindow(X *xgbutil.XUtil, id xproto.Window, img *xgraphics.Image) (error) {
    var err error

    img = xgraphics.New(X, image.Rect(0, 0, 1, 1))
    err = img.XSurfaceSet(id)
    if (err != nil) {
        return err
    }
    img.XDraw()
    img.XPaint(id)
    return nil
}

