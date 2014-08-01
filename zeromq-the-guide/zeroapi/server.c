#include <stdio.h>
#include <assert.h>

#include <zmq.h>

#include "utils.h"


int main(int argc, char *argv[])
{
    assert(argc > 2);
    IntList *ports = get_ports(argc, argv);
    printf("Starting puller on port %i.\n", ports->head->data);
    printf("Starting publisher on port %i.\n", ports->head->next->data);
    IntList_destroy(ports);
    printf("Stopping server.\n");
    return 0;
}
