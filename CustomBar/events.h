#ifndef EVENTS_H
#define EVENTS_H

#include <stdio.h>
#include <X11/Xlib.h>

int     listenClientEvents(void *widget, void *xutil, void *signals, void *app, void *config);

#endif
