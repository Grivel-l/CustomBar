#include "palib.h"

extern void set_volume(int volume, void *config);

pa_mainloop *loop;

void            destroy_con(void) {
    pa_mainloop_quit(loop, 0);
    pa_mainloop_free(loop);
}

static  void    cb_infos(pa_context *c, const pa_sink_info *infos, int eol, void *userData) {
    (void)c;
    if (eol == 1)
        return ;
    set_volume((int)((float)infos->volume.values[1] / (float)PA_VOLUME_NORM * 100), userData);
    destroy_con();
}

static void     cb(pa_context *c, void *userData) {
    if (pa_context_get_state(c) == PA_CONTEXT_READY) {
        pa_context_get_sink_info_list(c, &cb_infos, userData);
    }
}

int             create_con(char *appName, void *config) {
    pa_context  *ctx;

    if ((loop = pa_mainloop_new()) == NULL)
        return (1);
    if ((ctx = pa_context_new(pa_mainloop_get_api(loop), appName)) == NULL) {
        pa_mainloop_free(loop);
        return (1);
    }
    pa_context_set_state_callback(ctx, &cb, config);
    if (pa_context_connect(ctx, NULL, PA_CONTEXT_NOFLAGS, NULL) < 0) {
        pa_mainloop_free(loop);
        return (1);
    }
    if (pa_mainloop_run(loop, 0) < 0) {
        pa_mainloop_free(loop);
        return (1);
    }
    return (0);
}

