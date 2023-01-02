package main

import "fmt"

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

// Filter by color
func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}
	return result
}

// break open–closed principle (OCP)
// Filter by size
func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// break open–closed principle (OCP)
// Filter by size and color method
func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)
	
	for i, v := range products {
		if v.color == color && v.size == size {
			result = append(result, &products[i])
		}
	}
	return result
}

// Follow OCP, make extendable setup by using Specification pattern

// Implement specification interface
// Test whether or not a product specified, satisfies some criteria
type Specification interface {
	IsSatisfied(p *Product) bool
}

// To check for color make, ColorSpecification and specify color you want to filter on
type ColorSpecification struct {
	color Color
}

// Check whether or not a product matches a color specification
func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

// To check for size make, SizeSpecification and specify size you want to filter on
type SizeSpecification struct {
	size Size
}

// Check whether or not a product matches a size specification
func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

// Unlikely to ever modify
type BetterFilter struct {}

// Filter product based on specification interface
func (f *BetterFilter) Filter(product []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range product {
		if spec.IsSatisfied(&v) {
			result = append(result, &product[i])
		}
	}
	return result
}

// CompositeSpecification type
type AndSpecification struct {
	first, second Specification
}

// CompositeSpecification : Combines two different specification
func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.first.IsSatisfied(p) && a.second.IsSatisfied(p)
}


func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old) :\n")
	f := Filter{}
	for  _, v := range  f.FilterByColor(products, green) {
		fmt.Printf(" - %s is green\n", v.name)
	}
	
	// This approach gives more flexibility, because of we want to filter by a paticular new type
	// all we have to do is make new specification and make sure it confirms to specification interface...

	// The interface types is open for extension (we can implement this interface)
	// its closed for modification, unlikely to ever modify specification interface
	// similar fashion unlikely to modify BetterFilter

	fmt.Printf("Green products (new) :\n")
	greenSpec := ColorSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green\n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large Green products :\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf("- %s is large and green\n", v.name)
	}
}