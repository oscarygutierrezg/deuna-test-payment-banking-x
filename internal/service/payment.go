package service

import (
	"payment-banking-x/internal/client/banking-x"
	"payment-banking-x/internal/entity/enums"
	service "payment-banking-x/internal/kafka/producer"
	"payment-banking-x/pkg/dto"
)

type PaymentService interface {
	Create(origin *dto.PaymentRequest) error
}

func NewPaymentService(bankingXClient banking_x.Client, paymentProducer service.PaymentProducer) *paymentService {
	return &paymentService{
		bankingXClient:  bankingXClient,
		paymentProducer: paymentProducer,
	}
}

type paymentService struct {
	bankingXClient  banking_x.Client
	paymentProducer service.PaymentProducer
}

func (s *paymentService) Create(req *dto.PaymentRequest) error {
	res, err := s.bankingXClient.CreatePayment(banking_x.MapPaymentDTOToPaymentRequest(req))
	paymentResponse := dto.PaymentResponse{
		PaymentID: req.PaymentID,
	}
	if err != nil {
		paymentResponse.Status = enums.Failed
		paymentResponse.Msg = res.Data.Msg
	} else {
		paymentResponse.Status = enums.Approved
		paymentResponse.TransactionID = res.Data.TransactionID
	}
	return s.paymentProducer.Produce(paymentResponse)
}
