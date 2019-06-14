package parsing

import (
    "../structs"
)

func olkb(config *structs.OlkbConfig, property string, value string) {
    config.Enable = true
    switch (property) {
        case "order":
            config.Order = value
    }
}

