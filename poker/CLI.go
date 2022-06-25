package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	PlayerStore PlayerStore
	in          *bufio.Scanner
}

func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		PlayerStore: store,
		in:          bufio.NewScanner(in),
	}
}

func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.PlayerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
