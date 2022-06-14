package hello

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("lh", "")
		want := "Hello, lh"
		assertMsgEqual(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := englishHelloPrefix + "World"
		assertMsgEqual(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		//西班牙语
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertMsgEqual(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("lh", "French")
		want := "Bonjour, lh"
		assertMsgEqual(t, got, want)
	})
	t.Run("in chinese", func(t *testing.T) {
		got := Hello("吕欢", "Chinese")
		want := "你好，吕欢"
		assertMsgEqual(t, got, want)
	})
}

func assertMsgEqual(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
