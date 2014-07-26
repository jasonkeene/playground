#include <stdio.h>
#include <zmq.h>
#include <assert.h>
#include "zhelpers.h"


#define PORT "5556"
#define DEFAULT_ZIP "10001"


char *make_filter(int argc, char *argv[])
{
    if (argc > 1) {
        char buffer[15] = {"Weather{\0"};
        strncat(buffer, argv[1], 5);
        return strndup(buffer, sizeof(buffer));
    } else {
        return "Weather{" DEFAULT_ZIP;
    }
}


int main(int argc, char *argv[])
{
    printf("Starting weather client on port "PORT"\n");
    char *filter = make_filter(argc, argv);
    printf("Applying filter for zipcode starting with: '%s'\n", filter + 8);
    void *context = zmq_ctx_new();
    void *subscriber = zmq_socket(context, ZMQ_SUB);
    int rc = zmq_connect(subscriber, "tcp://localhost:"PORT);
    assert(rc == 0);

    // subscribe to weather updates from zipcode, default is NYC, 10001
    rc = zmq_setsockopt(subscriber, ZMQ_SUBSCRIBE, filter, strlen(filter));
    assert(rc == 0);

    int i;
    long total_temp = 0;
    for (i = 0; i < 100; i++) {
        char *weather_str = s_recv(subscriber);
        printf("received: %s\n", weather_str);
        int zip, temp, humidity;
        sscanf(weather_str, "Weather{%d, %d, %d}", &zip, &temp, &humidity);
        total_temp += temp;

        free(weather_str);
    }
    printf("Average temperature for zipcode starting with '%s' was %dF\n", filter + 8,
           (int) (total_temp / i));

    zmq_close(subscriber);
    zmq_ctx_destroy(context);
    return 0;
}
