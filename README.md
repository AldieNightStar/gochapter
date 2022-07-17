# Go Chapter
### Golang chapter parser for future novel engines

# Import
```go
import "github.com/AldieNightStar/gochapter"
```

# Usage
```go
// First of all - create Counter (It will count ID's)
cnt := gochapter.NewCounter(1)

// Parse single text (src)
// We will have *Scene list
scenes := gochapter.Parse(src, counter)

// Parse few texts
// It will also return *Scene list
scenes := gochapter.ParseFew(counter, src, src2, src3)
```

# Structures
* `Scene`
```go
type Scene struct {
	Id       int
	Lines    []string
	Commands []*Command
}
```
* `Command`
```go
type Command struct {
	Name string
	Args []string
}
```

# Scene view
* This is how scenes looks like
* Commands starts with `.`, but not `..`
```
This is first Scene
This is second line of first scene

This is second Scene
This is the second line of the second Scene
And third line

This is third scene
.background red
.animations stop




Last scene
.image 01.png
```