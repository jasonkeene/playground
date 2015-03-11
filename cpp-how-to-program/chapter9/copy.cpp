#include <iostream>

class Foo {
public:
    Foo(int num=0) : number(num)
    {
        std::cout << "Foo(" << number << ")" << std::endl;
    }
    Foo(const Foo &other) : number(other.number)
    {
        std::cout << "copy Foo(" << number << ")" << std::endl;
    }
    ~Foo()
    {
        std::cout << "~Foo(" << number << ")" << std::endl;
    }
    int getNumber()
    {
        return number;
    }
private:
    int number;
};

void display(Foo f)
{
    std::cout << "display: " << f.getNumber() << std::endl;
}

int main(int argc, char *argv[])
{
    Foo f1(1);
    Foo f2(2);
    display(f1);
    Foo f3 = f2;
    return 0;
}
