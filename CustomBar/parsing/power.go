package parsing

import (
    "../structs"
)

func power(config *structs.PowerConfig, property string, value string) {
    switch (property) {
        case "icon":
            config.Icon = value
    }
}

