#include "events.h"

extern void updateWorkspace(void);

static int  listenWorkspaces(Display *disp) {
    if (XSelectInput(disp, RootWindow(disp, 0), PropertyChangeMask) == BadWindow)
        return (None);
    return XInternAtom(disp, "_NET_CURRENT_DESKTOP", False);
}

int         listenClientEvents(void) {
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
            updateWorkspace();
    }
    XCloseDisplay(disp);
    return (0);
}

