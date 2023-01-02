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

}