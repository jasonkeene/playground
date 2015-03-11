
#ifndef TIME_H
#define TIME_H

class Time
{
public:
    explicit Time(int=0, int=0, int=0);
    ~Time();
    void setTime(int, int, int);
    void setHour(int);
    unsigned int getHour() const;
    void setMinute(int);
    unsigned int getMinute() const;
    void setSecond(int);
    unsigned int getSecond() const;
    void printUniversal() const;
    void printStandard() const;
private:
    unsigned int hour;
    unsigned int minute;
    unsigned int second;
};

#endif
