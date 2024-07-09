package banking_x

import (
	"payment-banking-x/pkg/dto"
	"time"
)

type RefundResponse struct {
	Data RefundData `json:"data"`
}

type RefundData struct {
	TransactionID string    `json:"transactionId"`
	Msg           string    `json:"msg"`
	RefundID      string    `json:"refundId"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	Timestamp     time.Time `json:"timestamp"`
}

type RefundRequest struct {
	TransactionID string  `json:"transactionId"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
}

func MapRefundDTOToRefundRequest(req *dto.PaymentRequest) *RefundRequest {
	return &RefundRequest{
		TransactionID: req.TransactionID,
		Amount:        req.Amount,
		Currency:      req.Currency,
	}
}
