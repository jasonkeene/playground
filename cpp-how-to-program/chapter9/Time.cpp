#include <iostream>
#include <iomanip>
#include <stdexcept>

#include "Time.h"

using namespace std;

Time::Time(int hour, int minute, int second)
{
    setTime(hour, minute, second);
}

Time::~Time()
{
    std::cout << "destructor called" << std::endl;
}

void Time::setTime(int h, int m, int s)
{
    setHour(h);
    setMinute(m);
    setSecond(s);
}

void Time::setHour(int h)
{
    if (h >= 0 && h < 24) {
        hour = h;
    } else {
        throw invalid_argument("hour was out of range");
    }
}

unsigned int Time::getHour() const
{
    return hour;
}

void Time::setMinute(int m)
{
    if (m >= 0 && m < 60) {
        minute = m;
    } else {
        throw invalid_argument("minute was out of range");
    }
}

unsigned int Time::getMinute() const
{
    return minute;
}

void Time::setSecond(int s)
{
    if (s >= 0 && s < 60) {
        second = s;
    } else {
        throw invalid_argument("second was out of range");
    }
}

unsigned int Time::getSecond() const
{
    return second;
}

void Time::printUniversal() const
{
    cout << setfill('0') << setw(2) << hour << ":"
        << setw(2) << minute << ":"
        << setw(2) << second;
}

void Time::printStandard() const
{
    int std_hour = hour % 12;
    std_hour = std_hour == 0 ? 12 : std_hour;

    cout << setfill('0') << setw(2) << std_hour << ":"
        << setw(2) << minute << ":"
        << setw(2) << second << " "
        << (hour < 12 ? "AM" : "PM");
}
