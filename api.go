package gochapter

import "strings"

func Parse(src string, counter *Counter) []*Scene {
	lines := strings.Split(
		strings.ReplaceAll(src, "\r", ""),
		"\n",
	)
	scenes := make([]*Scene, 0, 32)
	scene := newScene(counter.DoCount())
	for _, line := range lines {
		if len(strings.Trim(line, " \t")) < 1 {
			scenes = append(scenes, scene)
			scene = newScene(counter.DoCount())
			continue
		}
		scene.Lines = append(scene.Lines, line)
	}
	return scenes
}

func ParseFew(counter *Counter, srcs ...string) []*Scene {
	scenes := make([]*Scene, 0, 64)
	for _, src := range srcs {
		addAll(&scenes, Parse(src, counter))
	}
	return scenes
}
