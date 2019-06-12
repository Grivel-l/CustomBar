package parsing

import (
    "../structs"
)

func time(config *structs.TimeConfig, property string, value string) {
    switch (property) {
        case "click":
            if (value == "true") {
                config.Click = true
            } else {
                config.Click = false
            }
    }
}

