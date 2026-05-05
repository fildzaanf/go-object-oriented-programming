package inheritance

// inheritance in object-oriented programming allows a class
// to inherit properties (attributes) and behavior (methods) from another class.
//
// go does not support inheritance directly.
// instead, it uses composition (embedding) to achieve similar behavior.
//
// in go, inheritance-like behavior is achieved through:
// - embedding with reuse fields and methods
// - access embedded methods directly
// - define new methods on child struct

// BasePayment represents a base structure for payment.
// it contains common fields that can be shared
// by other payment types.
type BasePayment struct {
	ID     string
	Amount float64
}

// GetAmount returns the payment amount.
// this method can be reused by other structs
// that embed BasePayment.
func (b *BasePayment) GetAmount() float64 {
	return b.Amount
}

// EWallet represents a specific type of payment.
// it embeds BasePayment to reuse its fields and methods.
type EWallet struct {
	BasePayment // embedding (similar to inheritance)
	Fee         float64
}

// TotalPayment defines a custom behavior for EWallet.
// even though EWallet reuses BasePayment,
// it can still have its own specific logic.
func (e *EWallet) TotalPayment() float64 {
	return e.Amount + e.Fee
}

// Cash represents another type of payment.
// it also embeds BasePayment to reuse shared data.
type Cash struct {
	BasePayment
}

// TotalPayment defines behavior specific to cash.
func (c *Cash) TotalPayment() float64 {
	return c.Amount
}

// ExampleInheritance demonstrates inheritance-like behavior in go.
func ExampleInheritance() {

	ewallet := EWallet{
		BasePayment: BasePayment{
			ID:     "PAY001",
			Amount: 100000,
		},
		Fee: 2500,
	}

	cash := Cash{
		BasePayment: BasePayment{
			ID:     "PAY002",
			Amount: 50000,
		},
	}

	// embedding with reuse field from BasePayment
	println(ewallet.ID)

	// call BasePayment method directly from EWallet
	println(ewallet.GetAmount())

	// each struct has its own implementation
	println(ewallet.TotalPayment())
	println(cash.TotalPayment())
}