
typedef struct Weather {
    int zip;
    float temp;
    float humidity;
    char *str;
} Weather;


Weather *Weather_create(int zip, float temp, float humidity);

Weather *Weather_create_fake();

void Weather_destroy(Weather *weather);

void Weather_print(Weather *weather);

void Weather_str(Weather *weather);
