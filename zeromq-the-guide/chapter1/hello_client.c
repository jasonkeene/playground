#include <stdio.h>
#include <zmq.h>
#include "zhelpers.h"


int main(void)
{
    printf("Connecting to hello world server...\n");
    void *context = zmq_ctx_new();

    // create requester
    void *requester = zmq_socket(context, ZMQ_REQ);
    zmq_connect(requester, "tcp://localhost:5555");

    for (int i = 0; i < 10; i++) {
        printf("Sending Hello %d...\n", i);
        s_send(requester, "Hello");
        char *recv = s_recv(requester);
        printf("Received %s\n", recv);
    }
    zmq_close(requester);
    zmq_ctx_destroy(context);

    return 0;
}
