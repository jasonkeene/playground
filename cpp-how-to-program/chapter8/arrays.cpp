#include <iostream>

void double_it(int * const arr, const size_t len)
{
    for (size_t i = 0; i < len; i++) {
        arr[i] += arr[i];
    }
}

void print_it(const int * const arr, const size_t len)
{

    for (size_t i = 0; i < len; i++) {
        std::cout << arr[i] <<  (i < len-1 ? " " : "");
    }
    std::cout << std::endl;
}

int main()
{
    const size_t my_len = 6;
    int my_arr[my_len] = {4, 1, 16, 0, 8, 2};
    double_it(my_arr, my_len);
    print_it(my_arr, my_len);
    std::sort(std::begin(my_arr), std::end(my_arr));
    print_it(my_arr, my_len);
    return 0;
}
