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


// Adhere to ISP  
// Try to break Interface into separate parts, that people will definitely need
// Here we have  very granular kind of definitions
// Just grab the interface you need and you dont have any extra members in those interfaces
// For example if you're building an ordinary printer, just get the print method
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {

}

func (m MyPrinter) Print(d Document) {

}

type PhotoCopier struct {

}

func (p PhotoCopier) Print(d Document) {

}

func (p PhotoCopier) Scan(d Document) {

}


type MultiFunctionDevice interface {
	Printer
	Scanner
}

type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d) // reuse the functionality of printer that you already have
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d) // reuse the functionality of printer that you already have
}

func main() {

}