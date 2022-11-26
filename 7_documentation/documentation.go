// Документирование пакета
// Видно все, что примыкает к объявлению

// This package is for learning go comments and documentation
package main

// Каждый экспортируемый элемент должен быть задокументирован
// По конвенции, каждый коммент - законченное предложение с пунктуацией.

// Customer is a customer representation.
type Customer struct{}

// ID return customer identifier.
func (c Customer) ID() string { return "" }

// Deprecated: use ID() instead.
func (c Customer) OldID() string { return "" }

func main() {
	c := Customer{}
	c.OldID()
}
