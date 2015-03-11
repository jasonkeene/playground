#include <iostream>
#include <stdexcept>

#include "Time.h"

int main(int argc, char *argv[])
{
    Time &t = *(new Time());

    std::cout << "initial universal time: ";
    t.printUniversal();
    std::cout << std::endl;

    std::cout << "initial standard time: ";
    t.printStandard();
    std::cout << std::endl;

    std::cout << "trying to set to 13, 40, 30" << std::endl;
    t.setTime(13, 40, 30);

    std::cout << "new universal time: ";
    t.printUniversal();
    std::cout << std::endl;

    std::cout << "new standard time: ";
    t.printStandard();
    std::cout << std::endl;

    std::cout << "trying to set to 55, 40, 30" << std::endl;
    try {
        t.setTime(55, 40, 30);
    } catch (std::invalid_argument &e) {
        std::cout << "caught exception: " << e.what() << std::endl;
    }

    std::cout << "after exception universal time: ";
    t.printUniversal();
    std::cout << std::endl;

    std::cout << "after exception standard time: ";
    t.printStandard();
    std::cout << std::endl;

    return 0;
}
