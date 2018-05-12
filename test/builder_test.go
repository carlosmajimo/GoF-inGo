package test

import (
	"testing"
	. "github.com/mdwhatcott/gounit"
	. "GoF-inGo/project/builder"
)

func TestBuilder(t *testing.T) {

	var sender *Sender

	f := NewFixture("Test del patrón Builder", t)
	defer f.Run()

	f.Setup(func() {
		sender = &Sender{}
	})

	f.Test("Test mensaje JSON", func() {
		_, error := sender.BuildMessage(&JSONMessageBuilder{})
		f.So("...", error, ShouldEqual, nil)
	})

	f.Test("Test mensaje XML", func() {
		_, error := sender.BuildMessage(&XMLMessageBuilder{})
		f.So("...", error, ShouldEqual, nil)
	})
}