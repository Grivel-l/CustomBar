package main

import (
    "github.com/BurntSushi/xgbutil/ewmh"
    "github.com/BurntSushi/xgbutil/xwindow"
)

func setWindowOptions(window *xwindow.Window, config BarConfig) (error) {
    var err error

    err = ewmh.WmNameSet(window.X, window.Id, "myBar")
    if (err != nil) {
        return err
    }
    err = ewmh.WmWindowTypeSet(window.X, window.Id, []string{"_NET_WM_WINDOW_TYPE_DOCK"})
    if (err != nil) {
        return err
    }
    err = ewmh.WmStateSet(window.X, window.Id, []string{"_NET_WM_STATE_STICKY", "_NET_WM_STATE_ABOVE"})
    if (err != nil) {
        return err
    }
    err = ewmh.WmWindowOpacitySet(window.X, window.Id, config.opacity)
    if (err != nil) {
        return err
    }
    return nil
}

