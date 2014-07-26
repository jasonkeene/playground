#include <stdio.h>
#include <zmq.h>
#include <assert.h>
#include "zhelpers.h"
#include "weather.h"


#define PORT "5556"


int main(void)
{
    printf("Starting weather server on port " PORT "\n");
    void *context = zmq_ctx_new();
    void *publisher = zmq_socket(context, ZMQ_PUB);
    int rc = zmq_bind(publisher, "tcp://*:"PORT);
    assert(rc == 0);
    /* rc = zmq_bind(publisher, "ipc://weather.ipc"); */
    /* assert(rc == 0); */

    // init random number generator
    while (1) {
        //
        Weather *weather = Weather_create_fake();
        printf("sending: %s\n", weather->str);
        s_send(publisher, weather->str);
        Weather_destroy(weather);
        s_sleep(10);
    }

    zmq_close(publisher);
    zmq_ctx_destroy(context);
    return 0;
}
