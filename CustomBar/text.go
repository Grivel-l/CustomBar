package main

import (
    "image/color"
)

func printString(window Window, content string, pos Pos) {
    window.img.Text(pos.x, pos.y, color.RGBA{
        R: 0xff,
        G: 0xff,
        B: 0xff,
        A: 0xff,
    }, 16, window.font, content)
    window.img.XDraw()
    window.img.XPaint(window.win.Id)
}
