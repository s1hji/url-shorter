## Быстрый запуск через Docker(используется postgres хранилище) :

docker compose up --build 

## Запуск через терминал : 
-
go run main.go -storage memory - на локальном хранилище машины
---
go run main.go -storage postgres -connect="postgres://user:pass@localhost:5432/urlshorter?sslmode=disable" - создается база данных, если не создана и вся дальнейшая работа будет происходить в ней

## api:
-
curl -X POST http://localhost:8080/shorter -d '{"url":"telegram.org"}' (пример post запроса - должен выдать json типа {"short":"abc123"})
---
curl -v http://localhost:8080/(короткая ссылка) (вернет 302 Found и перебросит на ориг ссылку)
