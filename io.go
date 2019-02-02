package main

import (
    "fmt"
    "os"
    "errors"
    "strings"
    "io/ioutil"
)

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
        fmt.Printf("Content: %v", lines[i])
    }
    return nil
}
