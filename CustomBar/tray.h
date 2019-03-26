#ifndef TRAY_H
#define TRAY_H

#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <X11/Xlib.h>
#include <xcb/xcb.h>
#include <xcb/xproto.h>
#include <xcb/xcb_event.h>
#include <xcb/xcb_ewmh.h>

int createTrayManager(size_t width, size_t height, size_t opacity, size_t padding, void *layout);

#define SYSTEM_TRAY_REQUEST_DOCK    0

#endif
