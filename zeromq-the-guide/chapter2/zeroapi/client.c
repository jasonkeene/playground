#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

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

    // connect sockets
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

    for (int i = 0; i < 10; i++) {
        printf("%i\n", i);
        usleep(1000000);
        char *msg_str = "asdf";
        int msg_len = strnlen(msg_str, 50);
        zmq_msg_t msg;
        assert(zmq_msg_init_size(&msg, msg_len) == 0);
        memcpy(&msg, &msg_str, msg_len);
        zmq_msg_send(&msg, pusher, 0);
        zmq_msg_close(&msg);
        /* free(msg_str); */
    }

    // cleanup
    IntList_destroy(ports);
    zmq_close(pusher);
    zmq_close(subscriber);
    zmq_ctx_destroy(context);

    printf("Stopping client.\n");
    return 0;
}
