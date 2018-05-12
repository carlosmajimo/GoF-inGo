package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/decorator"
	"time"
)

func TestDecorator(t *testing.T) {

	var repositorio, repetidor IFetcher

	f := NewFixture("Test del patrón Decorator numero 1", t)
	defer f.Run()

	f.Setup(func() {
		repositorio = &Repositorio{}

		repetidor = &Retrier{
			RetryCount:   5,
			WaitInterval: time.Second,
			Fetcher:      repositorio,
		}

	})

	f.Test("Test Fetch en repositorio", func() {
		dataRepositorio, _ := repositorio.Fetch(Args{"id": "1"})
		f.So("...", dataRepositorio["user"], ShouldEqual, "root")
	})

	f.Test("Test Fetch en repetidor", func() {
		dataRepetidor, _ := repetidor.Fetch(Args{"id": "1"})
		f.So("...", dataRepetidor["user"], ShouldEqual, "root")
	})

	f.Test("Test Fetch en repetidor con fallo", func() {
		_, errorRepetidor := repetidor.Fetch(Args{})
		f.So("...", errorRepetidor, ShouldEqual, "No se han recibido argumentos")
	})
}