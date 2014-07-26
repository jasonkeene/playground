#include <stdio.h>
#include <zmq.h>


int main(void)
{
    printf("Connecting to hello world server...\n");
    void *context = zmq_ctx_new();

    // create requester
    void *requester = zmq_socket(context, ZMQ_REQ);
    zmq_connect(requester, "tcp://localhost:5555");

    for (int i = 0; i < 10; i++) {
        char buffer[10];
        printf("Sending Hello %d...\n", i);
        zmq_send(requester, "Hello", 5, 0);
        zmq_recv(requester, buffer, 10, 0);
        printf("Received %s\n", buffer);
    }
    zmq_close(requester);
    zmq_ctx_destroy(context);

    return 0;
}
