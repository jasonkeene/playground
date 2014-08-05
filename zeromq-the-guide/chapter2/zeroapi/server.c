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
        // pull and display message
        zmq_msg_t recv_msg;
        zmq_msg_init(&recv_msg);
        zmq_msg_recv(&recv_msg, puller, 0);

        int msg_len = zmq_msg_size(&recv_msg);
        char *msg_str = malloc(msg_len + 1);
        memcpy(msg_str, zmq_msg_data(&recv_msg), msg_len);
        msg_str[msg_len] = 0;

        printf("got %s\n", msg_str);

        zmq_msg_close(&recv_msg);

        // publish message
        zmq_msg_t send_msg;

        printf("sending: %s (len: %i)\n", msg_str, msg_len);

        zmq_msg_init_size(&send_msg, msg_len);
        memcpy(&send_msg, msg_str, msg_len);
        zmq_msg_send(&send_msg, publisher, 0);
    }

    // cleanup
    IntList_destroy(ports);
    zmq_close(puller);
    zmq_close(publisher);
    zmq_ctx_destroy(context);

    printf("Stopping server.\n");
    return 0;
}
