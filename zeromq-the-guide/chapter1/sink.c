#include <stdio.h>
#include "zhelpers.h"


int main(void)
{
    void *context = zmq_ctx_new();

    void *puller = zmq_socket(context, ZMQ_PULL);
    zmq_bind(puller, "tcp://*:5558");

    char *start = s_recv(puller);
    free(start);

    int64_t start_time = s_clock();

    int i;
    for (i = 0; i < 100; i++) {
        char *string = s_recv(puller);
        free(string);
        printf((i / 10) * 10 == i ? ":" : ".");
        fflush(stdout);
    }
    printf("Total elapsed time: %d msec\n", (int)(s_clock() - start_time));

    zmq_close(puller);
    zmq_ctx_destroy(context);
    return 0;
}
