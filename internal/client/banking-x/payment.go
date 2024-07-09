package banking_x

import (
	"payment-banking-x/pkg/dto"
	"time"
)

type PaymentResponse struct {
	Data PaymentData `json:"data"`
}

type PaymentData struct {
	TransactionID string    `json:"transactionId"`
	Msg           string    `json:"msg"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Merchant      string    `json:"merchant"`
	Timestamp     time.Time `json:"timestamp"`
}

type PaymentRequest struct {
	CardID      string  `json:"cardId"`
	CVC         string  `json:"cvc"`
	ExpiredDate string  `json:"expiredDate"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Merchant    string  `json:"merchant"`
}

func MapPaymentDTOToPaymentRequest(req *dto.PaymentRequest) *PaymentRequest {
	return &PaymentRequest{
		CardID:      req.CardID,
		CVC:         req.CVC,
		ExpiredDate: req.ExpiredDate,
		Amount:      req.Amount,
		Currency:    req.Currency,
		Merchant:    req.Merchant,
	}
}
