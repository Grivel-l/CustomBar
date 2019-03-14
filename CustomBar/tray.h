#ifndef TRAY_H
#define TRAY_H

#include <stdio.h>
#include <X11/Xlib.h>

int createTrayManager(void);

#define SYSTEM_TRAY_REQUEST_DOCK    0
#define SYSTEM_TRAY_BEGIN_MESSAGE   1
#define SYSTEM_TRAY_CANCEL_MESSAGE  2

#endif
