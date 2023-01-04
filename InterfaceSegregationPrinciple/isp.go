package main

type Document struct {

}

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

func (o OldFashionPrinter) Fax(d Document) {

}

func (o OldFashionPrinter) Scan(d Document) {

}

func
 main() {

}