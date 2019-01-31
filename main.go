package main

import (
    "fmt"
    "github.com/BurntSushi/xgbutil"
    /* "github.com/BurntSushi/xgb" */
    /* "github.com/BurntSushi/xgbutil/xwindow" */
    /* "github.com/BurntSushi/xgb/xproto" */
)

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func main() {
    var err     error
    var X       *xgbutil.XUtil
    var window  Window

    X, err = initX()
    if (err != nil) {
        errorHandler(err)
        return
    }
    window, err = createWindow(X)
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = setWindowOptions(window.win)
    if (err != nil) {
        errorHandler(err)
        return
    }
    window.win.Map()
    for {}
}
