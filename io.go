package main

import (
    "os"
    "errors"
    "strings"
)

func fillConfig(config *BarConfig) (error) {
    var err     error

    _, err = os.Stat(strings.Join([]string{os.Getenv("HOME"), "/.config/myBar/config"}, ""))
    if (os.IsNotExist(err)) {
        return errors.New("Config file is missing")
    } else if (err != nil && !os.IsExist(err)) {
        return err
    }
    return nil
}
