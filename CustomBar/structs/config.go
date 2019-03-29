package structs

type GeneralConfig struct {
    Height      int
    Width       int
    MarginTop   int
    MarginRight int
    MarginLeft  int
    FontSize    int
    Opacity     float64
}

type TrayConfig struct {
    Padding int
}

type VolumeConfig struct {
    Scroll  bool
    Icon    string
}

type WorkspacesConfig struct {
    Click           bool
    CurrentColor    string
}

type PowerConfig struct {
    Icon    string
}

type BarConfig struct {
    Tray       TrayConfig
    Power      PowerConfig
    Volume     VolumeConfig
    General    GeneralConfig
    Workspaces WorkspacesConfig
}

