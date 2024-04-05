package main

import (
	"testing"
)

func TestMainConnected(t *testing.T) {
	//time.Sleep(2 * time.Second)

	teste := Teste{
		TesteDuracao: 3,
		Msg:          "{0 2024-04-05 10:29:17.0535408 -0300 -03 m=+0.002230401 PM2.5 35.2}",
	}
	publiStart(teste)
}
