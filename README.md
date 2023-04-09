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

Criar arquivo go.mod
go mod init github.com/luizffdemoraes/curso-go/database

Após a criação do go.mod executar.:
go mod tidy

-> Consulta após inserção de novo produto no banco.:
select * from products

GORM - Ferramenta para Mapeamento de Objeto Relacional
https://gorm.io/

Em caso de falha ao tentar realizar o drop table.:
SET FOREIGN_KEY_CHECKS=0; -- to disable them
SET FOREIGN_KEY_CHECKS=1; -- to re-enable them

One to One (um para um) 
-> Belongs to 
-> Has One

One to many (um para muitos)
-> Has Many

Many to Many (Muitos para muitos)

Lock pessimista
-> Lock a tabela, a linha fica numa linha de espera para poder realizar alteração após comitar

Lock otimista
-> Ele versiona ao realizar alterações