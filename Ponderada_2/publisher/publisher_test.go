package main

import (
	"testing"
)

func TestMainConnected(t *testing.T) {
	//time.Sleep(2 * time.Second)
	teste := Teste{testeDuracao: 3}
	publiStart(teste)
}
