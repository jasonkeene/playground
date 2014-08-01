

typedef struct IntNode
{
    int data;
    struct IntNode *next;
} IntNode;


typedef struct IntList
{
    IntNode *head;
} IntList;


IntList *IntList_create();
void IntList_push(IntList *int_list, int data);
/* int IntList_pop(IntList *int_list); */
/* int IntList_lpop(IntList *int_list); */
void IntList_print(IntList *int_list);
void IntList_destroy(IntList *int_list);


IntList *get_ports(int argc, char *argv[]);

char *connection_str(int port);
