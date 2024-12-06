#!/bin/bash

# Подключаемся к Redis и выполняем команды
echo "Запуск миграции Redis..."

# Пример команд для добавления данных в Redis
redis-cli SET my_key "Hello, Redis"
redis-cli LPUSH my_list "Item 1" "Item 2" "Item 3"
redis-cli HSET my_hash field1 "value1" field2 "value2"

echo "Миграция Redis завершена!"