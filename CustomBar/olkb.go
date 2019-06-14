package main

import (
    "fmt"
    "time"
    "strings"
    "net/http"
    "io/ioutil"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
    "./structs"
)

func initOlkb(signals *Signals, config structs.OlkbConfig) {
    if (!config.Enable) {
        return
    }
    if (config.Order == "") {
        fmt.Printf("error: missing \"Order\" field in \"[olkb]\" module\n")
        return
    }
    texts["olkb"] = widgets.NewQLabel(nil, 0)
    texts["olkb"].SetAlignment(core.Qt__AlignCenter)
    texts["olkb"].SetStyleSheet("color: white")
    go checkOrders(signals, config)
}

func printErr(err error) {
    fmt.Printf("error: Couldn't get olkb order list\n")
    if (err != nil) {
        fmt.Printf("Infos: %v\n", err)
    }
}

func checkOrders(signals *Signals, config structs.OlkbConfig) {
    var i       int
    var err     error
    var data    []byte
    var lines   []string
    var res     *http.Response

    res, err = http.Get("https://raw.githubusercontent.com/olkb/orders/master/README.md")
    if (err != nil || res.StatusCode == 200) {
        data, err = ioutil.ReadAll(res.Body)
        res.Body.Close()
        if (err != nil) {
            printErr(err)
            return
        }
        lines = strings.Split(string(data[:]), "\n")
        for i = 0; i < len(lines); i++ {
            if (strings.Contains(lines[i], config.Order)) {
                signals.UpdateOrder(strings.Trim(strings.Split(lines[i], ".")[0], " "))
                break
            }
        }
    } else {
        printErr(err)
    }
    time.AfterFunc(600000000000, func() {checkOrders(signals, config)})
}

