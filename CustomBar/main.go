package main

// #include "./palib.h"
// #cgo pkg-config: libpulse
import "C"

import (
    "fmt"
    "unsafe"
    "errors"
    /* "strconv" */
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

//export set_volume
func set_volume(volume int, config unsafe.Pointer) {
    fmt.Printf("Volume is: %v\n", volume);
    /* printString(window, strconv.Itoa(volume), Pos{x: 1800, y: 0}) */
}

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func initPulseAudio(appName string, config *BarConfig) (error) {
    var cstring *C.char

    cstring = C.CString(appName)
    if (C.create_con(cstring, unsafe.Pointer(config)) != 0) {
        return errors.New("Couldn't init pulseaudio")
    }
    C.free(unsafe.Pointer(cstring))
    return nil
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
    printString(window, "HelloWorld", Pos{x: 0, y: 0})
    window.win.Map()
    for {}
}
