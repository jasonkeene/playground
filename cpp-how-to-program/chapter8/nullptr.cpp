#include <iostream>

int main()
{
    int *ptr1 = {};
    int *ptr2 = 0;
    int *ptr3 = nullptr;

    std::cout << "ptr1: " << ptr1 << std::endl;
    std::cout << "ptr2: " << ptr2 << std::endl;
    std::cout << "ptr3: " << ptr3 << std::endl;

    return 0;
}
