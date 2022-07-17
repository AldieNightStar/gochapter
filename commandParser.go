package gochapter

import "strings"

func parseComand(line string) *Command {
	if len(line) < 2 || strings.HasPrefix(line, "..") || !strings.HasPrefix(line, ".") {
		return nil
	}
	line = line[1:]
	if !strings.Contains(line, " ") {
		return &Command{Name: line, Args: []string{}}
	}
	lineCommands := strings.SplitN(line, " ", 2)
	if len(lineCommands) < 2 {
		return &Command{Name: line, Args: []string{}}
	}
	return &Command{Name: lineCommands[0], Args: splitArgs(lineCommands[1])}
}

func splitArgs(line string) []string {
	sb := strings.Builder{}
	sb.Grow(32)
	args := make([]string, 0, 32)
	esc := false
	for _, c := range line {
		if esc {
			esc = false
			sb.WriteRune(c)
			continue
		}
		if c == '\\' {
			esc = true
			continue
		}
		if c == ',' {
			args = append(args, sb.String())
			sb = strings.Builder{}
			sb.Grow(32)
			continue
		}
		sb.WriteRune(c)
	}
	if sb.Len() > 0 {
		args = append(args, strings.Trim(sb.String(), " \t"))
	}
	return args
}
