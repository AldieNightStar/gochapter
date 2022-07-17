package gochapter

type Command struct {
	Name string
	Args []string
}

type Scene struct {
	Id       int
	Lines    []string
	Commands []*Command
}

func newScene(id int) *Scene {
	return &Scene{
		Id:       id,
		Lines:    make([]string, 0, 8),
		Commands: make([]*Command, 0, 8),
	}
}

type Counter struct {
	Count int
}

func NewCounter(start int) *Counter {
	return &Counter{Count: start}
}

func (c *Counter) DoCount() (result int) {
	result = c.Count
	c.Count += 1
	return
}
