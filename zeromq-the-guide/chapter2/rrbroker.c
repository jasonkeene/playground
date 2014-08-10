#include <stdio.h>

#include <zmq.h>


int main(void)
{
    void *context = zmq_ctx_new();

    void *frontend = zmq_socket(context, ZMQ_ROUTER);
    void *backend = zmq_socket(context, ZMQ_DEALER);
    zmq_bind(frontend, "tcp://0.0.0.0:11111");
    zmq_bind(backend, "tcp://0.0.0.0:22222");

    zmq_proxy(frontend, backend, NULL);

    zmq_close(frontend);
    zmq_close(backend);
    zmq_ctx_destroy(context);

    return 0;
}
