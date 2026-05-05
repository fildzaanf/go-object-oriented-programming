package polymorphism

// polymorphism is the ability of different objects
// to be treated as the same type through an interface.
//
// it allows one interface to be used with different implementations.
//
// in go, polymorphism is achieved using interfaces.

// Payment interface defines a common behavior for all payment types.
type Payment interface {
	TotalPayment() float64
}

// EWallet represents a digital wallet payment.
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

// TotalPayment implements the Payment interface for EWallet.
func (e *EWallet) TotalPayment() float64 {
	return e.amount + e.fee
}

// Cash represents a cash payment.
type Cash struct {
	amount float64
}

// NewCash creates a new Cash object.
func NewCash(amount float64) *Cash {
	return &Cash{
		amount: amount,
	}
}

// TotalPayment implements the Payment interface for Cash.
func (c *Cash) TotalPayment() float64 {
	return c.amount
}

// ProcessPayment demonstrates polymorphism.
//
// it accepts any object that implements the Payment interface
// and processes it without knowing its concrete type.
func ProcessPayment(p Payment) {
	println(p.TotalPayment())
}

// ExamplePolymorphism demonstrates how polymorphism works.
func ExamplePolymorphism() {

	ewallet := NewEWallet(100000, 2500)
	cash := NewCash(50000)

	// same function, different objects
	ProcessPayment(ewallet)
	ProcessPayment(cash)
}