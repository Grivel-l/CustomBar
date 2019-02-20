package main

import (
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
)

func initAudio() {
    texts["audio"] = widgets.NewQLabel(nil, 0)
    texts["audio"].SetAlignment(core.Qt__AlignRight)
    texts["audio"].SetStyleSheet("color: white; background-color: blue")
}
