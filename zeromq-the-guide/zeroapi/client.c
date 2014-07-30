#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

#include <zmq.h>

#include "utils.h"


int main(int argc, char *argv[])
{
    // verify args
    assert(argc > 2);

    // create context and sockets
    void *context = zmq_ctx_new();
    void *pusher = zmq_socket(context, ZMQ_PUSH);
    void *subscriber = zmq_socket(context, ZMQ_SUB);

    // get ports list
    PortList *port_list = get_port_list(argc, argv);

    // connect pusher
    printf("Starting client on port %i.\n", port_list->head->port);
    char *connect_str = pusher_str(port_list->head->port);
    zmq_connect(pusher, connect_str);
    free(connect_str);

    // connect subscriber
    /* char *connect_str = _str(port); */
    /* zmq_connect(pusher, connect_str); */
    /* free(connect_str); */

    zmq_close(pusher);
    zmq_close(subscriber);
    zmq_ctx_destroy(context);

    printf("Stopping client.\n");
    return 0;
}
