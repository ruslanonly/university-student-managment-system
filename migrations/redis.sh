#!/bin/bash
set -e  # Exit on error

redis-server --logfile /var/log/redis/redis.log &

sleep 5

echo "Starting Redis migration script..."

redis-cli HSET student:1 full_name "Петров Алексей Сергеевич" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:2 full_name "Смирнов Дмитрий Валерьевич" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:3 full_name "Кузнецова Мария Александровна" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:4 full_name "Васильев Николай Игоревич" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:5 full_name "Соколова Ольга Вячеславовна" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:6 full_name "Григорьев Андрей Викторович" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:7 full_name "Зайцева Екатерина Анатольевна" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:8 full_name "Михайлов Артем Сергеевич" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:9 full_name "Попова Валерия Евгеньевна" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:10 full_name "Кравцов Кирилл Александрович" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:11 full_name "Иванов Иван Иванович" group_id 1 group_name "БСБО-01-21"
redis-cli HSET student:12 full_name "Фадеев Всеволод Вадимович" group_id 2 group_name "БСБО-02-21"
redis-cli HSET student:13 full_name "Усанкин Александр Александрович" group_id 2 group_name "БСБО-02-21"
redis-cli HSET student:14 full_name "Рзаев Руслан Халидович" group_id 2 group_name "БСБО-02-21"
redis-cli HSET student:15 full_name "Тимофеев Максим Павлович" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:16 full_name "Лебедев Артем Юрьевич" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:17 full_name "Фролова Дарина Васильевна" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:18 full_name "Ильин Иван Игоревич" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:19 full_name "Чернова Светлана Павловна" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:20 full_name "Андреев Павел Владиславович" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:21 full_name "Егорова Виктория Николаевна" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:22 full_name "Морозов Михаил Евгеньевич" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:23 full_name "Сидорова Ирина Борисовна" group_id 3 group_name "БСБО-04-21"
redis-cli HSET student:24 full_name "Борисов Сергей Юрьевич" group_id 3 group_name "БСБО-04-21"

echo "Redis migration completed."

tail -f /var/log/redis/redis.log