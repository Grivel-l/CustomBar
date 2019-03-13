#include "./tray.h"
#include <X11/Xatom.h>

// TODO RootWindow(disp, 0) should send to RootWindow(disp, DefaultScreen())
int     createTrayManager(void) {
    Display             *disp;
    XEvent              event;
    Window              window;
    Atom                trayManager;

    if ((disp = XOpenDisplay(NULL)) == NULL)
        return (1);
    if ((trayManager = XInternAtom(disp, "_NET_SYSTEM_TRAY_S0", False)) == None)
        return (1);
    window = XCreateSimpleWindow(disp, RootWindow(disp, 0), 0, 0, 1, 1, 0, XBlackPixel(disp, 0), XBlackPixel(disp, 0));
    if (XGetSelectionOwner(disp, trayManager) != None) {
        dprintf(1, "Tray manager already have an owner\n");
        return (2);
    }
    XSetSelectionOwner(disp, trayManager, window, CurrentTime);
    if (XGetSelectionOwner(disp, trayManager) == window) {
        dprintf(1, "Tray owned !\n");
        event.xclient.format = 32;
        event.xclient.message_type = XA_RESOURCE_MANAGER;
        event.xclient.data.l[0] = CurrentTime;
        event.xclient.data.l[1] = trayManager;
        event.xclient.data.l[2] = window;
        XSendEvent(disp, RootWindow(disp, 0), False, StructureNotifyMask, &event);
    }
    if (XMapWindow(disp, window) == BadWindow) {
        return (1);
    }
    while (1) {
        XNextEvent(disp, &event);
        dprintf(1, "Event received by tray\n");
    }
    return (0);
}

