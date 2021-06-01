# Unit Testing using Go-chi and SQLMock

## Overview

Some ProTip about Golang unit testing when the app using [go-chi](https://github.com/go-chi/chi) and [sqlmock](https://github.com/DATA-DOG/go-sqlmock).

## How to Run

### Docker

- Run your docker
- Run `docker-compose up`

### Non Docker

- Copy paste `.env.sample` to `.env` and set your database credential
- Run `go run app/main.go`

## Test

- Run `go test -v ./...`

## Routes

| Endpoint | Method | Description
| --- | --- | --- |
| `/v1/users` | `GET` | getting list of users |
| `v1/users/{id}` | `GET` | getting user by user id |