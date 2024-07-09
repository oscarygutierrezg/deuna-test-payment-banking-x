package banking_x

import (
	"payment-banking-x/internal/entity/enums"
	"time"
)

type Payment struct {
	PaymentID   string              `json:"cardId"`
	Status      enums.PaymentStatus `json:"status"`
	CreatedDate time.Time           `json:"createdDate"`
	UpdatedDate time.Time           `json:"updatedDate"`
}
