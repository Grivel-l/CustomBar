package main

import (
    "os"
    "fmt"
    "errors"
    "strings"
    "strconv"
    "io/ioutil"
)

func handleLine(config *BarConfig, line string) (error) {
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
    }
    return err
}

func fillConfig(config *BarConfig, appName string) (error) {
    var i       int
    var err     error
    var content []byte
    var lines   []string
    var path    string

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
    for i = 0; i < len(lines); i += 1 {
        err = handleLine(config, lines[i])
        if (err != nil) {
            return errors.New(fmt.Sprintf("Bad value at line %v of config file: %v", i + 1, strings.TrimSpace(strings.Split(lines[i], "=")[1])))
        }
    }
    return nil
}
