// Dependency Inversion Principle
// The principle states

// 1) High-level modules should not import anything from low-level modules. Both should depend on abstractions (e.g., interfaces).
// 2) Abstractions should not depend on details. Details (concrete implementations) should depend on abstractions.

package main

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

// Infrastructure  to have some sort of relationships
// For example: A is the Parent of B
type Info struct {
	from *Person
	relationship Relationship
	to *Person
}

// store above information
// have all data about the relationship between different people stored
// in some sort of type 
type Relationships struct {
	relations []Info
}

func (r Relationships) AddParentAndChild(parent , child *Person) {
}

func main() {

}