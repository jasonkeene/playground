#include <stdio.h>
#include <assert.h>

#include <zmq.h>

#include "utils.h"


int main(int argc, char *argv[])
{
    assert(argc > 2);
    PortList *port_list = get_port_list(argc, argv);
    printf("Starting puller on port %i.\n", port_list->head->port);
    printf("Starting publisher on port %i.\n", port_list->head->next->port);
    PortList_destroy(port_list);
    printf("Stopping server.\n");
    return 0;
}
