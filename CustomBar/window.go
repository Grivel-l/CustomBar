package main

import (
    "image"
    "golang.org/x/image/font/gofont/goregular"
    "github.com/BurntSushi/xgbutil"
    "github.com/BurntSushi/xgb/xproto"
    "github.com/BurntSushi/xgbutil/xwindow"
    "github.com/BurntSushi/xgbutil/xgraphics"
    "github.com/BurntSushi/freetype-go/freetype/truetype"
)

type Window struct {
    win     *xwindow.Window
    img     map[string]*xgraphics.Image
    font    *truetype.Font
}

func initX() (*xgbutil.XUtil, error) {
    var err error
    var X   *xgbutil.XUtil

    X, err = xgbutil.NewConn()
    if (err != nil) {
        return nil, err
    }
    return X, nil
}

func createWindow(X *xgbutil.XUtil, config BarConfig) (error) {
    var err     error

    window.img = make(map[string]*xgraphics.Image)
    window.win, err = xwindow.Generate(X)
    if (err != nil) {
        return err
    }
    window.win.Create(X.RootWin(),
        config.marginLeft,
        config.marginTop,
        config.width,
        config.height,
        xproto.CwBackPixel,
        0x0)
    window.img["wrapper"] = xgraphics.New(X, image.Rect(0, 0, config.width, config.height))
    window.font, err = truetype.Parse(goregular.TTF)
    if (err != nil) {
        return err
    }
    return window.img["wrapper"].XSurfaceSet(window.win.Id)
}

