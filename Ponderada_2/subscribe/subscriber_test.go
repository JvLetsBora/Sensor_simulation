package main

import (
	"fmt"
	"testing"
)

// TestRecebimento garante que os dados enviados pelo simulador são recebidos pelo broker.
func TestRecebimento(t *testing.T) {
	var momento_inicial string = "nil24613#2335" // Uma sequência de caracteres contrá casualidade
	ponteiro := &momento_inicial
	fmt.Println("Valor apontado pelo ponteiro:", ponteiro)
	teste := Teste{emTeste: true, testeDuracao: 10, textoTeste: ponteiro}
	mqttStart(teste)
	if *ponteiro != "MSG Teste" {
		t.Fatalf("A mensagem teste não foi validada! Esperado %s, o que veio foi %s", *ponteiro, "MSG Teste")

	}
	select {}
}

// TestValidacaoDados garante que os dados enviados pelo simulador chegam sem alterações.
func TestValidacaoDados(t *testing.T) {

}

// TestConfirmacaoTaxaDisparo garante que o simulador atende às especificações de taxa de disparo de mensagens dentro de uma margem de erro razoável.
func TestConfirmacaoTaxaDisparo(t *testing.T) {
	// Escreva testes aqui
}
