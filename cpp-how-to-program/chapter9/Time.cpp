#include <iostream>
#include <iomanip>
#include <stdexcept>

#include "Time.h"

using namespace std;

Time::Time() : hour(0), minute(0), second(0) {}

void Time::setTime(int h, int m, int s)
{
    if (
        h >= 0 &&
        h < 24 &&
        m >= 0 &&
        m < 60 &&
        s >= 0 &&
        s < 60
    ) {
        hour = h;
        minute = m;
        second = s;
    } else {
        throw invalid_argument(
            "hour, minute, and/or second was out of range");
    }
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
