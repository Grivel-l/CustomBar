package main

import (
    /* "fmt" */
    "image/color"
    "github.com/BurntSushi/xgbutil/xgraphics"
)

func printString(textType string, content string) (error) {
    var x   int
    var y   int
    var err error

    if (window.pos[textType].xEnd != -1) {
        x = window.pos[textType].xStart
        for (x < window.pos[textType].xEnd) {
            y = 0
            for (y < 40) {
                window.img.SetBGRA(x, y, xgraphics.BGRA{
                    B: 0,
                    G: 0,
                    R: 0,
                    A: 0,
                })
                y += 1
            }
            x += 1
        }
    }
    window.pos[textType].xEnd, _, err = window.img.Text(window.pos[textType].xStart, 0, color.RGBA{
        R: 0xff,
        G: 0xff,
        B: 0xff,
        A: 0xff,
    }, 16, window.font, content)
    if (err != nil) {
        return err
    }
    window.img.XDraw()
    window.img.XPaint(window.win.Id)
    return nil
}
