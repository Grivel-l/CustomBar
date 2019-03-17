#include "./tray.h"

// TODO Better error handling
static xcb_atom_t   getAtom(xcb_connection_t *conn, const char *name) {
    xcb_generic_error_t     *error;
    xcb_intern_atom_reply_t *reply;

    reply = xcb_intern_atom_reply(conn, xcb_intern_atom(conn, 0, strlen(name), name), &error);
    if (error != NULL) {
        dprintf(1, "Couldn't get %s atom\n", name);
        return (1);
    }
    return reply->atom;
}

static xcb_window_t getSelectionOwner(xcb_connection_t *conn, xcb_atom_t trayManager) {
    xcb_generic_error_t             *error;
    xcb_get_selection_owner_reply_t *reply;

    reply = xcb_get_selection_owner_reply(conn, xcb_get_selection_owner(conn, trayManager), &error);
    if (error != NULL) {
        dprintf(1, "Couldn't get tray owner. Error code: %i\n", error->error_code);
        return (XCB_NONE);
    }
    return reply->owner;
}

static xcb_window_t createWindow(xcb_connection_t *conn, xcb_screen_t *screen) {
    xcb_window_t    window;
    uint32_t        value[2];

    window = xcb_generate_id(conn);
    value[0] = screen->black_pixel;
    xcb_create_window(conn, XCB_COPY_FROM_PARENT, window, screen->root, 0, 0, 1, 1, 0, XCB_WINDOW_CLASS_INPUT_OUTPUT, screen->root_visual, XCB_CW_EVENT_MASK | XCB_CW_BACK_PIXEL, value);
    return window;
}

static int  notifySelection(xcb_connection_t *conn, xcb_screen_t *screen,
                xcb_window_t window, xcb_atom_t trayManager) {
    xcb_client_message_event_t  event;
    xcb_atom_t                  manager;

    event.format = 32;
    event.type = getAtom(conn, "MANAGER");
    event.response_type = XCB_CLIENT_MESSAGE;
    event.data.data32[0] = XCB_CURRENT_TIME;
    event.data.data32[1] = trayManager;
    event.data.data32[2] = window;
    event.data.data32[3] = 0;
    event.data.data32[4] = 0;
    xcb_send_event(conn, 0, screen->root, XCB_EVENT_MASK_STRUCTURE_NOTIFY, (const char *)(&event));
}

static int  handleEvent(xcb_connection_t *conn, xcb_client_message_event_t *clientMessage, xcb_atom_t opcode, xcb_window_t window, size_t *i) {
    dprintf(1, "ClientMessage received: Format: %i, AtomType: %i\n", clientMessage->type, clientMessage->type);
    if (clientMessage->format == 32 && clientMessage->type == opcode) {
        dprintf(1, "Requesting dock\n");
        xcb_reparent_window(conn, clientMessage->data.data32[2], window, *i * 20, 0);
        xcb_configure_window(conn, clientMessage->data.data32[2], XCB_CONFIG_WINDOW_WIDTH | XCB_CONFIG_WINDOW_HEIGHT, (uint32_t[]){20, 20});
        xcb_configure_window(conn, window, XCB_CONFIG_WINDOW_WIDTH | XCB_CONFIG_WINDOW_HEIGHT, (uint32_t[]){20 * (*i + 1), 20});
        xcb_map_window(conn, clientMessage->data.data32[2]);
        xcb_flush(conn);
        *i += 1;
    }
    return (0);
}

int     createTrayManager(void) {
    size_t              i;
    xcb_connection_t    *conn;
    xcb_generic_event_t *event;
    xcb_window_t        window;
    xcb_atom_t          opcode;
    xcb_screen_t        *screen;
    xcb_atom_t          trayManager;
    xcb_client_message_event_t  *clientMessage;

    conn = xcb_connect(NULL, NULL);
    opcode = getAtom(conn, "_NET_SYSTEM_TRAY_OPCODE");
    trayManager = getAtom(conn, "_NET_SYSTEM_TRAY_S0");
    screen = xcb_setup_roots_iterator(xcb_get_setup(conn)).data;
    window = createWindow(conn, screen);
    xcb_atom_t yo;
    yo = getAtom(conn, "_NET_WM_STATE_MODAL");
    xcb_change_property(conn, XCB_PROP_MODE_APPEND, window, getAtom(conn, "_NET_WM_STATE"), XCB_ATOM_ATOM, 32, 1, (const void *)(&yo));
    yo = getAtom(conn, "_NET_WM_STATE_SKIP_TASKBAR");
    xcb_change_property(conn, XCB_PROP_MODE_APPEND, window, getAtom(conn, "_NET_WM_STATE"), XCB_ATOM_ATOM, 32, 1, (const void *)(&yo));
    yo = getAtom(conn, "_NET_WM_ACTION_ABOVE");
    xcb_change_property(conn, XCB_PROP_MODE_REPLACE, window, getAtom(conn, "_NET_WM_ALLOWED_ACTIONS"), XCB_ATOM_WM_HINTS, 32, 1, (const void *)(&yo));
    if (getSelectionOwner(conn, trayManager) != XCB_NONE) {
        dprintf(1, "Tray already have an owner\n");
    }
    xcb_set_selection_owner(conn, window, trayManager, XCB_CURRENT_TIME);
    if (getSelectionOwner(conn, trayManager) == window) {
        dprintf(1, "Tray successfully owned\n");
        notifySelection(conn, screen, window, trayManager);
    }
    xcb_map_window(conn, window);
    xcb_flush(conn);
    dprintf(1, "Listening for events...\n");
    i = 0;
    while ((event = xcb_wait_for_event(conn)) != NULL) {
        if (XCB_EVENT_RESPONSE_TYPE(event) == XCB_CLIENT_MESSAGE) {
            handleEvent(conn, (xcb_client_message_event_t *)event, opcode, window, &i);
        }
        free(event);
    }
    return (0);
}

