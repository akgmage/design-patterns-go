// Dependency Inversion Principle
// The principle states

// 1) High-level modules should not import anything from low-level modules. Both should depend on abstractions (e.g., interfaces).
// 2) Abstractions should not depend on details. Details (concrete implementations) should depend on abstractions.

package main

import "fmt"

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

// Implement this interface on relationships
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// low-level module
// store above information
// have all data about the relationship between different people stored
// in some sort of type 
type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

// Add relationships
func (r *Relationships) AddParentAndChild(parent , child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// high-level module
type Research struct {
	// Break DIP
	// dependency doesn't have to be on a low level module directly
	// relationships Relationships 
	browser RelationshipBrowser // Depend on abstraction
}

// perform research
func (r *Research) Investigate() {
	// relations := r.relationships.relations
	// for _, rel := range relations {
	// 	if rel.from.name == "John" && rel.relationship == Parent {
	// 		fmt.Println("John has a child called ", rel.to.name)
	// 	}
	// }
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}
	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1) 
	relationships.AddParentAndChild(&parent, &child2)
	r := Research{&relationships}
	r.Investigate()
}