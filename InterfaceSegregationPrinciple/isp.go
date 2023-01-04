package main

type Document struct {
}

// Breaking ISP: by putting too much into an interface
// Here we put all Print, Fax and Scan into a single interface and 
// then we expect everyone to implement even if they dont have this functionality

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {

}

func (m MultiFunctionPrinter) Fax(d Document) {

}

func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionPrinter struct {
}

func (o OldFashionPrinter) Print(d Document) {

}

// Deprecated: ...
func (o OldFashionPrinter) Fax(d Document) {
	panic("Operation not supported")
}

// Deprecated: ...
func (o OldFashionPrinter) Scan(d Document) {
	panic("Operation not supported")
}

func main() {

}