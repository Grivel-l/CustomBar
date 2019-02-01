package main

import (
    "fmt"
    "github.com/BurntSushi/xgbutil"
)

type BarConfig struct {
}

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func main() {
    var err     error
    var X       *xgbutil.XUtil
    var window  Window
    var config  BarConfig;

    err = fillConfig(&config)
    if (err != nil) {
        errorHandler(err)
        return
    }
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
