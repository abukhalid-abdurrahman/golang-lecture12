package pkg

type PaymentType string

const (
	CREDIT PaymentType = "CREDIT"
	DEBIT PaymentType = "DEBIT"
	P2P PaymentType = "P2P"
)

type Payment struct {
	Type PaymentType
	Amount float64
	Sender string
	Reviever string
}