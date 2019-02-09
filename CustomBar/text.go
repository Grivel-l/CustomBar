package main

import (
    /* "fmt" */
    "image/color"
)

func printString(img string, content string, pos Pos) (error) {
    var err error

    window.img[img].ForExp(func(x int, y int) (r, g, b, a uint8) {
        return 0, 0, 0, 0
    })
    _, _, err = window.img[img].Text(pos.x, pos.y, color.RGBA{
        R: 0xff,
        G: 0xff,
        B: 0xff,
        A: 0xff,
    }, 16, window.font, content)
    if (err != nil) {
        return err
    }
    window.img[img].XDraw()
    window.img[img].XPaint(window.win.Id)
    return nil
}
