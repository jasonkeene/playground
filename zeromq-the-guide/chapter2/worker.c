#include <stdio.h>
#include "zhelpers.h"


int main(void)
{
    void *context = zmq_ctx_new();

    // socket to receive jobs from
    void *puller = zmq_socket(context, ZMQ_PULL);
    zmq_connect(puller, "tcp://127.0.0.1:5556");
    zmq_connect(puller, "tcp://127.0.0.1:5557");

    // socket to send jobs to sink
    void *pusher = zmq_socket(context, ZMQ_PUSH);
    zmq_connect(pusher, "tcp://127.0.0.1:5558");
    zmq_connect(pusher, "tcp://127.0.0.1:5559");

    for (int i = 0; ; i++) {
        char *string = s_recv(puller);
        printf("%d\n", i);
        fflush(stdout);
        s_sleep(atoi(string));
        free(string);
        s_send(pusher, "");
    }

    zmq_close(puller);
    zmq_close(pusher);
    zmq_ctx_destroy(context);

    return 0;
}
