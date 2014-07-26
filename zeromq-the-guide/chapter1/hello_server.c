#include <stdio.h>
#include <unistd.h>
#include <assert.h>
#include <zmq.h>


int main(void)
{
    printf("Starting hello world server...\n");
    void *context = zmq_ctx_new();

    // create responder
    void *responder = zmq_socket(context, ZMQ_REP);
    int rc = zmq_bind(responder, "tcp://0.0.0.0:5555");
    assert(rc == 0);

    while (1) {
        char buffer[10];
        zmq_recv(responder, buffer, 10, 0);
        printf("Received: %s...\n", buffer);
        sleep(1);  // do some work
        printf("Sending World\n");
        zmq_send(responder, "World", 5, 0);
    }
    return 0;
}
