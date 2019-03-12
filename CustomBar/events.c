#include "events.h"

extern void updateWorkspace(void *widget, void *xutil, void *signals, void *app);

static int  listenWorkspaces(Display *disp) {
    if (XSelectInput(disp, RootWindow(disp, 0), PropertyChangeMask) == BadWindow)
        return (None);
    return XInternAtom(disp, "_NET_CURRENT_DESKTOP", False);
}

int         listenClientEvents(void *widget, void *xutil, void *signals, void *app) {
    XEvent  event;
    Display *disp;
    Atom    currentDesktop;

    if ((disp = XOpenDisplay(NULL)) == NULL)
        return (1);
    if ((currentDesktop = listenWorkspaces(disp)) == None)
        return (1);
    while (1) {
        XNextEvent(disp, &event);
        if (event.xproperty.atom == currentDesktop)
            updateWorkspace(widget, xutil, signals, app);
        else
            dprintf(1, "Unknown event received...\n");
    }
    XCloseDisplay(disp);
    return (0);
}

