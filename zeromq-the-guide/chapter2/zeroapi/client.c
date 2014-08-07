#include <assert.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include <zmq.h>

#include "utils.h"


char *node_id;


int main(int argc, char *argv[])
{
    // verify args
    assert(argc > 2);
    assert(argc % 2 == 1);

    // set node id
    node_id = generate_node_id();
    printf("Node ID: %s\n", node_id);

    // create context and sockets
    void *context = zmq_ctx_new();
    void *pusher = zmq_socket(context, ZMQ_PUSH);
    void *subscriber = zmq_socket(context, ZMQ_SUB);
    zmq_setsockopt(subscriber, ZMQ_SUBSCRIBE, node_id, 8);

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
        // loop variables
        zmq_msg_t send_msg, recv_msg;
        char *msg_str;
        char buffer[30] = {0};
        int msg_len;

        // sleep for a wee bit
        usleep(1000000);

        // setup push message
        snprintf(buffer, sizeof(buffer), "%s:message #%i", node_id, i);
        msg_str = strndup(buffer, sizeof(buffer));
        msg_len = strnlen(msg_str, 50);
        zmq_msg_init_size(&send_msg, msg_len);
        memcpy(&send_msg, msg_str, msg_len);

        // notify console that you are about to send message
        printf("sending: %s (len: %i)\n", msg_str, msg_len);

        // send message
        zmq_msg_send(&send_msg, pusher, 0);

        // cleanup
        free(msg_str);
        zmq_msg_close(&send_msg);

        // setup subscribe message
        zmq_msg_init(&recv_msg);
        zmq_msg_recv(&recv_msg, subscriber, 0);
        msg_len = zmq_msg_size(&recv_msg);
        msg_str = malloc(msg_len + 1);
        memcpy(msg_str, zmq_msg_data(&recv_msg), msg_len);
        msg_str[msg_len] = 0;

        // notify console that you got a message
        printf("got %s\n", msg_str);

        // cleanup
        free(msg_str);
        zmq_msg_close(&recv_msg);
    }

    // cleanup
    IntList_destroy(ports);
    zmq_close(pusher);
    zmq_close(subscriber);
    zmq_ctx_destroy(context);
    free(node_id);

    printf("Stopping client.\n");
    return 0;
}
