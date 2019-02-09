package main

import (
    "image/color"
)

func printString(window Window, content string, pos Pos) (error) {
    var err error

    _, _, err = window.img.Text(pos.x, pos.y, color.RGBA{
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
}
