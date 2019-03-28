#ifndef EVENTS_H
#define EVENTS_H

int     sendClientMessage(char *msgType, int index);
int     listenClientEvents(void *widget, void *xutil, void *signals, void *app, void *config);

#endif
