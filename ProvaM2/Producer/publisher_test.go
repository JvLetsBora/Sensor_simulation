package main

import (
	"math/rand"
	"testing"
	"time"
)

func NovoSensor(TipoPoluente string, id string) Msg {
	return Msg{
		SensorId:     id,
		Timestamp:    time.Now().Unix(),
		TipoPoluente: TipoPoluente,
		Nivel:        rand.Intn(int(30)) + int(8),
	}
}

func TestMainConnected(t *testing.T) {
	//time.Sleep(2 * time.Second)

	teste := Teste{
		TesteDuracao: 1,
		Msg:          NovoSensor("PM2.5", "0"),
	}
	publiStart(teste)
	teste = Teste{
		TesteDuracao: 1,
		Msg:          NovoSensor("PM2.5", "0"),
	}
	publiStart(teste)
	teste = Teste{
		TesteDuracao: 1,
		Msg:          NovoSensor("PM2.5", "0"),
	}
	publiStart(teste)
	teste = Teste{
		TesteDuracao: 1,
		Msg:          NovoSensor("PM2.5", "0"),
	}
	publiStart(teste)
	teste = Teste{
		TesteDuracao: 1,
		Msg:          NovoSensor("PM2.5", "0"),
	}
	publiStart(teste)

}
