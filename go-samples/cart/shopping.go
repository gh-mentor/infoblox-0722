/*
This application demonstrates the use of methods with structs in Go.
The application creates a shopping cart and a cashier. The shopping cart
is a slice of items and the cashier is a struct with a method to checkout
the items in the cart.

The application creates a shopping cart and adds three items to the cart.
The application then checks out the items in the cart.
*/


package main

import (
	"fmt"
)

// Item type descibes an item in the cart
type Item struct{
	name string
	price float32
} 

// Cart type is a slice of items
type Cart []Item

// AddItem method adds an item to the cart
func (cart *Cart) AddItem(item Item){
	
	*cart = append(*cart, item)
}

// Cashier struct type describes a cashier
type Cashier struct{}

// Checkout method checks out the items in the cart
func (register Cashier) Checkout(cart Cart){
	for index, item := range cart{
		fmt.Printf("Item #%d %s $%5.2f\n", index + 1, item.name, item.price)
	}
}

/*
The main function creates a shopping cart and a cashier. The application
adds three items to the cart and then checks out the items in the cart.
*/
func main(){
	var cart = make(Cart, 0, 10)
	var cashier = Cashier{}

	cart.AddItem(Item{name: "milk", price: 1.50})
	cart.AddItem(Item{name: "butter", price: 2.00})
	cart.AddItem(Item{name: "bread", price: 3.75})

	cashier.Checkout(cart)
}