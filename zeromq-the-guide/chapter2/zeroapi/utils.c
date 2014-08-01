#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "utils.h"


IntList *IntList_create()
{
    IntList *int_list = malloc(sizeof(IntList));
    int_list->head = NULL;
    return int_list;
}


void IntList_push(IntList *int_list, int data)
{
    IntNode *new_node = malloc(sizeof(IntNode));
    new_node->data = data;
    new_node->next = NULL;
    if (int_list->head == NULL) {
        int_list->head = new_node;
    } else {
        IntNode *node = int_list->head;
        while (node->next != NULL) {
            node = node->next;
        }
        node->next = new_node;
    }
}


void IntList_print(IntList *int_list)
{
    IntNode *node = int_list->head;
    printf("[");
    while (node != NULL) {
        printf("%d", node->data);
        node = node->next;
        if (node != NULL) {
            printf(", ");
        }
    }
    printf("]\n");
}


void IntList_destroy(IntList *int_list)
{
    IntNode *node = int_list->head;
    while (node != NULL) {
        IntNode *next = node->next;
        free(node);
        node = next;
    }
    free(int_list);
}


IntList *get_ports(int argc, char *argv[])
{
    IntList *int_list = IntList_create();
    for (int i = 1; i < argc; i++) {
        int port = atoi(argv[i]);
        port = port < 8000 ? arc4random() % 57536 + 8000 : port;
        IntList_push(int_list, port);
    }
    return int_list;
}


char *connection_str(int port)
{
    char buffer[40];
    snprintf(buffer, sizeof(buffer), "tcp://127.0.0.1:%i", port);
    return strndup(buffer, sizeof(buffer));
}
