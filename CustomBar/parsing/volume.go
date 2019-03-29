package parsing

import (
    "../structs"
)

func volume(config *structs.VolumeConfig, property string, value string) {
    switch (property) {
        case "scroll":
            if (value == "true") {
                config.Scroll = true
            } else {
                config.Scroll = false
            }
        case "icon":
            config.Icon = value
    }
}

