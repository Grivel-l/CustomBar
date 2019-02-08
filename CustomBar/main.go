package main

// #include "./palib.h"
// #cgo pkg-config: libpulse
import "C"

import (
    "fmt"
    "unsafe"
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

//export set_volume
func set_volume(volume int, config unsafe.Pointer, window unsafe.Pointer) {
    var win     Window
    var conf    BarConfig

    conf = *(*BarConfig)(config)
    win = *(*Window)(window)
    fmt.Printf("Volume is: %v, %v, %v\n", volume, conf.height, win.win);
}

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func initPulseAudio(appName string, config *BarConfig, window *Window) {
    var cstring *C.char

    cstring = C.CString(appName)
    C.create_con(cstring, unsafe.Pointer(config), unsafe.Pointer(window))
    C.free(unsafe.Pointer(cstring))
}

func main() {
    var err     error
    var X       *xgbutil.XUtil
    var window  Window
    var appName string
    var config  BarConfig

    appName = "custombar"
    err = fillConfig(appName, &config)
    initPulseAudio(appName, &config, &window)
    if (err != nil) {
        errorHandler(err)
        return
    }
    X, err = initX()
    if (err != nil) {
        errorHandler(err)
        return
    }
    window, err = createWindow(X, config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    err = setWindowOptions(window.win, config)
    if (err != nil) {
        errorHandler(err)
        return
    }
    printString(window, "HelloWorld", Pos{x: 0, y: 0})
    window.win.Map()
    for {}
}
