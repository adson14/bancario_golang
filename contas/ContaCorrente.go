package contas

import "bancario_golang/clientes"

// Struct é um tipo que vai conter varios outros tipo
//Para tornar visível para outras parte a primeira letra deve ser Maiusculo
type ContaCorrente struct {
	Titular      clientes.Titular
	Agencia      int
	Conta        int
	Digito_conta int
	saldo        float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	permiteSacar := valorSaque > 0 && valorSaque <= c.saldo

	if permiteSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {

	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso o saldo atual é:", c.saldo
	} else {
		return "O valor para depósito não é válido saldo: ", c.saldo
	}

}

func (c *ContaCorrente) Extrato() float64 {
	return c.saldo
}
