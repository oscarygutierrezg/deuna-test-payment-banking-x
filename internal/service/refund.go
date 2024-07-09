package service

import (
	bankingx "payment-banking-x/internal/client/banking-x"
	"payment-banking-x/internal/entity/enums"
	service "payment-banking-x/internal/kafka/producer"
	"payment-banking-x/pkg/dto"
)

type RefundService interface {
	Create(origin dto.PaymentRequest) error
}

func NewRefundService(bankingXClient bankingx.Client, refundProducer service.PaymentProducer) *refundService {
	return &refundService{
		bankingXClient: bankingXClient,
		refundProducer: refundProducer,
	}
}

type refundService struct {
	bankingXClient bankingx.Client
	refundProducer service.PaymentProducer
}

func (s *refundService) Create(req *dto.PaymentRequest) error {
	res, err := s.bankingXClient.CreateRefund(bankingx.MapRefundDTOToRefundRequest(req))
	refundResponse := dto.PaymentResponse{
		TransactionID: req.TransactionID,
		PaymentID:     req.PaymentID,
	}
	if err != nil {
		refundResponse.Status = enums.Failed
		refundResponse.Msg = res.Data.Msg
	} else {
		refundResponse.Status = enums.Cancelled
		refundResponse.RefundID = res.Data.RefundID
	}
	return s.refundProducer.Produce(refundResponse)
}
