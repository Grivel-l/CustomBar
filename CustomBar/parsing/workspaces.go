package parsing

import (
    "../structs"
)

func workspaces(config *structs.WorkspacesConfig, property string, value string) {
    switch (property) {
        case "click":
            if (value == "true") {
                config.Click = true
            } else {
                config.Click = false
            }
        case "current-color":
            config.CurrentColor = value
        case "position":
            config.Position = value
        case "alignment":
            config.Alignment = value
    }
}

