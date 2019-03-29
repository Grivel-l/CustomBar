package main

// #include "./palib.h"
import "C"

import (
    "unsafe"
    "github.com/therecipe/qt/gui"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
    "./structs"
)

func initAudio(ctx unsafe.Pointer, config structs.VolumeConfig) {
    var filter  *core.QObject
    var wheelEvent  *gui.QWheelEvent

    texts["audio"] = widgets.NewQLabel(nil, 0)
    texts["audio"].SetAlignment(core.Qt__AlignCenter)
    texts["audio"].SetStyleSheet("color: white")
    if (config.Scroll) {
        filter = core.NewQObject(nil)
        filter.ConnectEventFilter(func (watched *core.QObject, event *core.QEvent) bool {
            if (event.Type() == core.QEvent__Wheel) {
                wheelEvent = gui.NewQWheelEventFromPointer(event.Pointer())
                C.update_volume(ctx, C.int(wheelEvent.AngleDelta().Y()))
            }
            return false
        })
        texts["audio"].SetEnabled(true)
        texts["audio"].InstallEventFilter(filter)
    }
}
