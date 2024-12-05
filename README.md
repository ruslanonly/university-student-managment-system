Инструкция по миграциям (а то neo4j будет офигевать):
 - Запустить `docker-compose up` впервые
 - Подождать пока все проинитится
 - Выключить контейнеры `docker-compose down`
 - Запустить `docker-compose -f .\docker-compose-run-migrations.yml up`
 - Подождать пока neo4j и elastic-init выйдут
 - Выключить контейнеры `docker-compose down`
 - В дальнейшем запускать `docker-compose up`