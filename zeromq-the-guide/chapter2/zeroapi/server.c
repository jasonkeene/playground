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

    while (1) {
        zmq_msg_t msg;
        zmq_msg_t msg2;
        zmq_msg_init(&msg);

        // get and display message
        zmq_msg_recv(&msg, puller, 0);
        int size = zmq_msg_size(&msg);
        char *string = malloc(size + 1);
        memcpy(string, zmq_msg_data(&msg), size);
        string[size] = 0;
        printf("got %s\n", string);

        // publish message
        zmq_msg_init_data(&msg2, string, strlen(string), NULL, NULL);
        zmq_msg_send(&msg2, publisher, 0);
        zmq_msg_close(&msg);
        zmq_msg_close(&msg2);
    }

    // cleanup
    IntList_destroy(ports);
    zmq_ctx_destroy(context);

    printf("Stopping server.\n");
    return 0;
}
