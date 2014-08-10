#include <stdio.h>

#include <zmq.h>


int main(void)
{
    void *context = zmq_ctx_new();

    void *frontend = zmq_socket(context, ZMQ_ROUTER);
    void *backend = zmq_socket(context, ZMQ_DEALER);
    zmq_bind(frontend, "tcp://0.0.0.0:11111");
    zmq_bind(backend, "tcp://0.0.0.0:22222");

    zmq_pollitem_t items[] = {
        {frontend, 0, ZMQ_POLLIN, 0},
        {backend, 0, ZMQ_POLLIN, 0},
    };

    while (1) {
        zmq_msg_t message;
        zmq_poll(items, 2, -1);

        if (items[0].revents & ZMQ_POLLIN) {
            while (1) {
                zmq_msg_init(&message);
                zmq_msg_recv(&message, frontend, 0);
                int more = zmq_msg_more(&message);
                zmq_msg_send(&message, backend, more ? ZMQ_SNDMORE : 0);
                zmq_msg_close(&message);
                if (!more) break;
            }
        }
        if (items[1].revents & ZMQ_POLLIN) {
            while (1) {
                zmq_msg_init(&message);
                zmq_msg_recv(&message, backend, 0);
                int more = zmq_msg_more(&message);
                zmq_msg_send(&message, frontend, more ? ZMQ_SNDMORE : 0);
                zmq_msg_close(&message);
                if (!more) break;
            }
        }
    }

    zmq_close(frontend);
    zmq_close(backend);
    zmq_ctx_destroy(context);

    return 0;
}
