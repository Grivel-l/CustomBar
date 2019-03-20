package main

import (
    "time"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
)

func initDate(signals *Signals) {
    texts["time"] = widgets.NewQLabel(nil, 0)
    texts["time"].SetAlignment(core.Qt__AlignCenter)
    texts["time"].SetStyleSheet("color: white; background-color: red")
    printDate(signals)
}

func printDate(signals *Signals) {
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
    signals.UpdateWidget("time", string(parsed[:]))
    tmp = timestamp.Second()
    time.AfterFunc(time.Duration((60 - tmp) * 1000000000), func() {printDate(signals)})
}
