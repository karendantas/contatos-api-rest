# API REST com GoLang e Gin

Api simples para CRUD de contatos

## Tecnologias

- Go
- PostgreSQL
- Docker

## Iniciando um projeto Go

```bash
    go mod init go-api
    go get github.com/gin-gonic/gin
    cd cmd
    go run main.go
```

## Como rodar localmente

```bash
    docker compose up -d go_db
    docker container ls

```

# SQL

<pre lang="markdown"> 
```sql create table if not exists contact ( contact_id serial primary key, contact_name varchar(50) not null, email varchar(100) unique not null ); -- insert into contact (contact_name, email) values ('Karen', 'karen@gmail.com'); -- insert into contact (contact_name, email) values ('Mirla', 'mirla@gmail.com'); select * from contact; 
``` 
</pre>
