package main

// #include "./palib.h"
// #cgo pkg-config: libpulse
import "C"

import (
    "fmt"
    "unsafe"
    "errors"
)

//export set_volume
func set_volume(volume int, config unsafe.Pointer) {
    /* var err error */

    fmt.Printf("Volume is: %v\n", volume)
    /* err = printString("volume", strconv.Itoa(volume)) */
    /* if (err != nil) { */
    /*     fmt.Printf("Error: %v\n", err) */
    /* } */
    return
}

func initPulseAudio(appName string, config *BarConfig) (error) {
    var cstring *C.char

    cstring = C.CString(appName)
    if (C.create_con(cstring, nil) != 0) {
        return errors.New("Couldn't init pulseaudio")
    }
    C.free(unsafe.Pointer(cstring))
    return nil
}

