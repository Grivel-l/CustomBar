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

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

var texts    map[string]*widgets.QLabel

func main() {
    var err     error
    var app     *widgets.QApplication
    var widget  *widgets.QWidget
    var appName string
    var config  BarConfig

    appName = "custombar"
    texts = make(map[string]*widgets.QLabel)
    err = fillConfig(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    app, widget = initWindow(config)
    err = initPulseAudio(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    widget.Layout().AddWidget(texts["audio"])
    fmt.Printf("HelloWorld")
    app.Exec()
}
