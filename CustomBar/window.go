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
    widget.Show()
    return app, widget
}

func createLayout(widget *widgets.QWidget) (error) {
    var i           int
    var err         error
    var workspaces  []string
    var grid        *widgets.QGridLayout
    var box         [3]*widgets.QBoxLayout

    workspaces, err = getWorkspaces()
    if (err != nil) {
        return err
    }
    grid = widgets.NewQGridLayout2()
    box[0] = widgets.NewQBoxLayout(widgets.QBoxLayout__LeftToRight, nil)
    box[1] = widgets.NewQBoxLayout(widgets.QBoxLayout__LeftToRight, nil)
    box[2] = widgets.NewQBoxLayout(widgets.QBoxLayout__LeftToRight, nil)
    for i = 0; i < len(workspaces); i++ {
        box[0].AddWidget(texts[workspaces[i]], 0, 0)
    }
    box[1].AddWidget(texts["time"], 0, 0)
    box[2].AddWidget(texts["audio"], 0, 0)
    if (texts["power"] != nil) {
        box[2].AddWidget(texts["power"], 0, 0)
    }
    grid.AddLayout(box[0], 0, 0, 0)
    grid.AddLayout(box[1], 0, 1, 0)
    grid.AddLayout(box[2], 0, 2, 0)
    widget.SetLayout(grid)
    widget.SetLayoutDirection(core.Qt__LeftToRight)
    return nil
}

