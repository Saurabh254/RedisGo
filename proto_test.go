package main

import (
	"fmt"
	"testing"
)

func TestProtocol(t *testing.T) {
	raw := "*3\r\n$3\r\nSET\r\n$10\r\nwhatisthis\r\n$4\r\nbroo\r\n"
	cmd, err := parseCommand(raw)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cmd)
}
