package service

type Services struct {
	Payment *paymentService
	Refund  *refundService
}
