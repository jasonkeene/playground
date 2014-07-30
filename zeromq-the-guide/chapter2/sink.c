#include <stdio.h>
#include "zhelpers.h"


int main(int argc, char *argv[])
{
    void *context = zmq_ctx_new();

    // socket to read messages from
    void *puller = zmq_socket(context, ZMQ_PULL);
    char buffer[30];
    snprintf(buffer, sizeof(buffer), "tcp://*:%s", argv[1]);
    printf("Binding to: %s\n", buffer);
    zmq_bind(puller, buffer);

    for (int i = 0; ; i++) {
        char *string = s_recv(puller);
        free(string);
        printf("%d\n", i);
        fflush(stdout);
    }

    zmq_close(puller);
    zmq_ctx_destroy(context);

    return 0;
}
