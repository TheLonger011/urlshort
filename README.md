# URL Shortener Ru

Сервис для сокращения ссылок на Go с веб-интерфейсом.

## Возможности

- Сокращение длинных URL
- Кастомные алиасы для ссылок
- Автоматическая генерация коротких ссылок
- Веб-интерфейс для удобного использования
- REST API для интеграции

## Технологии

- **Go* 1.25.0
- **PostgreSQL** 18
- **chi/router** - маршрутизация
- **HTML/CSS/JS** - frontend

## Установка и запуск

### 1. Клонирование репозитория

git clone https://github.com/TheLonger011/urlshort.git
cd urlshortener

### 2. Настройка

отредактируйте файл local.yaml

storage_path: "postgres://postgres:pass@localhost:5432/name?sslmode=disable"

### 3. запуск 

make run 
или
go run cmd/url-shortener/main.go



# URL Shortener EN

URL shortening service written in Go with a web interface.

## Features

- Shorten long URLs
- Custom aliases for links
- Automatic short link generation
- Web interface for easy use
- REST API for integration

## Technologies

- **Go** 1.25.0
- **PostgreSQL** 18
- **chi/router** - routing
- **HTML/CSS/JS** - frontend

## Installation and Running

### 1. Clone the repository

git clone https://github.com/TheLonger011/urlshort.git
cd urlshortener

### 2. Configuration
Edit the config/local.yaml file:

storage_path: "postgres://postgres:pass@localhost:5432/name?sslmode=disable"

### 3. Run

make run
or
go run cmd/url-shortener/main.go
