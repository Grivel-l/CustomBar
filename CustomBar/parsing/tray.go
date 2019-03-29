package parsing

import (
    "strconv"
    "../structs"
)

func tray(config *structs.TrayConfig, property string, value string) (error) {
    var err error

    err = nil
    switch (property) {
        case "padding":
            config.Padding, err = strconv.Atoi(value)
    }
    return (err)
}

