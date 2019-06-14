#include "palib.h"
#include  <stdio.h>

extern void setVolume(int volume, void *signals, void *volumeIcon);

/* void            destroy_con(void) { */
/*     pa_threaded_mainloop_stop(loop); */
/*     pa_threaded_mainloop_free(loop); */
/* free(params); */
/* } */

static  void    cb_infos(pa_context *c, const pa_sink_info *infos, int eol, void *userData) {
    void    **params;

    params = userData;
    if (eol == 1) {
        return ;
    }
    setVolume((int)((float)infos->volume.values[0] / (float)PA_VOLUME_NORM * 100), params[0], params[1]);
}

static void     event_cb(pa_context *c, pa_subscription_event_type_t type, uint32_t idx, void *userData) {
    pa_context_get_sink_info_by_index(c, 0, cb_infos, userData);
}

static void     event_success_cb(pa_context *c, int success, void *userData) {
    pa_threaded_mainloop_signal(userData, 0);
}

static void     state_cb(pa_context *c, void *userData) {
    if (pa_context_get_state(c) == PA_CONTEXT_READY) {
        pa_threaded_mainloop_signal(userData, 0);
    }
}

static  void    set_cb_infos(pa_context *ctx, const pa_sink_info *infos, int eol, void *userData) {
    int         i;
    pa_cvolume  volume;

    if (eol == 1) {
        return ;
    }
    i = 0;
    if ((int)((uintptr_t)userData) > 0) {
        while (i < infos->volume.channels) {
            volume.values[i] = infos->volume.values[i] + (PA_VOLUME_NORM / 100 * 2);
            i += 1;
        }
    } else {
        if ((PA_VOLUME_NORM / 100 * 2) > infos->volume.values[0]) {
            volume.values[0] = 0;
            volume.values[1] = 0;
        } else {
            while (i < infos->volume.channels) {
                volume.values[i] = infos->volume.values[i] - (PA_VOLUME_NORM / 100 * 2);
                i += 1;
            }
        }
    }
    volume.channels = infos->volume.channels;
    pa_context_set_sink_volume_by_index(ctx, infos->index, &(volume), NULL, NULL);
}

void            *update_volume(void *ctxP, int increase) {
    pa_context  *ctx;

    ctx = (pa_context *)ctxP;
    pa_context_get_sink_info_by_index(ctx, 0, set_cb_infos, (void *)((uintptr_t)increase));
}

void            *create_con(char *appName, void *signals, char *volumeIcon) {
    pa_operation    *op;
    pa_context      *ctx;
    void            **params;
    pa_threaded_mainloop *loop;

    if ((params = malloc(sizeof(void *) * 2)) == NULL)
        return (NULL);
    if ((loop = pa_threaded_mainloop_new()) == NULL)
        return (NULL);
    pa_threaded_mainloop_lock(loop);
    if ((ctx = pa_context_new(pa_threaded_mainloop_get_api(loop), appName)) == NULL) {
        pa_threaded_mainloop_unlock(loop);
        pa_threaded_mainloop_free(loop);
        return (NULL);
    }
    pa_context_set_state_callback(ctx, &state_cb, loop);
    if (pa_context_connect(ctx, NULL, PA_CONTEXT_NOFLAGS, NULL) < 0) {
        pa_threaded_mainloop_unlock(loop);
        pa_threaded_mainloop_free(loop);
        return (NULL);
    }
    if (pa_threaded_mainloop_start(loop) < 0) {
        pa_threaded_mainloop_unlock(loop);
        pa_threaded_mainloop_free(loop);
        return (NULL);
    }
    pa_threaded_mainloop_wait(loop);
    op = pa_context_subscribe(ctx, PA_SUBSCRIPTION_MASK_SINK, event_success_cb, loop);
    while (pa_operation_get_state(op) == PA_OPERATION_RUNNING)
        pa_threaded_mainloop_wait(loop);
    // Check error here
    params[0] = signals;
    params[1] = volumeIcon;
    pa_context_set_subscribe_callback(ctx, event_cb, (void *)params);
    pa_threaded_mainloop_unlock(loop);
    pa_context_get_sink_info_by_index(ctx, 0, cb_infos, (void *)params);
    return (ctx);
}

