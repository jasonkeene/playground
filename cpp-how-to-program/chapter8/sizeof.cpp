#include <iostream>

int main()
{
    int x = 12;
    int my_arr[] = {1, 2, 3, 4, 5, 6};

    std::cout << "sizeof x: " << sizeof x << std::endl;
    std::cout << "sizeof my_arr: " << sizeof my_arr << std::endl;
    std::cout << "sizeof my_arr[0]: " << sizeof my_arr[0] << std::endl;
    std::cout << "sizeof my_arr/sizeof my_arr[0]: " << sizeof my_arr/sizeof my_arr[0] << std::endl;

    char c;
    short s;
    int i;
    long l;
    long long ll;
    float f;
    double d;
    long double ld;
    int array[20];
    int *ptr = array;

    std::cout << std::endl;
    std::cout << "sizeof c: " << sizeof c << std::endl;
    std::cout << "sizeof (char): " << sizeof (char) << std::endl;
    std::cout << "sizeof s: " << sizeof s << std::endl;
    std::cout << "sizeof (short): " << sizeof (short) << std::endl;
    std::cout << "sizeof i: " << sizeof i << std::endl;
    std::cout << "sizeof (int): " << sizeof (int) << std::endl;
    std::cout << "sizeof l: " << sizeof l << std::endl;
    std::cout << "sizeof (long): " << sizeof (long) << std::endl;
    std::cout << "sizeof ll: " << sizeof ll << std::endl;
    std::cout << "sizeof (long long): " << sizeof (long long) << std::endl;
    std::cout << "sizeof f: " << sizeof f << std::endl;
    std::cout << "sizeof (float): " << sizeof (float) << std::endl;
    std::cout << "sizeof d: " << sizeof d << std::endl;
    std::cout << "sizeof (double): " << sizeof (double) << std::endl;
    std::cout << "sizeof ld: " << sizeof ld << std::endl;
    std::cout << "sizeof (long double): " << sizeof (long double) << std::endl;
    std::cout << "sizeof array: " << sizeof array << std::endl;
    std::cout << "sizeof ptr: " << sizeof ptr << std::endl;

    return 0;
}
