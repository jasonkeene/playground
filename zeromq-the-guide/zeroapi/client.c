#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

#include <zmq.h>

#include "utils.h"


int main(int argc, char *argv[])
{
    // verify args
    assert(argc > 2);
    assert(argc % 2 == 1);

    // create context and sockets
    void *context = zmq_ctx_new();
    void *pusher = zmq_socket(context, ZMQ_PUSH);
    void *subscriber = zmq_socket(context, ZMQ_SUB);

    // get ports list
    IntList *ports = get_ports(argc, argv);
    IntList_print(ports);

    // connect pusher
    IntNode *port = ports->head;
    for (int i = 0; port != NULL; i++) {
        char *connect_str = connection_str(port->data);
        if (i % 2 == 0) {
            printf("Connecting pusher to port %i.\n", port->data);
            zmq_connect(pusher, connect_str);
        } else {
            printf("Connecting subscriber to port %i.\n", port->data);
            zmq_connect(subscriber, connect_str);
        }
        free(connect_str);
        port = port->next;
    }

    IntList_destroy(ports);

    zmq_close(pusher);
    zmq_close(subscriber);
    zmq_ctx_destroy(context);

    printf("Stopping client.\n");
    return 0;
}
