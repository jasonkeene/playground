

typedef struct PortNode
{
    int port;
    struct PortNode *next;
} PortNode;


typedef struct PortList
{
    PortNode *head;
} PortList;


PortList *PortList_create();
void PortList_push(PortList *port_list, int port);
void PortList_destroy(PortList *port_list);
PortList *get_port_list(int argc, char *argv[]);


char *pusher_str(int port);
