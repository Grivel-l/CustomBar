package main

import (
    "os"
    "fmt"
    "errors"
    "strings"
    "strconv"
    "io/ioutil"
)

func handleLine(line string, config *BarConfig) (error) {
    var err     error
    var option  []string
    var value   string

    err = nil
    option = strings.Split(line, "=")
    if (len(option) != 2) {
        return nil
    }
    value = strings.TrimSpace(option[1])
    switch (strings.TrimSpace(option[0])) {
        case "margin-top":
            config.marginTop, err = strconv.Atoi(value)
        case "margin-right":
            config.marginRight, err = strconv.Atoi(value)
        case "margin-left":
            config.marginLeft, err = strconv.Atoi(value)
        case "height":
            config.height, err = strconv.Atoi(value)
        case "width":
            config.width, err = strconv.Atoi(value)
        case "opacity":
            config.opacity, err = strconv.ParseFloat(value, 64)
        case "font-size":
            config.fontSize, err = strconv.Atoi(value)
        case "current-workspace":
            config.currentWorkspace = value
        case "volume-icon":
            config.volumeIcon = value
        case "tray-padding":
            config.trayPadding, err = strconv.Atoi(value)
    }
    return err
}

func defaultConfig(config *BarConfig) {
    config.height = 33
    config.width = 1920
    config.marginTop = 0
    config.marginLeft = 0
    config.marginRight = 0
    config.opacity = 50
    config.fontSize = 16
    config.currentWorkspace = "#0053a0"
    config.volumeIcon = ""
    config.trayPadding = 5
}

func fillConfig(appName string, config *BarConfig) (error) {
    var i       int
    var err     error
    var content []byte
    var lines   []string
    var path    string

    defaultConfig(config)
    path = strings.Join([]string{os.Getenv("HOME"), "/.config/", appName, "/config"}, "")
    _, err = os.Stat(path)
    if (os.IsNotExist(err)) {
        return errors.New("Config file is missing")
    } else if (err != nil && !os.IsExist(err)) {
        return err
    }
    content, err = ioutil.ReadFile(path)
    if (err != nil) {
        return err
    }
    lines = strings.Split(string(content), "\n")
    for i = 0; i < len(lines); i++ {
        err = handleLine(lines[i], config)
        if (err != nil) {
            return fmt.Errorf("Bad value at line %v of config file: %v", i + 1, strings.TrimSpace(strings.Split(lines[i], "=")[1]))
        }
    }
    return nil
}
