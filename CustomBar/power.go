package main

import (
    "log"
    "fmt"
    "time"
    "strings"
    "strconv"
    "io/ioutil"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
)

func listenEvents(max int) {
    var remaining   int
    var err         error
    var content     []byte

    content, err = ioutil.ReadFile("/sys/class/power_supply/BAT1/charge_now")
    if (err != nil) {
        log.Println(err)
    }
    remaining, err = strconv.Atoi(strings.Trim(string(content), "\n"))
    fmt.Printf("Remaining: %v, Max: %v\n", remaining, max)
    time.AfterFunc(60000000000, func() {listenEvents(max)})
}

func initPower() (error) {
    var max     int
    var err     error
    var content []byte

    texts["power"] = widgets.NewQLabel(nil, 0)
    texts["power"].SetAlignment(core.Qt__AlignRight)
    texts["power"].SetStyleSheet("color: white; background-color: blue")
    texts["power"].SetText("HelloWorld")
    content, err = ioutil.ReadFile("/sys/class/power_supply/BAT1/charge_full")
    if (err != nil) {
        return err
    }
    max, err = strconv.Atoi(strings.Trim(string(content), "\n"))
    if (err != nil) {
        return err
    }
    go listenEvents(max)
    return err
}

