package payment

type Payment string

const (
	MONEY  Payment = "money"
	CREDIT Payment = "credit"
	DEBIT  Payment = "debit"
	PIX    Payment = "pix"
)
