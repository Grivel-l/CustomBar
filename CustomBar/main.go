package main

import (
    "fmt"
    "github.com/BurntSushi/xgbutil"
)

type BarConfig struct {
    height      int
    width       int
    marginTop   int
    marginRight int
    marginLeft  int
    opacity     float64
}

type Pos struct {
    x   int
    y   int
}

var window  Window

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func main() {
    var err     error
    var X       *xgbutil.XUtil
    var appName string
    var config  BarConfig

    appName = "custombar"
    err = fillConfig(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    X, err = initX()
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = createWindow(X, config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = setWindowOptions(config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = initPulseAudio(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    window.pos["wrapper"] = &TextPos{
        xStart: 0,
        xEnd: -1,
    }
    printString("wrapper", "HelloWorld")
    window.win.Map()
    for {}
}
