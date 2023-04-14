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

Caso seja necessário trabalhar com dois modulos locais e possivel utilizar esse exemplo para que procure local.:

go mod edit -replace github.com/luizffdemoraes/curso-go/7-Packaging/3/math=../math

Obs.: Ao subir em um repositório na nuvem pois não existe url relativa na pratica ele é uma solução local.

Criação do go.work, ele informa no arquivos as url relativas e podemos colocar ele no gitignore deve ser criado na raiz das pastas que detem o go.mod. O go.work é algo independente do sistema ao adicionalo não será possível utilizar o comando go mod tidy para baixar dependencia. Pasolucionar essa ocorrência temos as seguintes opções.:

1 go get github.com/google/uuid 
2 go mod tidy -e
3 publicar a aplicação para não utilizar o go workspaces

go work init ./math ./sistema

No processo de criação de teste deve-se utilizar a nomenclatura para cada metodo de validação implementado com a inicial.:
"Test"

Verificar cobertura de código.:
go test -coverprofile=coverage

Rodar test unitários.:
go test

Verificar ponto onde falta a cobertura gerando html apresentado.:
go tool cover -html=coverage

Rodar testes de BenchMarck.:
go test -bench=.

Rodar somente teste Benchmark.:
go test -bench=. -run=^#

Executar de 10 em 10.:
go test -bench=. -run=^# -count=10

Executar por 3 segundos apenas por operação.:
go test -bench=. -run=^# -count=10 -benchtime=3s

Pegar a documentação do Go Test.:
go help test