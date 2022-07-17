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

func TestAllScenes(t *testing.T) {
	src := "Line1\nLine2\n\nLine1\n.command a, b\n\n\n\n\n\nLine1"
	scenes := Parse(NewCounter(1), src)
	if len(scenes) != 3 {
		t.Fatal("Scenes len need to be 3")
	}
	if scenes[0].Lines[0] != "Line1" || scenes[0].Lines[1] != "Line2" {
		t.Fatal("First scene lines is wrong")
	}
	if scenes[1].Lines[0] != "Line1" {
		t.Fatal("Second scene line 0 is wrong")
	}
	if scenes[1].Commands[0].Name != "command" || scenes[1].Commands[0].Args[0] != "a" || scenes[1].Commands[0].Args[1] != "b" {
		t.Fatal("Second scene command is bad!")
	}
	if scenes[2].Lines[0] != "Line1" {
		t.Fatal("Third scene line 0 is wrong")
	}
}
