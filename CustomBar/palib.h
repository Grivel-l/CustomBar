#ifndef PALIB_H
#define PALIB_H

#include <stdlib.h>
#include <stdio.h>
#include <pulse/pulseaudio.h>

void    destroy_con(void);
void    *update_volume(void *ctx, int increase);
void    *create_con(char *appName, void *signals, char *volumeIcon);

#endif
