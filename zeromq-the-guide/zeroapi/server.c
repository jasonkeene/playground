#include <stdio.h>

#include <zmq.h>

#include "utils.h"


int main(int argc, char *argv[])
{
    int port = get_port(argc, argv);
    printf("Starting server on port %i.\n", port);
    printf("Stopping server.\n");
    return 0;
}
