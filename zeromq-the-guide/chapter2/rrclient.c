#include <stdio.h>

#include <zmq.h>

#include "zhelpers.h"


int main(void)
{
    void *context = zmq_ctx_new();

    void *requester = zmq_socket(context, ZMQ_REQ);
    zmq_connect(requester, "tcp://localhost:11111");

    for (int i = 0; i < 10; i++) {
        printf("Sending request %i \"%s\"\n", i, "Hello");
        s_send(requester, "Hello");
        char *string = s_recv(requester);
        printf("Got reply %i \"%s\"\n", i, string);
        free(string);
    }

    zmq_close(requester);
    zmq_ctx_destroy(context);

    return 0;
}
