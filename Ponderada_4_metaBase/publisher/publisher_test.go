package main

import (
	"testing"
)

func TestMainConnected(t *testing.T) {
	test := Teste{tested: true, sec: 3}
	s := NewSensor("sensor", -100, 100)
	s.On(test)
}
