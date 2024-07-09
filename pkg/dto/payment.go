package dto

import (
	"payment-banking-x/internal/entity/enums"
	enums2 "payment-banking-x/pkg/dto/enums"
)

type PaymentRequest struct {
	PaymentID     string              `json:"paymentId"`
	TransactionID string              `json:"transactionId"`
	Status        enums.PaymentStatus `json:"status"`
	CardID        string              `json:"cardId"`
	CVC           string              `json:"cvc"`
	ExpiredDate   string              `json:"expiredDate"`
	Amount        float64             `json:"amount"`
	Type          enums2.PaymentType  `json:"type"`
	Currency      string              `json:"currency"`
	Merchant      string              `json:"merchant"`
}

type PaymentResponse struct {
	PaymentID     string              `json:"paymentID"`
	TransactionID string              `json:"transactionID"`
	Status        enums.PaymentStatus `json:"status"`
	Msg           string              `json:"msg"`
	RefundID      string              `json:"refundID"`
}
