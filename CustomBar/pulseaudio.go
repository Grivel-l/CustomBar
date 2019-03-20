package main

// #include "./palib.h"
// #cgo pkg-config: libpulse
import "C"

import (
    "fmt"
    "unsafe"
    "errors"
    "strconv"
)

//export set_volume
func set_volume(volume int, signalsP unsafe.Pointer) {
    var signals *Signals

    signals = *(**Signals)(signalsP)
    fmt.Printf("Updating audio with: %v...\n", strconv.Itoa(volume))
    signals.UpdateWidget("audio", strconv.Itoa(volume))
}

func initPulseAudio(appName string, signals unsafe.Pointer) (error) {
    var cstring *C.char

    cstring = C.CString(appName)
    if (C.create_con(cstring, signals) != 0) {
        return errors.New("Couldn't init pulseaudio")
    }
    C.free(unsafe.Pointer(cstring))
    initAudio()
    return nil
}

