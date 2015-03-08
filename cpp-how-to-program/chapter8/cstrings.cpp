#include <iostream>

int main()
{
    char str1[] = "test1";
    char str2[] = "test2";
    const char *str3 = "test3";
    char str4[] = "test4";

    std::cout << "str1: " << (long)str1 << std::endl;
    std::cout << "str2: " << (long)str2 << std::endl;
    std::cout << "str3: " << (long)str3 << std::endl;
    std::cout << "str4: " << (long)str4 << std::endl;

    return 0;
}
