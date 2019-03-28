#include "events.h"
#include <X11/Xlib.h>

extern void updateWorkspace(void *widget, void *xutil, void *signals, void *app, void *config);

static int  listenWorkspaces(Display *disp) {
    if (XSelectInput(disp, RootWindow(disp, 0), PropertyChangeMask) == BadWindow)
        return (None);
    return XInternAtom(disp, "_NET_CURRENT_DESKTOP", False);
}

int         sendClientMessage(char *msgType, int index) {
    Display             *disp;
    XEvent              event;
    XClientMessageEvent clientMessage;

    if ((disp = XOpenDisplay(NULL)) == NULL)
        return (1);
    clientMessage.type = ClientMessage;
    clientMessage.display = disp;
    clientMessage.message_type = XInternAtom(disp, msgType, False);
    clientMessage.send_event = True;
    clientMessage.window = DefaultRootWindow(disp);
    clientMessage.format = 32;
    clientMessage.data.l[0] = index;
    clientMessage.data.l[1] = CurrentTime;
    clientMessage.data.l[2] = 0;
    clientMessage.data.l[3] = 0;
    clientMessage.data.l[4] = 0;
    event.type = ClientMessage;
    event.xclient = clientMessage;
    XSendEvent(disp, DefaultRootWindow(disp), False, SubstructureNotifyMask | SubstructureRedirectMask, &event);
    XFlush(disp);
    return (0);
}

int         listenClientEvents(void *widget, void *xutil, void *signals, void *app, void *config) {
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
            updateWorkspace(widget, xutil, signals, app, config);
    }
    XCloseDisplay(disp);
    return (0);
}

