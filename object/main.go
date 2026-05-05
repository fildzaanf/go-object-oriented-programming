package object

import "fmt"

// object is an instance of a class.
//
// it represents a real example created from a class
// and contains actual data based on defined attributes.
//
// in go, an object is an instance of a struct.


// Payment represents a payment entity.
// it contains fields that describe the data of a payment object.
type Payment struct {
	ID        string
	Method    string  
	Amount    float64
	Fee       float64
}

// constructor is a common pattern used to create and initialize an object.
//
// NewPayment is a constructor used to create and initialize a payment object.
// it takes input data and returns a new Payment instance.
func NewPayment(id, method string, amount, fee float64) Payment {
	return Payment{
		ID:     id,
		Method: method,
		Amount: amount,
		Fee:    fee,
	}
}

// TotalPayment is a method that represents the behavior of the payment object.
// it calculates the total payment amount using the objects data. 
func (p *Payment) TotalPayment() float64 {
	return p.Amount + p.Fee
}

// ExampleObject demonstrates how to create and use an object.
// it shows how an object is created using a constructor
// and how its method (behavior) is called.
func ExampleObject() {

	// create an object (instance) using constructor
	payment := NewPayment("PAY001", "e-wallet", 100000, 2500)

	// call method (behavior) of the object
	total := payment.TotalPayment()
	fmt.Println(total)
}