#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "utils.h"


PortList *PortList_create()
{
    PortList *port_list = malloc(sizeof(PortList));
    port_list->head = NULL;
    return port_list;
}


void PortList_push(PortList *port_list, int port)
{
    PortNode *new_node = malloc(sizeof(PortNode));
    if (port_list->head == NULL) {
        port_list->head = new_node;
        return;
    } else {
        PortNode *node = port_list->head;
        while (node->next != NULL) {
            node = node->next;
        }
        node->next = new_node;
    }
}


void PortList_destroy(PortList *port_list)
{
    PortNode *node = port_list->head;
    while (node != NULL) {
        PortNode *next = node->next;
        free(node);
        node = next;
    }
    free(port_list);
}


PortList *get_port_list(int argc, char *argv[])
{
    PortList *port_list = PortList_create();
    for (int i = 1; i < argc; i++) {
        int port = atoi(argv[i]);
        port = port < 8000 ? arc4random() % 57536 + 8000 : port;
        PortList_push(port_list, port);
    }
    return port_list;
}


char *pusher_str(int port)
{
    char buffer[40];
    snprintf(buffer, sizeof(buffer), "tcp://127.0.0.1:%i", port);
    return strndup(buffer, sizeof(buffer));
}
