# Ozon Test Task - Link Shortener
"Задача:

Реализовать сервис, предоставляющий API по созданию сокращённых ссылок.

Ссылка должна быть:
* Уникальной; на один оригинальный URL должна ссылаться только одна сокращенная ссылка;
*  Длиной 10 символов;
*  Из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание).

Сервис должен быть написан на Go и принимать следующие запросы по http:
1. Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый.
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный.
   Условие со звёздочкой(будет большим плюсом):
   Сделать работу сервиса через GRPC, то есть составить proto и реализовать сервис с двумя соответствующими эндпойнтами

Решение должно соответствовать условиям:
*  Сервис распространён в виде Docker-образа;
*  В качестве хранилища ожидаем in-memory решение и PostgreSQL. Какое хранилище использовать, указывается параметром при запуске сервиса;
*  Реализованный функционал покрыт тестами.
"

## Installation & Run
```bash
# Download this project
git clone https://github.com/Snegniy/ozon-testtask.git
```


```bash
# Build and Run
cd github.com/Snegniy/testTaskResponseApi/cmd
go build -o testapp
./testapp

```

## Command
```bash
# Запустить приложение в контейнере с локальным хранилищем
make local
```

```bash
# Запустить приложение в контейнере с PostgreSQL хранилищем
make postgres
```

```bash
# Запустить тесты
make test
```

```bash
# Запустить локально
make run
```

## Structure
```
├── api
│   ├── proto
│   │   ├── shortlinks.proto // gRPC proto файл
├── cmd
│   ├── main.go          // запуск приложения
├── internal
│   ├── config
│   │   ├── config.go   // инициализация конфигурации приложения 
│   ├── model
│   │   ├── model.go // модель данных
│   ├── repository
│   │   ├── memdb
│   │   │  ├── repository.go // локальное хранилище
│   │   ├── postgre
│   │   │  ├── repository.go // postgreSQL хранилище
│   ├── service
│   │   ├── service.go // бизнес-логика
│   ├── transport
│   │   ├── grpc
│   │   │  ├── grpc.go // gRPC обработчики
│   │   │  ├── grpcserver.go // gRPC сервер
│   │   ├── rest
│   │   │  ├── rest.go // HTTP обработчики
│   │   │  ├── response.go // отправка ответа в формате JSON
├── migrations
│   ├── init.sql        // начальные настройки БД
├── pkg
│   ├── api
│   │   ├── shortlinks.pb.go // protoc-gen-go-grpc
│   │   ├── shortlinks_grpc.pb.go // protoc-gen-go-grpc
│   ├── graceful
│   │   ├── server.go  // запуск graceful HTTP сервера
│   ├── logger
│   │   ├── logger.go // инициализация логгера
│   ├── postgres
│   │   ├── postgres.go // инициализация PostgreSQL хранилища
│   ├── urlgenerator
│   │   ├── urlgenerator.go // генератор короткой ссылки
├── .env  // конфигурационные установки по умолчанию
├── gocker-compose.yml
├── Dockerfile
├── go.mod
├── Makefile
```

## HTTP API

#### /
* `GET` : {"url":"site:} - Запрос на получение короткой ссылки

#### /
* `POST` : {"url":"short_url:} - Запрос на получение полной ссылки
