# Лекция 8. OWASP и безопасность веб-приложений

[← Программа курса](../../README.md#программа-курса)

## Основные вопросы
- OWASP и его роль в безопасности программного обеспечения.
- Проекты OWASP: Top 10, ASVS, WSTG, Cheat Sheet Series, ZAP.
- OWASP Top 10:2025.
- Клиентские и серверные уязвимости веб-приложений.
- Cross-Site Scripting (XSS).
- Reflected, Stored и DOM-based XSS.
- Sources и Sinks в клиентском JavaScript.
- Защита от XSS: Output Encoding, безопасные DOM API, HTML Sanitization.
- Content Security Policy (CSP).
- CSP Report-Only и основные директивы CSP.
- Subresource Integrity (SRI).
- Cross-Site Request Forgery (CSRF) и механизмы защиты.
- Clickjacking и защита от встраивания страниц.
- Same-Origin Policy (SOP).
- Cross-Origin Resource Sharing (CORS).
- CORS Preflight и типичные ошибки конфигурации.
- Изоляция iframe и взаимодействие через postMessage.
- HTTP Security Headers.
- Многоуровневая защита веб-приложений (Defense in Depth).
- Основные принципы безопасной веб-разработки.

## Примеры

- [Уязвимости браузерных приложений и способы защиты](../../examples/browser-vulnerabilities) — учебный стенд для демонстрации XSS, CSRF, clickjacking, CORS и защитных HTTP-заголовков.
