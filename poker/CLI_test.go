package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &StubPlayerStore{}

		cli := NewCLI(playerStore, in)
		cli.PlayPoker()

		assertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &StubPlayerStore{}
		cli := NewCLI(playerStore, in)
		cli.PlayPoker()
		assertPlayerWin(t, playerStore, "Cleo")
	})
}

func assertPlayerWin(t *testing.T, playerStore *StubPlayerStore, winner string) {
	t.Helper()
	if len(playerStore.winCalls) < 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(playerStore.winCalls), 1)
	}
	got := playerStore.winCalls[0]
	if got != winner {
		t.Errorf("didn't record correct winner,got '%s',want '%s'", got, winner)
	}
}
