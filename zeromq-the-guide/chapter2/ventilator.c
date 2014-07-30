#include <stdio.h>
#include "zhelpers.h"


int main(int argc, char *argv[])
{
    void *context = zmq_ctx_new();

    // socket to send messages on
    void *pusher = zmq_socket(context, ZMQ_PUSH);
    char buffer[30];
    snprintf(buffer, sizeof(buffer), "tcp://*:%s", argv[1]);
    printf("Binding to: %s\n", buffer);
    zmq_bind(pusher, buffer);

    printf("Sending tasks to workers.\n");

    // seed random number generator
    srandom((unsigned)time(NULL));

    for (int i = 0; i < 10000; i++) {
        printf("Sending %i\n", i);
        int workload;
        workload = randof(10) + 1;
        char string[10];
        snprintf(string, sizeof(string), "%d", workload);
        s_send(pusher, string);
        s_sleep(2);
    }

    zmq_close(pusher);
    zmq_ctx_destroy(context);

    return 0;
}
