package gochapter

type Scene struct {
	Id    int
	Lines []string
}

func newScene(id int) *Scene {
	return &Scene{Lines: make([]string, 0, 8), Id: id}
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
