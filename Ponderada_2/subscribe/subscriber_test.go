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
	if *ponteiro == "nil24613#2335" {
		t.Fatalf("Menssagem não recebida")

	}
	select {}
}

// TestValidacaoDados garante que os dados enviados pelo simulador chegam sem alterações.
func TestValidacaoDados(t *testing.T) {
	var momento_inicial string = "nil24613#2335" // Uma sequência de caracteres contrá casualidade
	ponteiro := &momento_inicial
	fmt.Println("Valor apontado pelo ponteiro:", ponteiro)
	teste := Teste{emTeste: true, testeDuracao: 10, textoTeste: ponteiro}
	mqttStart(teste)
	if *ponteiro != "MSG Teste" {
		t.Fatalf("A mensagem teste não foi validada! Esperado %s, o que veio foi %s", *ponteiro, "MSG Teste")

	}
}
