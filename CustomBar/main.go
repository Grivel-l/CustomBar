package main

import (
    "fmt"
    "github.com/therecipe/qt/widgets"
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

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func main() {
    var err     error
    var app     *widgets.QApplication
    var widget  *widgets.QWidget
    var appName string
    var config  BarConfig

    appName = "custombar"
    err = fillConfig(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    app, widget = initWindow(config)
    widget.Layout()
    err = initPulseAudio(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    app.Exec()
}
