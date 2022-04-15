# Стандартный веб-сервер

## Config
В Go принято, что конфиги хранятся в файлах:
* env
* toml
Всегда есть дефолтные значения(кроме каких-то секретных данных, например: лог пасс)
Можно передавать через флаг при запуске приложения 
``` api.exe -path configs/api.toml```

## Database
Шаги реализации взаимодействия с БД:
* Определить модель данных
* Обработчик модели
* Выделение публичных обработчиков

### Библиотеки для работы с бд
* database/sql
* sqlx
* gosql

### Инициализация модели хранилища
Содержит инстанс конфига и конструктор. Атрибутом конфига является лишь строка подключения

### Мигация
Установим либу
```
brew install golang-migrate
 ```

Создадим миграционную папку
```
migrate create -ext sql -dir migrations UsersCreationMigration
```

Применение миграции
```
migrate -path migrations -database "postgres://localhost:5432/restapi?sslmode=disable&user=postgres&password=postgres" up
```