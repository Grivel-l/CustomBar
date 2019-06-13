package main

import (
    "fmt"
    "time"
    "strings"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
    "./structs"
)

func initDate(signals *Signals, config structs.TimeConfig) {
    var accurate    bool
    var filter      *core.QObject
    var timer       *time.Timer

    accurate = false
    texts["time"] = widgets.NewQLabel(nil, 0)
    texts["time"].SetAlignment(core.Qt__AlignCenter)
    texts["time"].SetStyleSheet("color: white")
    if (config.Click) {
        filter = core.NewQObject(nil)
        filter.ConnectEventFilter(func (watched *core.QObject, event *core.QEvent) bool {
            if (event.Type() == core.QEvent__MouseButtonPress) {
                accurate = !accurate
                timer.Stop()
                printDate(signals, &accurate, &timer)
            }
            return false
        })
        texts["time"].InstallEventFilter(filter)
    }
    printDate(signals, &accurate, &timer)
}

func printDate(signals *Signals, accurate *bool, timer **time.Timer) {
    var tmp         int
    var timestamp   time.Time
    var parsed      [8]byte
    var builder     strings.Builder

    timestamp = time.Now()
    tmp = timestamp.Hour()
    parsed[0] = byte(tmp / 10 + 48)
    parsed[1] = byte(tmp % 10 + 48)
    parsed[2] = ':'
    tmp = timestamp.Minute()
    parsed[3] = byte(tmp / 10 + 48)
    parsed[4] = byte(tmp % 10 + 48)
    if (*accurate) {
        parsed[5] = ':'
        tmp = timestamp.Second()
        parsed[6] = byte(tmp / 10  + 48)
        parsed[7] = byte(tmp % 10  + 48)
        fmt.Fprintf(&builder, "%v %v %v, %v", timestamp.Weekday().String()[:3], timestamp.Day(), timestamp.Month().String(), string(parsed[:]))
    } else {
        fmt.Fprintf(&builder, "%v %v %v, %v", timestamp.Weekday().String()[:3], timestamp.Day(), timestamp.Month().String(), string(parsed[:5]))
    }
    signals.UpdateWidget("time", builder.String())
    tmp = timestamp.Second()
    if (*accurate) {
        *timer = time.AfterFunc(time.Duration(1000000000), func() {printDate(signals, accurate, timer)})
    } else {
        *timer = time.AfterFunc(time.Duration((60 - tmp) * 1000000000), func() {printDate(signals, accurate, timer)})
    }
}
