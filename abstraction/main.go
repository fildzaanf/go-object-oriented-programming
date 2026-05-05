package abstraction

// abstraction is the concept of hiding implementation details
// and exposing only the essential behavior of an object.
//
// it allows us to define what an object can do
// without specifying how it does it.
//
// in go, abstraction is implemented using interfaces.

// Payment defines the abstraction of a payment behavior.
//
// it declares what actions are available without
// exposing the implementation details.
type Payment interface {
	TotalPayment() float64
	UpdateAmount(amount float64)
}

// EWallet represents a specific type of payment.
type EWallet struct {
	amount float64
	fee    float64
}

// NewEWallet creates a new EWallet object.
func NewEWallet(amount, fee float64) *EWallet {
	return &EWallet{
		amount: amount,
		fee:    fee,
	}
}

// TotalPayment calculates the total payment amount.
func (e *EWallet) TotalPayment() float64 {
	return e.amount + e.fee
}

// UpdateAmount updates the payment amount.
func (e *EWallet) UpdateAmount(amount float64) {
	if amount >= 0 {
		e.amount = amount
	}
}

// Cash represents another type of payment.
type Cash struct {
	amount float64
}

// NewCash creates a new cash object.
func NewCash(amount float64) *Cash {
	return &Cash{
		amount: amount,
	}
}

// TotalPayment processes payment for cash.
func (c *Cash) TotalPayment() float64 {
	return c.amount
}

// UpdateAmount updates the cash amount.
func (c *Cash) UpdateAmount(amount float64) {
	if amount >= 0 {
		c.amount = amount
	}
}

// ExampleAbstraction demonstrates how abstraction works.
func ExampleAbstraction() {
	var payment Payment

	payment = NewEWallet(100000, 2500)

	// accessible through abstraction
	payment.UpdateAmount(120000)
	println(payment.TotalPayment())

	payment = NewCash(50000)

	// accessible through abstraction
	payment.UpdateAmount(60000)
	println(payment.TotalPayment())

}