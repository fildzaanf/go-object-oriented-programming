package encapsulation

// encapsulation is the concept of restricting direct access
// to the objects data and allowing access only through methods.
//
// it helps protect the data and ensures it is used in a controlled way.
//
// in go, encapsulation is implemented by:
// - using unexported fields (lowercase)
// - providing exported methods (getters/setters) to control access

// Payment represents a payment entity with restricted data access.
// all fields are unexported (lowercase), so they cannot be accessed
// directly from outside the package.
type Payment struct {
	id     string
	method string
	amount float64
	fee    float64
}

// NewPayment is a constructor used to create and initialize a Payment object.
// it takes input data and returns a new payment instance.
func NewPayment(id, method string, amount, fee float64) *Payment {
	return &Payment{
		id:     id,
		method: method,
		amount: amount,
		fee:    fee,
	}
}

// GetAmount is a getter method that returns the payment amount.
// it provides read access to the unexported field 'amount'
// without exposing the field directly.
func (p *Payment) GetAmount() float64 {
	return p.amount
}

// SetAmount is a setter method that updates the payment amount.
// it provides controlled write access to the 'amount' field
// and includes validation to ensure the value is valid.
func (p *Payment) SetAmount(amount float64) {
	if amount >= 0 {
		p.amount = amount
	}
}