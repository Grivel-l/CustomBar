package main

import (
    "github.com/BurntSushi/xgbutil/ewmh"
    "github.com/BurntSushi/xgbutil/xwindow"
)

func setWindowOptions(window *xwindow.Window) (error) {
    var err error

    err = ewmh.WmWindowTypeSet(window.X, window.Id, []string{"_NET_WM_WINDOW_DOCK"});
    if (err != nil) {
        errorHandler(err)
        return err
    }
    err = ewmh.WmStateSet(window.X, window.Id, []string{"_NET_WM_STATE_STICKY"});
    if (err != nil) {
        errorHandler(err)
        return err
    }
    err = ewmh.WmDesktopSet(window.X, window.Id, ^uint(0))
    if (err != nil) {
        errorHandler(err)
        return err
    }
    err = ewmh.WmNameSet(window.X, window.Id, "myBar")
    if (err != nil) {
        errorHandler(err)
        return err
    }
    return nil
}
