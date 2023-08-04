package main

import (
	"fmt"

	"bancario_golang/clientes"
	"bancario_golang/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

//A interface obriga a implementação dos métodos presentes
type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	fmt.Println("=========Conta da Critina=======")
	contaCritina := new(contas.ContaCorrente)
	contaCritina.Agencia = 1234
	contaCritina.Conta = 545455
	contaCritina.Depositar(12)
	fmt.Println(contaCritina.Extrato()) // Ponteiros
	fmt.Println(contaCritina.Sacar(100))
	status, valor := contaCritina.Depositar(-100)
	fmt.Println(status, valor)
	fmt.Println("=================================")

	fmt.Println("=========Conta do Bruno=======")
	contaDoBruno := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Souza",
		CPF:       "123.222.154-54",
		Profissao: "Programdor"},
		Agencia:      1234,
		Conta:        1245455,
		Digito_conta: 1,
	}

	contaDoBruno.Depositar(120)

	fmt.Println(contaDoBruno)
	fmt.Println("=================================")

	fmt.Println("=========Conta do Danilo=======")
	contaDanilo := contas.ContaPoupanca{}
	contaDanilo.Depositar(100)
	contaDanilo.Operacao = 13
	contaDanilo.Agencia = 0012
	contaDanilo.Conta = 1124
	contaDanilo.Digito_conta = 1
	contaDanilo.Titular = clientes.Titular{Nome: "Danilo Souza",
		CPF:       "123.222.174-54",
		Profissao: "Design"}

	PagarBoleto(&contaDanilo, 110)
	fmt.Println(contaDanilo)
}
