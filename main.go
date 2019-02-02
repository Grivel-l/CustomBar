package main

import (
    "fmt"
    "github.com/BurntSushi/xgbutil"
)

type BarConfig struct {
    marginTop   int
    marginRight int
    marginLeft  int
}

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func main() {
    var err     error
    var X       *xgbutil.XUtil
    var window  Window
    var config  BarConfig
    var appName string

    appName = "myBar"
    err = fillConfig(&config, appName)
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
