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

func updatePower(remaining int, max int) {
    var percentage  int

    percentage = int(float32(remaining) / float32(max) * 100)
    texts["power"].SetText(strconv.Itoa(percentage))
}

func listenEvents(max int) {
    var remaining   int
    var err         error
    var content     []byte

    content, err = ioutil.ReadFile("/sys/class/power_supply/BAT1/charge_now")
    if (err != nil) {
        log.Println(err)
    }
    remaining, err = strconv.Atoi(strings.Trim(string(content), "\n"))
    if (err != nil) {
        log.Println(err)
    }
    fmt.Printf("Remaining: %v, Max: %v\n", remaining, max)
    updatePower(remaining, max)
    time.AfterFunc(60000000000, func() {listenEvents(max)})
}

func initPower() (error) {
    var max     int
    var err     error
    var content []byte

    texts["power"] = widgets.NewQLabel(nil, 0)
    texts["power"].SetAlignment(core.Qt__AlignRight)
    texts["power"].SetStyleSheet("color: white; background-color: blue")
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

