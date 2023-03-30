module curso-go

go 1.20

// Criar modulo
// go mod init curso-go
// Baixar pacote
// go get golang.org/x/exp/constraints
// verifica um pacote que esta sendo utilizado e não foi baixado
// caso não tenha sido usado realiza remoção
// go mod tidy

require (
	github.com/google/uuid v1.3.0 // indirect
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29 // indirect
)
