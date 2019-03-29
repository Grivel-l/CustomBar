package parsing

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
    "../structs"
)

func handleLine(line string, config *structs.BarConfig, module *string) (error) {
    var err     error
    var option  []string
    var value   string

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
    value = strings.TrimSpace(option[1])
    switch (strings.TrimSpace(option[0])) {
        case "margin-top":
            config.MarginTop, err = strconv.Atoi(value)
        case "margin-right":
            config.MarginRight, err = strconv.Atoi(value)
        case "margin-left":
            config.MarginLeft, err = strconv.Atoi(value)
        case "height":
            config.Height, err = strconv.Atoi(value)
        case "width":
            config.Width, err = strconv.Atoi(value)
        case "opacity":
            config.Opacity, err = strconv.ParseFloat(value, 64)
        case "font-size":
            config.FontSize, err = strconv.Atoi(value)
        case "current-workspace":
            config.CurrentWorkspace = value
        case "volume-icon":
            config.VolumeIcon = value
        case "power-icon":
            config.PowerIcon = value
        case "tray-padding":
            config.TrayPadding, err = strconv.Atoi(value)
        case "volume-scroll":
            if (value == "true") {
                config.VolumeScroll = true
            } else {
                config.VolumeScroll = false
            }
        case "workspace-click":
            if (value == "true") {
                config.WorkspaceClick = true
            } else {
                config.WorkspaceClick = false
            }
    }
    return err
}

func defaultConfig(config *structs.BarConfig, width int) {
    config.Height = 33
    config.Width = width
    config.MarginTop = 0
    config.MarginLeft = 0
    config.MarginRight = 0
    config.Opacity = 40
    config.FontSize = 16
    config.CurrentWorkspace = "#0053a0"
    config.VolumeIcon = ""
    config.PowerIcon = ""
    config.TrayPadding = 5
    config.VolumeScroll = true
    config.WorkspaceClick = true
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
        fmt.Printf("Module is: %v\n", module)
        if (err != nil) {
            return fmt.Errorf("Bad value at line %v of config file: %v", i + 1, strings.TrimSpace(strings.Split(lines[i], "=")[1]))
        }
    }
    return nil
}
