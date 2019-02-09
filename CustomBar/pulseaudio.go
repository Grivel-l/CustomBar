package main

// #include "./palib.h"
// #cgo pkg-config: libpulse
import "C"

import (
    "fmt"
    "image"
    "unsafe"
    "errors"
    "strconv"
    "github.com/BurntSushi/xgbutil/xgraphics"
)

//export set_volume
func set_volume(volume int, config unsafe.Pointer) {
    var err error

    fmt.Printf("Volume is: %v\n", volume);
    err = printString("volume", strconv.Itoa(volume), Pos{x: 1800, y: 0})
    if (err != nil) {
        fmt.Printf("Error: %v\n", err)
    }
}

func initPulseAudio(appName string, config *BarConfig) (error) {
    var cstring *C.char

    cstring = C.CString(appName)
    if (C.create_con(cstring, unsafe.Pointer(config)) != 0) {
        return errors.New("Couldn't init pulseaudio")
    }
    C.free(unsafe.Pointer(cstring))
   window.img["volume"]  = xgraphics.New(window.win.X, image.Rect(0, 0, config.width, config.height))
    return window.img["volume"].XSurfaceSet(window.win.Id)
}

