package main

import (
    "image/color"
)

func printString(window Window, config BarConfig) {
    window.img.Text(0, 0, color.RGBA{
        R: 0xff,
        G: 0xff,
        B: 0xff,
        A: 0xff,
    }, 16, window.font, "HelloWorld")
    window.img.XDraw()
    window.img.XPaint(window.win.Id)
}
