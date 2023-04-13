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

Comandos dockers utilizados para criar o banco de dados em container e acessa-lo.:
-> docker-compose up -d

-> docker-compose exec mysql bash

-> mysql -u root -p goexpert
-> senha

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

1- Packaging
- Ao iniciar um projeto com Go devemos criar um modulo
- Fazendo isso esta sendo dado opção para trabalhar nesse diretorio não no Go Path

Exemplo
go mod init github.com/luizffdemoraes/curso-go/7-Packaging/1-Introduzindo-go-mod

- Para poder habilitar o export de metodos, struct, atributos e parametros é necessário utilizar letra inicial Maiúscula

Avalia o código e imports e baseado nisso e baixa os pacotes ou realiza remoção.
go mod tidy 

Ele baixa o pacote solicitado
go get github.com/google/uuid

Adiciona as dependencias utilizadas no projeto
go.mod

Garante as versões do seu projeto
go.sum 