package main

import (
    "fmt"
    /* "github.com/BurntSushi/xgbutil" */
    /* "github.com/BurntSushi/xgb" */
    /* "github.com/BurntSushi/xgbutil/xwindow" */
    "github.com/BurntSushi/xgbutil/ewmh"
    /* "github.com/BurntSushi/xgb/xproto" */
)

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func main() {
    X, err := initX()
    if (err != nil) {
        errorHandler(err)
        return
    }
    window, err := createWindow(X)
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = ewmh.WmWindowTypeSet(window.win.X, window.win.Id, []string{"_NET_WM_WINDOW_DOCK"});
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = ewmh.WmStateSet(window.win.X, window.win.Id, []string{"_NET_WM_STATE_STICKY"});
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = ewmh.WmDesktopSet(window.win.X, window.win.Id, ^uint(0))
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = ewmh.WmNameSet(window.win.X, window.win.Id, "myBar")
    if (err != nil) {
        errorHandler(err)
        return
    }
    for {}
    fmt.Printf("HelloWorld!\n")
}
