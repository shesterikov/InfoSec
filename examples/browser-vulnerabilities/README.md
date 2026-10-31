# OWASP Browser Security Lab

Локальный учебный стенд для демонстрации XSS, CSRF, clickjacking, CORS и защитных HTTP-заголовков. Приложение намеренно содержит уязвимости и использует тестовые данные.

## Запуск

Нужны Docker и Docker Compose.

```bash
docker compose up --build -d
```

- http://localhost:3000 — уязвимое приложение
- http://localhost:3002 — приложение с защитными механизмами
- http://localhost:3001 — атакующий сайт

Используйте `localhost` для всех адресов и отдельный профиль браузера: cookie не разделяются по портам.

## Структура кода

В `app/` находится целевое приложение, в `attacker/` — атакующий сайт. В каждом `server.js` запускает сервер, `src/app.js` подключает маршруты, `src/routes/` содержит сценарии, а `src/views/` — HTML-шаблон страницы. В целевом приложении `src/middleware/` отвечает за сессии и защитные заголовки, `src/config.js` — за режим и порт, `src/utils/` — за экранирование HTML.

## Сценарии

Проходите каждый сценарий сначала на порту `3000`, затем на `3002`, чтобы сравнить результат.

### 1. Reflected XSS

Откройте http://localhost:3000/search?q=hello, вставьте в поиск и нажмите «Искать»:

```html
<img src=x onerror=alert('reflected-xss')>
```

На порту `3000` появится alert. Повторите на http://localhost:3002/search?q=hello — ввод отобразится как текст.

### 2. Stored XSS

Откройте http://localhost:3000/comments и добавьте комментарий:

```html
<img src=x onerror=alert('stored-xss')>
```

Обновите страницу: alert появится снова. Повторите на http://localhost:3002/comments — комментарий отобразится как текст.

### 3. DOM-based XSS

Откройте адрес:

```text
http://localhost:3000/dom?name=%3Cimg%20src%3Dx%20onerror%3Dalert('dom-xss')%3E
```

Появится alert. Замените порт на `3002`: значение параметра отобразится как текст.

### 4. CSRF

1. Откройте http://localhost:3000/account и запомните email.
2. В том же браузере откройте http://localhost:3001/csrf?target=3000 — форма отправится автоматически.
3. Откройте http://localhost:3000/account снова: email изменится на `changed-by-csrf@example.test`.
4. Откройте http://localhost:3002/account, затем http://localhost:3001/csrf?target=3002 — сервер вернёт `403`.
5. Вернитесь на http://localhost:3002/account: email останется прежним.

Порты localhost имеют разные origin, но один site, поэтому cookie с `SameSite=Lax` здесь не предотвращает CSRF.

### 5. Clickjacking

1. Откройте http://localhost:3001/clickjack?target=3000 и попробуйте нажать видимую область кнопки. Целевая страница загружена поверх неё в полупрозрачном iframe; его можно изучить в DevTools.
2. Откройте http://localhost:3001/clickjack?target=3002 — браузер заблокирует iframe. Причину можно увидеть в консоли DevTools.

### 6. CORS

1. Откройте http://localhost:3000/account, чтобы создать сессию.
2. Откройте http://localhost:3001/cors?target=3000 и нажмите «Выполнить fetch»: появится JSON с тестовым email.
3. Откройте http://localhost:3002/account, затем http://localhost:3001/cors?target=3002.
4. Нажмите «Выполнить fetch»: браузер заблокирует чтение ответа. Посмотрите ошибку в консоли DevTools.

### 7. CSP и HTTP-заголовки

Сравните заголовки приложений:

```bash
curl -I http://localhost:3000
curl -I http://localhost:3002
```

На порту `3002` проверьте `Content-Security-Policy` с nonce и `frame-ancestors 'none'`, `X-Frame-Options: DENY`, `X-Content-Type-Options: nosniff` и `Referrer-Policy`. На порту `3000` CSP разрешает `unsafe-inline` и `unsafe-eval`.

В браузере эти заголовки доступны в DevTools → Network → запрос документа → Response Headers.

## Управление стендом

Статус контейнеров:

```bash
docker compose ps
```

Логи:

```bash
docker compose logs -f
```

Остановка и сброс тестовых данных, которые хранятся в памяти:

```bash
docker compose down
```
