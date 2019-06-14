package main

import "C"

import (
    "unsafe"
    "strings"
    "strconv"
    "github.com/therecipe/qt/core"
    "github.com/BurntSushi/xgbutil"
    "github.com/therecipe/qt/widgets"
    "github.com/BurntSushi/xgbutil/ewmh"
    "./structs"
)

//export updateMargin
func updateMargin(layoutP unsafe.Pointer, size int) {
    var layout  *widgets.QHBoxLayout

    layout = (*widgets.QHBoxLayout)(layoutP)
    layout.SetContentsMargins(0, 0, size, 0)
}

func initWindow(config structs.GeneralConfig, widget *widgets.QWidget) {
    widget.SetMinimumSize2(config.Width, config.Height)
    widget.SetMaximumSize2(config.Width, config.Height)
    widget.SetAttribute(core.Qt__WA_X11NetWmWindowTypeDock, true)
    widget.SetAttribute(core.Qt__WA_TranslucentBackground, true)
    widget.SetStyleSheet("background-color: rgba(0, 0, 0, 200)")
    widget.Show()
}

func createLayout(widget *widgets.QWidget, xutil *xgbutil.XUtil, config structs.GeneralConfig) (error) {
    var i           int
    var err         error
    var workspaces  []string
    var builder     strings.Builder
    var grid        *widgets.QHBoxLayout
    var box         [3]*widgets.QHBoxLayout

    workspaces, err = ewmh.DesktopNamesGet(xutil)
    if (err != nil) {
        return err
    }
    grid = widgets.NewQHBoxLayout()
    grid.SetContentsMargins(0, 0, 0, 0)
    grid.SetSpacing(0)
    box[0] = widgets.NewQHBoxLayout()
    box[0].SetSpacing(0)
    box[1] = widgets.NewQHBoxLayout()
    box[2] = widgets.NewQHBoxLayout()
    for i = 0; i < len(workspaces); i++ {
        box[0].AddWidget(texts[workspaces[i]], 0, 0)
    }
    box[0].AddWidget(widgets.NewQWidget(nil, 0), 1, 0)
    box[1].AddWidget(texts["time"], 0, 0)
    box[2].AddWidget(widgets.NewQWidget(nil, 0), 1, 0)
    if (texts["olkb"] != nil) {
      box[2].AddWidget(texts["olkb"], 0, 0)
    }
    box[2].AddWidget(texts["audio"], 0, 0)
    texts["audio"].SetContentsMargins(10, 0, 10, 0)
    if (texts["power"] != nil) {
        box[2].AddWidget(texts["power"], 0, 0)
        texts["power"].SetContentsMargins(10, 0, 10, 0)
    }
    grid.AddLayout(box[0], 1)
    grid.AddLayout(box[1], 1)
    grid.AddLayout(box[2], 1)
    grid.SetAlignment2(box[0], core.Qt__AlignLeft)
    grid.SetAlignment2(box[2], core.Qt__AlignRight)
    widget.SetLayout(grid)
    widget.SetLayoutDirection(core.Qt__LeftToRight)
    builder.WriteString("background-color: rgba(0, 0, 0, ")
    builder.WriteString(strconv.Itoa(int(config.Opacity * 255 / 100)))
    builder.WriteByte(')')
    widget.SetStyleSheet(builder.String())
    return nil
}

