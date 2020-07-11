package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyweight(t *testing.T) {
	viewer1 := NewViewer("1")
	viewer2 := NewViewer("1")

	viewer1.Display()
	viewer2.Display()

	fmt.Println(viewer1.Flyweight == viewer2.Flyweight)
}
