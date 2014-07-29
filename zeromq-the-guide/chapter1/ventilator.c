#include <stdio.h>
#include "zhelpers.h"


int main(void)
{
    void *context = zmq_ctx_new();

    // socket to send messages on
    void *sender = zmq_socket(context, ZMQ_PUSH);
    zmq_bind(sender, "tcp://*:5557");

    // socket to send start of batch message on
    void *sink = zmq_socket(context, ZMQ_PUSH);
    zmq_connect(sink, "tcp://localhost:5558");

    printf("Press enter to start messages...\n");
    getchar();
    printf("Sending tasks to workers.\n");

    // seed random number generator
    srandom((unsigned)time(NULL));

    // signal sink to start counting
    s_send(sink, "0");

    int i;
    int total_msec = 0;
    for (i = 0; i < 100; i++) {
        int workload;
        workload = randof(100) + 1;
        total_msec += workload;
        printf("%d\n", total_msec);
        char string[10];
        snprintf(string, sizeof(string), "%d", workload);
        s_send(sender, string);
    }
    printf("Total expected cost: %d msec\n", total_msec);

    zmq_close(sender);
    zmq_ctx_destroy(context);
    return 0;
}
