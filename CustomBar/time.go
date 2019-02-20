package main

import (
    "time"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
)

func initDate() {
    texts["time"] = widgets.NewQLabel(nil, 0)
    texts["time"].SetAlignment(core.Qt__AlignCenter)
    texts["time"].SetStyleSheet("color: white; background-color: red")
    printDate()
}

func printDate() {
    var tmp         int
    var timestamp   time.Time
    var parsed      [5]byte

    timestamp = time.Now()
    tmp = timestamp.Hour()
    parsed[0] = byte(tmp / 10 + 48)
    parsed[1] = byte(tmp % 10 + 48)
    parsed[2] = ':'
    tmp = timestamp.Minute()
    parsed[3] = byte(tmp / 10 + 48)
    parsed[4] = byte(tmp % 10 + 48)
    texts["time"].SetText(string(parsed[:]))
    tmp = timestamp.Second()
    time.AfterFunc(time.Duration((60 - tmp) * 1000000000), printDate)
}
