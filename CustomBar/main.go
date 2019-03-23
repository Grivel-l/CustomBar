package main

// #include "./events.h"
// #include "./tray.h"
// #cgo pkg-config: x11 xcb xcb-util
import "C"

import (
    "os"
    "fmt"
    "unsafe"
    "github.com/therecipe/qt/gui"
    "github.com/BurntSushi/xgbutil"
    "github.com/therecipe/qt/widgets"
)

// BarConfig export
type BarConfig struct {
    height              int
    width               int
    marginTop           int
    marginRight         int
    marginLeft          int
    fontSize            int
    currentWorkspace    string
    opacity             float64
}

func errorHandler(err error) {
    fmt.Fprintf(os.Stderr, "An error occured: %v\n", err)
}

var texts    map[string]*widgets.QLabel

func initConfigs(app *widgets.QApplication, config BarConfig) {
    var font    *gui.QFont

    font = gui.NewQFont()
    font.SetPixelSize(config.fontSize)
    app.SetFont(font, "")
}

func main() {
    var err         error
    var appName     string
    var signals     *Signals
    var config      BarConfig
    var xutil       *xgbutil.XUtil
    var app         *widgets.QApplication
    var widget      *widgets.QWidget

    appName = "custombar"
    texts = make(map[string]*widgets.QLabel)
    go C.listenClientEvents(unsafe.Pointer(&widget), unsafe.Pointer(&xutil), unsafe.Pointer(&signals), unsafe.Pointer(&app), unsafe.Pointer(&config))
    xutil, err = xgbutil.NewConn()
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = fillConfig(appName, &config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    app, widget = initWindow(config)
    initConfigs(app, config)
    err = initWorkspaces(config, xutil)
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = initPower()
    if (err != nil) {
        errorHandler(err)
        return
    }
    signals = NewSignals(nil)
    err = initPulseAudio(appName, unsafe.Pointer(signals))
    if (err != nil) {
        errorHandler(err)
        return
    }
    initDate(signals)
    createLayout(widget, xutil, config)
    go C.createTrayManager(C.ulong(config.width), C.ulong(config.height), unsafe.Pointer(widget.Layout().ItemAt(2).Layout()))
    app.Exec()
}
