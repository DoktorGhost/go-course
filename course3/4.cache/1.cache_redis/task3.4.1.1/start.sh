#!/bin/bash

# Проверяем, запущен ли контейнер с именем my-redis
if [ "$(docker ps -q -f name=my-redis)" ]; then
    echo "Контейнер my-redis уже запущен."
else
    if [ "$(docker ps -aq -f name=my-redis)" ]; then
        # Контейнер my-redis существует, но остановлен, запускаем его
        echo "Контейнер my-redis существует. Запускаем его..."
        docker start my-redis
    else
        # Если контейнер не существует, создаем и запускаем его
        echo "Контейнер my-redis не существует. Создаем и запускаем..."
        docker run --name my-redis -p 6379:6379 -d redis
    fi
fi
