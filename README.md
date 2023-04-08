# curso-go
Fullcycle Go Intensivo - Fundação

Intalador do GoLang.:
https://go.dev/learn/

Documentação para Template.:
https://pkg.go.dev/html/template

Conversor de JSON para Go struct.:
https://mholt.github.io/json-to-go/

Compilador online do GoLang.:
https://go.dev/play/


-> Aula - Preparando base de código:

Comandos dockers utilizados.:
docker-compose up -d
docker-compose exec mysql bash

Script de criação de tabela.:
create table products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));

Após a criação do go.mod executar.:
go mod tidy

-> Consulta após inserção de novo produto no banco.:
select * from products