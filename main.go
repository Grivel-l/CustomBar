package main

import (
    "fmt"
    "github.com/BurntSushi/xgbutil"
    /* "github.com/BurntSushi/xgb" */
    "github.com/BurntSushi/xgbutil/xwindow"
    /* "github.com/BurntSushi/xgb/xproto" */
)

func main() {
    xutil, err := xgbutil.NewConn()
    /* var xutil *XUtil, err error = xgbutil.NewConn() */
    if (err != nil) {
        return
    }
    /* window, err := xwindow.Generate(xutil) */
    if (err != nil) {
        return
    }
    /* if (err != nil) { */
    /*     return */
    /* } */
    /* conn, err := xgb.NewConn() */
    /* fmt.Printf("Hello: %+v\n", conn); */
    if (err != nil) {
        return
    }
    win := xutil.RootWin()
    if (err != nil) {
        fmt.Printf("Return 1 %v\n", err);
        return
    }
    wini, err := xwindow.Create(xutil, win);
    /* xwindow.Generate(X) */
    xwindow.Generate(xutil)
    if (err != nil) {
        fmt.Printf("Return %v\n", err);
        return
    }
    fmt.Printf("%+v\n", wini);
    /* win.Map() */
    fmt.Printf("HelloWorld!\n")
}
