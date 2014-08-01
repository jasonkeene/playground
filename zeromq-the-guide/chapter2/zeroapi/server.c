#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <zmq.h>

#include "utils.h"


int main(int argc, char *argv[])
{
    // verify args
    assert(argc > 2);

    // create context and sockets
    void *context = zmq_ctx_new();
    void *puller = zmq_socket(context, ZMQ_PULL);
    void *publisher = zmq_socket(context, ZMQ_PUB);

    // get ports list
    IntList *ports = get_ports(argc, argv);

    // connect sockets
    printf("Starting puller on port %i.\n", ports->head->data);
    char *connect_str = connection_str(ports->head->data);
    zmq_bind(puller, connect_str);
    free(connect_str);

    printf("Starting publisher on port %i.\n", ports->head->next->data);
    connect_str = connection_str(ports->head->next->data);
    zmq_bind(publisher, connect_str);
    free(connect_str);

    for (int i = 0; ; i++) {
        printf("%i\n", i);

        zmq_msg_t msg;
        zmq_msg_init(&msg);
        zmq_msg_recv(&msg, puller, 0);
        int size = zmq_msg_size(&msg);
        char *string = malloc(size + 1);
        memcpy(string, zmq_msg_data(&msg), size);
        zmq_msg_close(&msg);
        string[size] = 0;

        printf("%s", string);
    }

    // cleanup
    IntList_destroy(ports);
    zmq_ctx_destroy(context);

    printf("Stopping server.\n");
    return 0;
}
