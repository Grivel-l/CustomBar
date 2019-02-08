package main

import (
    "github.com/BurntSushi/xgbutil/ewmh"
)

func setWindowOptions(config BarConfig) (error) {
    var err error

    err = ewmh.WmNameSet(window.win.X, window.win.Id, "myBar")
    if (err != nil) {
        return err
    }
    err = ewmh.WmWindowTypeSet(window.win.X, window.win.Id, []string{"_NET_WM_WINDOW_TYPE_DOCK"})
    if (err != nil) {
        return err
    }
    err = ewmh.WmStateSet(window.win.X, window.win.Id, []string{"_NET_WM_STATE_STICKY", "_NET_WM_STATE_ABOVE"})
    if (err != nil) {
        return err
    }
    err = ewmh.WmWindowOpacitySet(window.win.X, window.win.Id, config.opacity)
    if (err != nil) {
        return err
    }
    return nil
}

