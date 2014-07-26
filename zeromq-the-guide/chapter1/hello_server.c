#include <stdio.h>
#include <assert.h>
#include <zmq.h>
#include "zhelpers.h"


int main(void)
{
    printf("Starting hello world server...\n");
    void *context = zmq_ctx_new();

    // create responder
    void *responder = zmq_socket(context, ZMQ_REP);
    int rc = zmq_bind(responder, "tcp://0.0.0.0:5555");
    assert(rc == 0);

    while (1) {
        char *recv = s_recv(responder);
        printf("Received: %s...\n", recv);
        s_sleep(500);  // do some work
        printf("Sending World\n");
        s_send(responder, "World");
    }
    return 0;
}
