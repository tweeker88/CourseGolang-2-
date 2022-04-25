# CRUD сервис и БД

Выполним команду ```migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" down```

## Модели

Это "отображение" таблиц бд на нашем проекте

* ```./internal/app/models/user.go```
* ```./internal/app/models/article.go```

## Репозитории

Работать с моделями будем через репозитории. Для этого инициализируем 2 файла:
* ```storage/userrepository.go```
* ```storage/articlerepository.go```

Определяем публичные методы репозиториев для манипуляции с данными

