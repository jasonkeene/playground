#!/bin/bash
container=$(docker run -Pd elasticsearch)
echo $container
docker port $container | head -n 1 | awk '{print $3}'
