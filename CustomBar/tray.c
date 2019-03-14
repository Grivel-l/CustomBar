#include "./tray.h"
#include <X11/Xatom.h>
#include <unistd.h>

static int  sendEvent(Display *disp, Window window, Atom trayManager) {
    XEvent  event;
    Atom    manager;

    if ((manager = XInternAtom(disp, "MANAGER", False)) == None)
        return (1);
    event.xclient.type = ClientMessage;
    event.xclient.display = disp;
    event.xclient.window = window;
    event.xclient.format = 32;
    event.xclient.message_type = manager;
    event.xclient.data.l[0] = CurrentTime;
    event.xclient.data.l[1] = trayManager;
    event.xclient.data.l[2] = window;
    XSendEvent(disp, RootWindow(disp, 0), False, StructureNotifyMask, &event);
    return (0);
}

static int  handleEvent(Display *disp, XEvent event, Window window) {
    int     ret;
    Atom    opcode;

    if ((opcode = XInternAtom(disp, "_NET_SYSTEM_TRAY_OPCODE", False)) == None)
        return (1);
    if (event.type == ClientMessage) {
        if (event.xclient.message_type == opcode && event.xclient.format == 32) {
            dprintf(1, "Opcode received\n");
            if ((int)(event.xclient.data.l[1]) == SYSTEM_TRAY_REQUEST_DOCK) {
                dprintf(1, "Requesting dock\n");
                ret = XReparentWindow(disp, (Window)(event.xclient.data.l[2]), window, 0, 0);
                if (ret == BadMatch || ret == BadWindow) {
                    dprintf(1, "Error during docking\n");
                    return (1);
                }
                XMoveResizeWindow(disp, (Window)(event.xclient.data.l[2]), 0, 0, 100, 100);
                XMapRaised(disp, (Window)(event.xclient.data.l[2]));
                XFlush(disp);
                /* XUnmapWindow(disp, (Window)(event.xclient.data.l[2])); */
                /* XReparentWindow(disp, (Window)(event.xclient.data.l[2]), RootWindow(disp, 0), 0, 0); */
                /* XDestroyWindow(disp, (Window)(event.xclient.data.l[2])); */
            } else if ((int)(event.xclient.data.l[1]) == SYSTEM_TRAY_BEGIN_MESSAGE) {
                dprintf(1, "Beginning message\n");
            } else if ((int)(event.xclient.data.l[1]) == SYSTEM_TRAY_CANCEL_MESSAGE) {
                dprintf(1, "Cancelling message\n");
            }
        }
        dprintf(1, "Client message received\n");
        dprintf(1, "Atom name: %s\n", XGetAtomName(disp, event.xclient.message_type));
    }
    dprintf(1, "Event received by tray %i\n", event.type);
    return (0);
}

// TODO RootWindow(disp, 0) should send to RootWindow(disp, DefaultScreen())
int     createTrayManager(void) {
    Display             *disp;
    XEvent              event;
    Window              window;
    Atom                trayManager;
    Time                timestamp;

    if ((disp = XOpenDisplay(NULL)) == NULL)
        return (1);
    if ((trayManager = XInternAtom(disp, "_NET_SYSTEM_TRAY_S0", False)) == None)
        return (1);
    window = XCreateSimpleWindow(disp, RootWindow(disp, 0), 0, 0, 1, 1, 0, XBlackPixel(disp, 0), XBlackPixel(disp, 0));
    if (XGetSelectionOwner(disp, trayManager) != None) {
        dprintf(1, "Tray manager already have an owner\n");
        return (2);
    }
    timestamp = CurrentTime;
    XSetSelectionOwner(disp, trayManager, window, timestamp);
    XSelectInput(disp, window, NoEventMask);
    if (XGetSelectionOwner(disp, trayManager) == window) {
        dprintf(1, "Tray owned !\n");
        if (sendEvent(disp, window, trayManager) == 1)
            return (1);
    }
    if (XMapWindow(disp, window) == BadWindow) {
        return (1);
    }
    while (1) {
        XNextEvent(disp, &event);
        if (handleEvent(disp, event, window) == 1)
            return (1);
    }
    XSetSelectionOwner(disp, trayManager, None, timestamp);
    XUnmapWindow(disp, window);
    XDestroyWindow(disp, window);
    XCloseDisplay(disp);
    return (0);
}

