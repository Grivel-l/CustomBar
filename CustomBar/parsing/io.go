package parsing

import (
    "os"
    "fmt"
    "strings"
    "io/ioutil"
    "../structs"
)

func handleLine(line string, config *structs.BarConfig, module *string) (error) {
    var err     error
    var option  []string

    switch (line) {
        case "[volume]", "[power]", "[workspaces]", "[tray]", "[general]":
            *module = line[1:len(line) - 1]
            return nil
    }
    option = strings.Split(line, "=")
    if (len(option) != 2) {
        return nil
    }
    err = nil
    switch (*module) {
        case "general":
            err = general(&config.General, strings.TrimSpace(option[0]), strings.TrimSpace(option[1]))
        case "power":
            power(&config.Power, strings.TrimSpace(option[0]), strings.TrimSpace(option[1]))
        case "workspaces":
            workspaces(&config.Workspaces, strings.TrimSpace(option[0]), strings.TrimSpace(option[1]))
        case "tray":
            err = tray(&config.Tray, strings.TrimSpace(option[0]), strings.TrimSpace(option[1]))
        case "volume":
            volume(&config.Volume, strings.TrimSpace(option[0]), strings.TrimSpace(option[1]))
        case "":
        default:
    }
    return err
}

func defaultConfig(config *structs.BarConfig, width int) {
    config.General.Height = 33
    config.General.Height = 33
    config.General.Width = width
    config.General.MarginTop = 0
    config.General.MarginLeft = 0
    config.General.MarginRight = 0
    config.General.Opacity = 40
    config.General.FontSize = 16
    config.Workspaces.CurrentColor = "#0053a0"
    config.Workspaces.Click = true
    config.Volume.Icon = ""
    config.Volume.Scroll = true
    config.Power.Icon = ""
    config.Tray.Padding = 5
}

func FillConfig(appName string, config *structs.BarConfig, width int) (error) {
    var i       int
    var err     error
    var content []byte
    var lines   []string
    var path    string
    var module  string

    defaultConfig(config, width)
    path = strings.Join([]string{os.Getenv("HOME"), "/.config/", appName, "/config"}, "")
    _, err = os.Stat(path)
    if (os.IsNotExist(err)) {
        fmt.Printf("WARNING: Config file is missing, using default config\n")
        return nil
    } else if (err != nil && !os.IsExist(err)) {
        return err
    }
    content, err = ioutil.ReadFile(path)
    if (err != nil) {
        return err
    }
    module = ""
    lines = strings.Split(string(content), "\n")
    for i = 0; i < len(lines); i++ {
        err = handleLine(lines[i], config, &module)
        if (err != nil) {
            return fmt.Errorf("Bad value at line %v of config file: %v", i + 1, strings.TrimSpace(strings.Split(lines[i], "=")[1]))
        }
    }
    return nil
}
