package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/prototype"
)

func TestPrototype(t *testing.T) {

	var Director Node
	var Clon Node

	f := NewFixture("Test del patron Prototype", t)
	defer f.Run()

	f.Setup(func() {
		Director = NewDirector()
		Director.AddChild(NewIngeniero())

		Clon = Director.Clone()
	})

	f.Test("Test de clonacion de Director", func() {
		f.So("...", Director.String(), ShouldEqual, Clon.String())
	})

}