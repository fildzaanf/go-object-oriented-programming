package class

// class is a blueprint or template in object oriented programming 
// used to create objects.
//
// in go there is no "class" keyword.
// structs define data and methods define behavior.
// together they provide class-like functionality.

// structs are used to define attributes (what the object has)
//
// attributes represent the data of an object.
// these attributes are used by methods to handle the objects data.
//
// in go, they are defined as fields in a struct
// and describe the state of the object.
//
// Payment represents a payment entity.
// it contains fields that describe the data of a payment object.
type Payment struct {
	ID        string
	Method    string  
	Amount    float64
	Fee       float64
}

// methods are used to define behavior (what the object can do).
//
// the receiver represents the object that calls the method.
// it allows the method to access and work with the objects data.

// method with value receiver (p Payment)
//
// a value receiver passes a copy of the object to the method.
// any changes made inside the method do not affect the original object.
//
// use value receiver when:
// - the method does not modify the object
// - the struct is small
//
// TotalPayment is a method with a value receiver that calculates
// the total payment amount using the objects data.
func (p Payment) TotalPayment() float64 {
	return p.Amount + p.Fee
}

// method with pointer receiver (p *Payment) 
//
// a pointer receiver passes the address of the object to the method.
// it allows the method to modify the original object.
//
// use pointer receiver when:
// - the method needs to modify the object
// - the struct is large
// - consistency is needed across methods
// 
// UpdateAmount is a method with a pointer receiver.
// it updates the amount of the payment by setting a new value.
func (p *Payment) UpdateAmount(amount float64) {
	p.Amount = amount
}