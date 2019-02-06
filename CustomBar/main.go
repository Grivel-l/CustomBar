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

var config  BarConfig

//export set_volume
func set_volume(volume int) {
    fmt.Printf("Volume is: %v\n", volume);
}

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func main() {
    var err     error
    var X       *xgbutil.XUtil
    var window  Window
    var appName string
    var cstring *C.char

    appName = "custombar"
    cstring = C.CString(appName)
    C.create_con(cstring)
    C.free(unsafe.Pointer(cstring))
    err = fillConfig(appName)
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
    printString(window, "HelloWorld", Pos{x: 0, y: 0})
    window.win.Map()
    for {}
}
