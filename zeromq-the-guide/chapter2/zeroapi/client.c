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
        // sleep for a wee bit
        usleep(1000000);

        // push message
        zmq_msg_t send_msg;
        char buffer[30];
        snprintf(buffer, sizeof(buffer), "message #%i", i);
        char *msg_str = strndup(buffer, sizeof(buffer));
        printf("sending: %s\n", msg_str);
        int msg_len = strnlen(msg_str, 50);
        zmq_msg_init_data(&send_msg, msg_str, msg_len, NULL, NULL);
        zmq_msg_send(&send_msg, pusher, 0);
    }
    while (1) {
        // get and display message
        zmq_msg_t recv_msg;
        zmq_msg_init(&recv_msg);
        zmq_msg_recv(&recv_msg, subscriber, 0);
        int msg_len = zmq_msg_size(&recv_msg);
        char *string = malloc(msg_len + 1);
        memcpy(string, zmq_msg_data(&recv_msg), msg_len);
        string[msg_len] = 0;
        printf("got %s\n", string);
        zmq_msg_close(&recv_msg);
    }

    // cleanup
    IntList_destroy(ports);
    zmq_close(pusher);
    zmq_close(subscriber);
    zmq_ctx_destroy(context);

    printf("Stopping client.\n");
    return 0;
}
