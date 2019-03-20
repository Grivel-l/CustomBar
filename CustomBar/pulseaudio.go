package main

// #include "./palib.h"
// #cgo pkg-config: libpulse
import "C"

import (
    "unsafe"
    "errors"
    "strconv"
    "strings"
)

//export set_volume
func set_volume(volume int, signalsP unsafe.Pointer) {
    var signals *Signals
    var builder strings.Builder

    builder.WriteString(strconv.Itoa(volume))
    builder.WriteByte('%')
    signals = (*Signals)(signalsP)
    signals.UpdateWidget("audio", builder.String())
}

func initPulseAudio(appName string, signals unsafe.Pointer) (error) {
    var cstring *C.char

    cstring = C.CString(appName)
    if (C.create_con(cstring, signals) == nil) {
        return errors.New("Couldn't init pulseaudio")
    }
    C.free(unsafe.Pointer(cstring))
    initAudio()
    return nil
}

