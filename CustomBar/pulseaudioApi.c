#include "palib.h"

extern void set_volume(int volume, void *config);

pa_threaded_mainloop *loop;

void            destroy_con(void) {
    pa_threaded_mainloop_stop(loop);
    pa_threaded_mainloop_free(loop);
}

static  void    cb_infos(pa_context *c, const pa_sink_info *infos, int eol, void *userData) {
    if (eol == 1)
        return ;
    set_volume((int)((float)infos->volume.values[1] / (float)PA_VOLUME_NORM * 100), userData);
}

static void     event_cb(pa_context *c, pa_subscription_event_type_t type, uint32_t idx, void *userData) {
    pa_context_get_sink_info_list(c, &cb_infos, userData);
}

static void     event_success_cb(pa_context *c, int success, void *userData) {
    pa_context_set_subscribe_callback(c, &event_cb, userData);
}

static void     cb(pa_context *c, void *userData) {
    if (pa_context_get_state(c) == PA_CONTEXT_READY) {
        pa_context_subscribe(c, PA_SUBSCRIPTION_MASK_SINK, &event_success_cb, userData);
        pa_context_get_sink_info_list(c, &cb_infos, userData);
    }
}

int             create_con(char *appName, void *config) {
    pa_context  *ctx;

    if ((loop = pa_threaded_mainloop_new()) == NULL)
        return (1);
    if ((ctx = pa_context_new(pa_threaded_mainloop_get_api(loop), appName)) == NULL) {
        pa_threaded_mainloop_free(loop);
        return (1);
    }
    pa_context_set_state_callback(ctx, &cb, config);
    if (pa_context_connect(ctx, NULL, PA_CONTEXT_NOFLAGS, NULL) < 0) {
        pa_threaded_mainloop_free(loop);
        return (1);
    }
    if (pa_threaded_mainloop_start(loop) < 0) {
        pa_threaded_mainloop_free(loop);
        return (1);
    }
    return (0);
}

