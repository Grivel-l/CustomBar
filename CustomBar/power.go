package main

import (
     "os"
    "log"
    "fmt"
    "time"
    "strings"
    "strconv"
    "io/ioutil"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
)

func updatePower(remaining int, max int, icon string) {
    var builder     strings.Builder

    builder.WriteString(icon)
    builder.WriteString("  ")
    builder.WriteString(strconv.Itoa(int(float32(remaining) / float32(max) * 100)))
    builder.WriteByte('%')
    texts["power"].SetText(builder.String())
}

func listenEvents(max int, icon string) {
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
    updatePower(remaining, max, icon)
    time.AfterFunc(60000000000, func() {listenEvents(max, icon)})
}

func initPower(icon string) (error) {
    var max     int
    var err     error
    var content []byte

    _, err = os.Stat("/sys/class/power_supply/BAT1/charge_full")
    if (os.IsNotExist(err)) {
        return nil
    }
    err = nil
    texts["power"] = widgets.NewQLabel(nil, 0)
    texts["power"].SetAlignment(core.Qt__AlignCenter)
    texts["power"].SetStyleSheet("color: white")
    content, err = ioutil.ReadFile("/sys/class/power_supply/BAT1/charge_full")
    if (err != nil) {
        return err
    }
    max, err = strconv.Atoi(strings.Trim(string(content), "\n"))
    if (err != nil) {
        return err
    }
    go listenEvents(max, icon)
    return err
}

