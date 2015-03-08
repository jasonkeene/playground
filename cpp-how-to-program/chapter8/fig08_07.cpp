#include <iostream>

using namespace std;

void cubeByReference(int *nPtr)
{
    *nPtr = *nPtr * *nPtr * *nPtr;
}

int main()
{
    int number = 5;

    cout << "The original value of number is " << number << endl;
    cubeByReference(&number);
    cout << "The new value of number is " << number << endl;

    return 0;
}
