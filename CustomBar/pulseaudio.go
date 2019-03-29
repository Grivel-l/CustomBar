package main

// #include "./palib.h"
// #cgo pkg-config: libpulse
import "C"

import (
    "unsafe"
    "errors"
    "strconv"
    "strings"
    "./structs"
)

//export setVolume
func setVolume(volume int, signalsP unsafe.Pointer, volumeIcon *C.char) {
    var signals *Signals
    var builder strings.Builder

    builder.WriteString(C.GoString(volumeIcon))
    builder.WriteString("  ")
    builder.WriteString(strconv.Itoa(volume))
    builder.WriteByte('%')
    signals = (*Signals)(signalsP)
    signals.UpdateWidget("audio", builder.String())
}

func initPulseAudio(appName string, signals unsafe.Pointer, config structs.VolumeConfig) (error) {
    var cstring *C.char
    var ctx     unsafe.Pointer

    cstring = C.CString(appName)
    ctx = C.create_con(cstring, signals, C.CString(config.Icon))
    if (ctx == nil) {
        return errors.New("Couldn't init pulseaudio")
    }
    C.free(unsafe.Pointer(cstring))
    initAudio(ctx, config)
    return nil
}

