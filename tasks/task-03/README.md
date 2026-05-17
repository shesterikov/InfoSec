# Практическая работа 3

## Хеширование паролей

### Цель работы

Реализовать безопасное хранение паролей с использованием хеша и соли

### Технические требования:

- Наличие интернет-соединения
- Наличие [Docker](https://docs.docker.com/desktop/) и [Docker Compose](https://docs.docker.com/compose/install/)
- Наличие [cURL](https://curl.se/download.html) / [Postman](https://www.postman.com/downloads/) / [Insomnia](https://insomnia.rest/download)

### Ход работы:

1. Запустите приложение при помощи команды
   `docker compose up -d --build`

2. Убедитесь, что приложение запущено корректно при помощи следующего cURL запроса

```cURL
curl --location 'http://localhost:3000/health'
```

В случае успешного запуска при выполнении cURL запроса отображается следующий текст
![Рисунок 3.1 - Главный экран Postman с успешным ответом от API](/assets/task-03/1.png)

3. Выполните настройку подключения и проверку существования таблицы users. Для этого воспользуйтесь информацией, размещенной в разделе “Установка подключения к СУБД PostgreSQL посредством PGAdmin”

4. Выполните следующий cURL запрос для создания (регистрации) нового пользователя

```cURL
curl --location 'http://localhost:3000/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test@example.com",
    "password": "super_secret_password"
}'
```

![Рисунок 3.2 - cURL запрос для создания нового пользователя](/assets/task-03/2.png)

5. Убедитесь в корректности сохраняемых в БД данных при помощи SQL запроса.

```SQL
SELECT * FROM public.users;
```

![Рисунок 3.3 - SQL запрос для отображения содержимого таблицы users после вставки данных](/assets/task-03/3.png)
Обратите внимание, что пароль представлен в открытом виде. Соль при этом имеет `null` значение.

6. Модифицируйте исходный код файла `server.js` для реализации безопасного хранения пользовательского пароля с использованием хеширования и соли.

### Документация:

[Sequelize](https://sequelize.org/docs/v6/core-concepts/model-basics/) для расширения модели БД

[Node:Crypto](https://nodejs.org/api/crypto.html) для работы с криптографическим API Node

### Контрольные вопросы:

1. Что такое хеш функция?
2. Что такое соль?
3. Что такое коллизии?
