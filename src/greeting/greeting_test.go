package main

import "testing"

func TestGreeting(t *testing.T) {
	texto := greeting("Teste")
	if texto != "<b>Teste</b>" {
		t.Error("Teste incorreto")
	}
}
