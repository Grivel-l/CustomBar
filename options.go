package main

import (
    "github.com/BurntSushi/xgbutil/ewmh"
    "github.com/BurntSushi/xgbutil/xwindow"
)

func setWindowOptions(window *xwindow.Window) (error) {
    var err             error
    /* var strut           ewmh.WmStrut */
    var strutPartial    ewmh.WmStrutPartial

    /* err = ewmh.WmWindowTypeSet(window.X, window.Id, []string{"_NET_WM_WINDOW_DOCK"}); */
    /* if (err != nil) { */
    /*     return err */
    /* } */
    /* err = ewmh.WmStateSet(window.X, window.Id, []string{"_NET_WM_STATE_STICKY"}); */
    /* if (err != nil) { */
    /*     return err */
    /* } */
    err = ewmh.WmDesktopSet(window.X, window.Id, ^uint(0))
    if (err != nil) {
        return err
    }
    err = ewmh.WmNameSet(window.X, window.Id, "myBar")
    if (err != nil) {
        return err
    }
    strutPartial.Left = 0
    strutPartial.Right = 0
    strutPartial.Top = 100
    strutPartial.Bottom = 0
    strutPartial.LeftStartY = 0
    strutPartial.LeftEndY = 0
    strutPartial.RightStartY = 0
    strutPartial.RightEndY = 0
    strutPartial.TopStartX = 1
    strutPartial.TopEndX = 1920
    strutPartial.BottomStartX = 0
    strutPartial.BottomEndX = 0
    err = ewmh.WmStrutPartialSet(window.X, window.Id, &strutPartial);
    if (err != nil) {
        return err
    }
    /* strut.Left = 0 */
    /* strut.Right = 0 */
    /* strut.Top = 100 */
    /* strut.Bottom = 0 */
    /* err = ewmh.WmStrutSet(window.X, window.Id, &strut); */
    /* if (err != nil) { */
    /*     return err */
    /* } */
    return nil
}

