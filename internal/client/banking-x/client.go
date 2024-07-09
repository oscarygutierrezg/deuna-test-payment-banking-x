package banking_x

import (
	"payment-banking-x/internal/client/http"
)

type Client interface {
	CreatePayment(req *PaymentRequest) (*PaymentResponse, error)
	CreateRefund(req *RefundRequest) (*RefundResponse, error)
}

func NewClient(restClient http.RestClient) *client {
	return &client{
		restClient: restClient,
	}
}

type client struct {
	restClient http.RestClient
}

func (c *client) CreatePayment(req *PaymentRequest) (*PaymentResponse, error) {
	res := PaymentResponse{
		Data: PaymentData{},
	}
	respPost, err := c.restClient.Post("payments", req, &res)
	if err != nil {
		if errorDTO, ok := respPost.(http.ErrorDTO); ok {
			res.Data.Msg = errorDTO.Data
		}
	}
	return &res, err
}

func (c *client) CreateRefund(req *RefundRequest) (*RefundResponse, error) {
	res := RefundResponse{
		Data: RefundData{},
	}
	respPost, err := c.restClient.Post("refunds", req, &res)
	if err != nil {
		if errorDTO, ok := respPost.(http.ErrorDTO); ok {
			res.Data.Msg = errorDTO.Data
		}
	}
	return &res, nil
}
