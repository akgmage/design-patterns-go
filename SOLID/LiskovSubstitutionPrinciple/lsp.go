// Liskov Substitution Principle
// It states that if you continue to use generalization like interfaces
// then you should not have implementors of those generalizations break some
// of the assumptions which are set up at the higher level.
package main

import (
	"fmt"
)

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
func (r *Rectangle) SetWidth(width int) {
	r.width = width
}
func (r *Rectangle) GetHeight() int {
	return r.height
}
func (r *Rectangle) SetHeight(height int) {
	r.height = height
}


type Square struct {
	Rectangle
}

// Enforce the idea of width being equal to height
func NewSquare(size int) *Square {
	sq := Square{}
	sq.width = size
	sq.height = size
	return &sq
}

// Breaking assumption by setting both width and height 
// violate LSP
func (s * Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

// violate LSP
func (s * Square) SetHeight(height int) {
	s.height = height
	s.width = height
}

// If you are expecting some sort of something up the hierarchy, in the argument
// then it should continue to work, even if you proceed to extend objects which 
// are already Sized type.
// Here we took a rectangle and we decided to sort of extend the rectangle and make it a square
// it should continue to work but it doesn't, and if we try to plug in the square 
// we can see how it goes wrong
// Assumption: Here we assume that if we have a Sized interface and you set its height
// then you're just setting its height, not both the height and width

func UseIt(sized Sized) {
	width := sized.GetWidth()
	// The call to SetHeight actually sets not just the height, it also sets the width
	// so the internal width of the square becomes inconsistent with the value of variable  
	// As a result we will get different values for expected and actual area
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

type Square2 struct {
	size int // width, height
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rc := &Rectangle{10, 20}
	UseIt(rc)

	sq := NewSquare(6)
	UseIt(sq)

	// Outputs
	// Expected an area of 100, but got 100
	// Expected an area of 60, but got 100

	
}	