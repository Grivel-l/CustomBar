package main

import (
    "fmt"
    "image"
    "github.com/BurntSushi/xgbutil"
    /* "github.com/BurntSushi/xgb" */
    /* "github.com/BurntSushi/xgbutil/xwindow" */
    /* "github.com/BurntSushi/xgb/xproto" */
    "github.com/BurntSushi/xgbutil/xgraphics"
)

func errorHandler(err error) {
    fmt.Printf("An error occured: %v\n", err)
}

func main() {
    var err     error
    var X       *xgbutil.XUtil
    var window  *Window
    var img     *xgraphics.Image

    X, err = initX()
    if (err != nil) {
        errorHandler(err)
        return
    }
    window, err = createWindow(X)
    if (err != nil) {
        errorHandler(err)
        return
    }
    setWindowOptions(window.win);
    img = xgraphics.New(window.win.X, image.Rect(0, 0, 1, 1))
    err = img.XSurfaceSet(window.win.Id)
    if (err != nil) {
        errorHandler(err)
        return
    }
    img.XDraw()
    img.XPaint(window.win.Id)
    for {}
    img.Destroy();
}
