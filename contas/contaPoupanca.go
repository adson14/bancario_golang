package contas

import "bancario_golang/clientes"

type ContaPoupanca struct {
	Titular                                clientes.Titular
	Agencia, Conta, Digito_conta, Operacao int
	saldo                                  float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	permiteSacar := valorSaque > 0 && valorSaque <= c.saldo

	if permiteSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {

	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso o saldo atual é:", c.saldo
	} else {
		return "O valor para depósito não é válido saldo: ", c.saldo
	}

}

func (c *ContaPoupanca) Extrato() float64 {
	return c.saldo
}
