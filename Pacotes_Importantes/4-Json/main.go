package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Utilizar - na tag omite a informação no momento de Unmarshal
type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 1}
	//	Transformar uma Conta em Json
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	//	Json sempre retorna em bytes e necessário converter para string
	fmt.Println(string(res))

	//	Ao utilizar o Encoder faz o processo de serialização e retorna em algum lugar
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	jsonPuro := []byte(`{"n":2,"s":200}`)
	var contaX Conta
	//	Referência na memória
	//  Transformar Json em um conta
	err = json.Unmarshal(jsonPuro, &contaX)
	if err != nil {
		panic(err)
	}
	println(contaX.Saldo)
}
