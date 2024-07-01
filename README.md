# ZeroAgency

Этот проект представляет собой JSON REST API сервер, написанный на языке Go с использованием фреймворка Fiber. Сервер предоставляет две основные ручки:

### *Возможные функции:*
- *POST /edit/*
изменение новости по заданному идентификатору (Id). Для изменения доступны поля Title, Content и категории новости.
- *GET /list* - получение списка всех новостей в системе. Каждая новость представлена полями Id, Title, Content и списком связанных категорий.

### *База данных:*

*Для хранения данных используется СУБД PostgreSQL. В проекте определены две таблицы:*
- News - содержит информацию о новостях, включая идентификатор, заголовок и контент.
- NewsCategories - связующая таблица между новостями и категориями, хранит связи между NewsId и CategoryId.

## Авторизация

Для авторизации в системе используется JWT (JSON Web Token). Для получения JWT токена необходимо отправить запрос методом `GET` /login

## Документация Swagger

Проект включает документацию Swagger для облегчения работы с API. После запуска сервера вы можете получить доступ к документации: /swagger/index.html
### Пример использования Swagger с авторизацией

1. Перейдите по URL документации Swagger.
2. Нажмите на кнопку `Authorize`.
3. Вставьте ваш JWT токен в поле `Value` в формате `Bearer <YOUR_TOKEN>`.
4. Нажмите на кнопку `Authorize`.

Теперь вы можете использовать все защищенные эндпоинты API, вставив JWT токен для авторизации.

### *Формат входных данных:*

- Для ручки POST /edit/
  данные принимаются в формате JSON и включают поля:
    ```json
    {
    "Id": 64,
    "Title": "Lorem ipsum",
    "Content": "Dolor sit amet <b>foo</b>",
    "Categories": [1, 2, 3]
    }
    ```
  Поля Title, Content и Categories являются опциональными. Если они не указаны, соответствующие данные не изменяются.

### *Формат данных на выходе:*
- Ручка GET /list возвращает список новостей в формате JSON:
  ```json
    {
      "Success": true,
      "News": [
        {
          "Id": 64,
          "Title": "Lorem ipsum",
          "Content": "Dolor sit amet <b>foo</b>",
          "Categories": [1, 2, 3]
        },
        {
          "Id": 1,
          "Title": "first",
          "Content": "tratata",
          "Categories": [1]
        }
      ]
    }
    ```
  
## Окружение

### Требования
- Docker 20.x.x
- Docker Compose 1.29.x
- Go 1.16+

### Переменные окружения

| Параметр                 | Описание                                         | Default                 |
|--------------------------|--------------------------------------------------|-------------------------|
| `MODE`                   | Режим окружения                                  | `development`           |
| `PREFORK`                | Флаг предзагрузки процессов                      | `true`                  |
| `TOKEN`                  | Секретный ключ для подписи JWT токенов           | `secret`                |
| `HTTP_SERVER_PORT `      | Порт HTTP сервера                                | `8080`                  |
| `DB_HOST`                | Хост базы данных                                 | `localhost`             |
| `DB_PORT`                | Порт базы данных                                 | `5432`                  |
| `DB_NAME`                | Имя базы данных                                  | `pkk_db`                |
| `DB_USER_NAME`           | Имя пользователя базы данных                     | `user`                  |
| `DB_PASSWORD`            | Пароль пользователя базы данных                  | `secret`                |
| `DB_SSL_MODE`            | Режим SSL для базы данных                        | `disable`               |
| `DB_DRIVER_NAME`         | Имя драйвера базы данных                         | `postgres`              |
| `DB_MAX_CONNS`           | Максимальное количество соединений БД            | `20`                    |
| `DB_MAX_IDEL_CONNS`      | Максимальное количество неактивных соединений БД | `10`                    |
| `DB_CONNS_MAX_LIFE_TIME` | Максимальное время жизни соединений БД           | `5s`                    |
| `MIGRATION_URL`          | URL для миграций                                 | `file://migrations`     |

## Установка

**PROJECT**

- Создать новую директорию для проекта. В консоли перейти в созданную директорию и написать: git clone https://github.com/ShevelevEvgeniy/ZeroAgencyTest.git

**DOCKER**

*Сборка:*

Скопировать файл .env.dist и переименовать в .env, настроить параметры окружения cp .env.dist .env
Для развертывания, запустите установку, выполнив команду ниже: make install

*Служебное:*
- make migrate-up - Запуск миграций
- make migrate-down - Откат миграций
- make migrate-create name="$" - Создание новой миграции

Если у вас возникли вопросы или проблемы, вы можете связаться со мной по адресу Z_shevelev@mail.ru