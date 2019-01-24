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
    var window  *Window

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
    setWindowOptions(window.win)
    err = colorWindow(window.win.X, window.win.Id, &window.img)
    if (err != nil) {
        errorHandler(err)
        return
    }
    for {}
    window.img.Destroy()
}
