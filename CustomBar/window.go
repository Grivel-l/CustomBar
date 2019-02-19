package main

import (
    "os"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/widgets"
)

func initWindow(config BarConfig) (*widgets.QApplication, *widgets.QWidget) {
    var app     *widgets.QApplication
    var widget  *widgets.QWidget

    app = widgets.NewQApplication(len(os.Args), os.Args)
    widget = widgets.NewQWidget(nil,
            core.Qt__WindowStaysOnTopHint |
            core.Qt__FramelessWindowHint)
    widget.SetMinimumSize2(config.width, config.height)
    widget.SetMaximumSize2(config.width, config.height)
    widget.SetAttribute(core.Qt__WA_X11NetWmWindowTypeDock, true)
    widget.SetWindowOpacity(config.opacity)
    widget.SetStyleSheet("background-color: black")
    widget.SetLayout(widgets.NewQVBoxLayout())
    widget.Show()
    return app, widget
}
