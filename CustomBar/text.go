package main

import (
    "image"
    "image/color"
    "github.com/BurntSushi/xgbutil"
    "github.com/BurntSushi/xgb/xproto"
    "github.com/BurntSushi/xgbutil/xgraphics"
    "github.com/BurntSushi/freetype-go/freetype/truetype"
    "golang.org/x/image/font/gofont/goregular"
)

func printString(X *xgbutil.XUtil, id xproto.Window, config BarConfig) (error) {
    var err     error
    var img     *xgraphics.Image
    var font    *truetype.Font

    img = xgraphics.New(X, image.Rect(0, 0, config.width, config.height))
    err = img.XSurfaceSet(id)
    if (err != nil) {
        return err
    }
    font, err = truetype.Parse(goregular.TTF)
    if (err != nil) {
        return err
    }
    img.Text(0, 0, color.RGBA{
        R: 0xff,
        G: 0xff,
        B: 0xff,
        A: 0xff,
    }, 16, font, "HelloWorld")
    img.XDraw()
    img.XPaint(id)
    return nil
}
