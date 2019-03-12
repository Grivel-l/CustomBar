#include <stdio.h>
#include <X11/Xlib.h>

int     createTrayManager(void) {
    Display *disp;
    Window  owner;
    Window  window;
    Atom    trayManager;

    if ((disp = XOpenDisplay(NULL)) == NULL)
        return (1);
    if ((trayManager = XInternAtom(disp, "_NET_SYSTEM_TRAY_S0", False)) == None)
        return (1);
    window = XCreateSimpleWindow(disp, RootWindow(disp, 0), 0, 0, 100, 100, 0, 0, 0);
    if ((owner = XGetSelectionOwner(disp, trayManager)) != None) {
        dprintf(1, "Tray manager already have an owner\n");
        return (2);
    }
    XSetSelectionOwner(disp, trayManager, window, CurrentTime);
    if (XGetSelectionOwner(disp, trayManager) == window) {
        dprintf(1, "Tray owned !\n");
    }
    return (0);
}

