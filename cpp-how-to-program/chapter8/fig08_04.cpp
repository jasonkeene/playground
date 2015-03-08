#include <iostream>

using namespace std;

int main()
{
    int a = 7;
    int *aPtr = &a;

    cout << "The address of a is " << &a << endl
         << "The value of aPtr is " << aPtr << endl
         << endl
         << "The value of a is " << a << endl
         << "The value of *aPtr is " << *aPtr << endl;

    return 0;
}
