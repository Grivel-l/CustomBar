package parsing

import (
    "strconv"
    "../structs"
)

func general (config *structs.GeneralConfig, property string, value string) (error) {
    var err error

    err = nil
    switch (property) {
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
    }
    return (err)
}

