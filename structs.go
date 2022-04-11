package structs

import (
	"fmt"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip:   0,
	}

	return b
}

//this extra set of parentheses says it only works with bill objects
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	//total. -25 makes the string 25 characters long. +25 would make spacing on the left of those characters.
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// func (b *bill) save() {
// 	data := []byte(b.format())

// 	//0644 is permissions
// 	// err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	fmt.Println("bill was saved to file")
// }
