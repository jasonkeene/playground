#include <stdlib.h>


int get_port(int argc, char *argv[])
{
    int port = 0;
    if (argc == 2) {
        port = atoi(argv[1]);
    }
    if (port < 8000) {
        port = arc4random() % 57536 + 8000;
    }
    return port;
}
