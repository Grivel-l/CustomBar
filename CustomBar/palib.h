#ifndef PALIB_H
#define PALIB_H

#include <stdlib.h>
#include <stdio.h>
#include <pulse/pulseaudio.h>

void    destroy_con(void);
void    *create_con(char *appName, void *signals);

#endif
