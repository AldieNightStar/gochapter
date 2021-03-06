package gochapter

import "strings"

func Parse(counter *Counter, src string) []*Scene {
	lines := strings.Split(
		strings.ReplaceAll(src, "\r", ""),
		"\n",
	)
	scenes := make([]*Scene, 0, 32)
	scene := newScene(counter.DoCount())
	for _, line := range lines {
		if len(strings.Trim(line, " \t")) < 1 {
			if !scene.isEmpty() {
				scenes = append(scenes, scene)
				scene = newScene(counter.DoCount())
			}
			continue
		}
		if cmd := parseComand(line); cmd != nil {
			scene.Commands = append(scene.Commands, cmd)
		} else {
			scene.Lines = append(scene.Lines, line)
		}
	}
	if !scene.isEmpty() {
		scenes = append(scenes, scene)
	}
	return scenes
}

func ParseFew(counter *Counter, srcs ...string) []*Scene {
	scenes := make([]*Scene, 0, 64)
	for _, src := range srcs {
		scenes = append(scenes, Parse(counter, src)...)
	}
	return scenes
}
