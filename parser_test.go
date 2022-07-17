package gochapter

import "testing"

func TestSplitArgs(t *testing.T) {
	src := "a,b\\,c,d"
	args := splitArgs(src)
	if len(args) != 3 {
		t.Fatal("Len not 3")
	}
	if args[0] != "a" || args[1] != "b,c" || args[2] != "d" {
		t.Fatal("Data has to be: a :: b,c :: d")
	}
}

func TestSplitArgsOneArg(t *testing.T) {
	src := "a"
	args := splitArgs(src)
	if len(args) != 1 {
		t.Fatal("Len not 1")
	}
	if args[0] != "a" {
		t.Fatal("One args is not 'a'")
	}
}

func TestCommandParse(t *testing.T) {
	src := ".say Hello, day"
	cmd := parseComand(src)
	if cmd.Name != "say" {
		t.Fatal("Command is not 'say'")
	}
	if cmd.Args[0] != "Hello" {
		t.Fatal("Arg 0 is not 'Hello'")
	}
	if cmd.Args[1] != "day" {
		t.Fatal("Arg 1 is not 'day'")
	}
}

func TestCommandParseNoArgs(t *testing.T) {
	src := ".done"
	cmd := parseComand(src)
	if cmd.Name != "done" {
		t.Fatal("Command is not 'done'")
	}
	if len(cmd.Args) != 0 {
		t.Fatal("Args is present. Even if there are not specified")
	}
}
