#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "weather.h"


#define WEATHER_FMT "Weather{%05d, %.1f, %.2f}"


Weather *Weather_create(int zip, float temp, float humidity)
{
    Weather *weather = malloc(sizeof(Weather));
    weather->zip = zip;
    weather->temp = temp;
    weather->humidity = humidity;
    weather->str = NULL;
    Weather_str(weather);
    return weather;
}


Weather *Weather_create_fake()
{
    return Weather_create(arc4random() % 99999,
                          (int)(arc4random() % 215) - 80,
                          (int)(arc4random() % 50) + 10);
}


void Weather_destroy(Weather *weather)
{
    if (weather->str != NULL) free(weather->str);
    free(weather);
}


void Weather_print(Weather *weather)
{
    printf(WEATHER_FMT"\n", weather->zip, weather->temp,
           weather->humidity);
}


void Weather_str(Weather *weather)
{
    if (weather->str != NULL) free(weather->str);
    char buffer[256];
    snprintf(buffer, sizeof(buffer), WEATHER_FMT, weather->zip, weather->temp,
             weather->humidity);
    weather->str = strndup(buffer, sizeof(buffer));
}

