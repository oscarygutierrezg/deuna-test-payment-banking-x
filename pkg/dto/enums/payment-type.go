package enums

import (
	"encoding/json"
	"errors"
)

type PaymentType string

const (
	Payment = "Payment"
	Refund  = "Refund"
)

var statusToString = map[PaymentType]string{
	Payment: "Payment",
	Refund:  "Refund",
}

var stringToStatus = map[string]PaymentType{
	"Payment": Payment,
	"Refund":  Refund,
}

func (s PaymentType) String() string {
	return statusToString[s]
}

func (s *PaymentType) UnmarshalJSON(data []byte) error {
	var statusStr string
	if err := json.Unmarshal(data, &statusStr); err != nil {
		return err
	}

	status, ok := stringToStatus[statusStr]
	if !ok {
		return errors.New("invalid status value")
	}

	*s = status
	return nil
}
