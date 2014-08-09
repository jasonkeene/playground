#include <stdio.h>
#include <unistd.h>

#include <zmq.h>

#include "zhelpers.h"


int main(void)
{
    void *context = zmq_ctx_new();

    void *responder = zmq_socket(context, ZMQ_REP);
    zmq_connect(responder, "tcp://localhost:22222");

    while (1) {
        char *string = s_recv(responder);
        printf("Got \"%s\"\n", string);
        free(string);
        sleep(1);
        printf("Sending \"%s\"\n", "World");
        s_send(responder, "World");
    }

    zmq_close(responder);
    zmq_ctx_destroy(context);

    return 0;
}
