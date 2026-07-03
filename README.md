```markdown
# EduSphere — Управление образовательным пространством

Веб-приложение для управления базой знаний, службами колледжа и обращениями студентов. Состоит из backend на Go (Chi) и одностраничного интерфейса (HTML/CSS/JS), который общается с API.

## Технологии

- **Backend**: Go 1.25+, библиотеки:
  - [chi](https://github.com/go-chi/chi) – маршрутизация
  - [go-chi/cors](https://github.com/go-chi/cors) – CORS
  - [validator/v10](https://github.com/go-playground/validator) – валидация
  - [lib/pq](https://github.com/lib/pq) – драйвер PostgreSQL
  - [joho/godotenv](https://github.com/joho/godotenv) – загрузка .env
- **База данных**: PostgreSQL 14+
- **Frontend**: чистый HTML + CSS + JavaScript (без фреймворков)

## Установка и запуск

### 1. Установите PostgreSQL и Go

- [Скачать PostgreSQL](https://www.postgresql.org/download/)
- [Скачать Go](https://golang.org/dl/) (версия 1.23 или выше)

### 2. Клонируйте репозиторий

```bash
git clone https://github.com/your-repo/college_faq.git
cd college_faq