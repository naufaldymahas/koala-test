# REST-Test

REST API using golang, gorm, postgresql

## Minimal Requirement
- [Golang](https://golang.org/dl/) version go1.15.2
- [PostgreSQL](https://www.postgresql.org/download/) version 12.2

## Usage
Restore dump to your PostgreSQL
```
dump-koala-202012310211.sql
```

Setting .env before start the program, .env example:
```
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_NAME=koala
DB_USER=dev
DB_PASSWORD=password

SALT_JWT=secret
EXP_NORMAL_TOKEN_JWT=10 #minute
EXP_REFRESH_TOKEN_JWT=2 #hour
```
Run this command on your terminal
```bash
go run main.go
```

## Author
Naufaldy Mahas (naufaldymahas@gmail.com)