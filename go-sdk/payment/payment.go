package payment

import (
	"fmt"
	"strings"
)

type PaymentDetails struct {
	TransactionID           string
	PaymentMethod           string
	CountryCode             string
	Amount                  float64
	CustomerMobileNumber    string
	PaymentMethodProperties map[string]interface{}
}

type PaymentProcessor interface {
	ProcessPayment(details PaymentDetails) string
	Validate(details PaymentDetails) bool
}

// missing fields
func ValidatePaymentDetails(details PaymentDetails) error {
	var missingFields []string

	fieldValidations := map[string]func(PaymentDetails) bool{
		"TransactionID":                   func(d PaymentDetails) bool { return d.TransactionID != "" },
		"PaymentMethod":                   func(d PaymentDetails) bool { return d.PaymentMethod != "" },
		"CountryCode":                     func(d PaymentDetails) bool { return d.CountryCode != "" },
		"CustomerMobileNumber":            func(d PaymentDetails) bool { return d.CustomerMobileNumber != "" },
		"Amount (must be greater than 0)": func(d PaymentDetails) bool { return d.Amount > 0 },
		"PaymentMethodProperties":         func(d PaymentDetails) bool { return d.PaymentMethodProperties != nil },
	}

	for fieldName, validation := range fieldValidations {
		if !validation(details) {
			missingFields = append(missingFields, fieldName)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("Missing or invalid mandatory fields: %s", strings.Join(missingFields, ", "))
	}

	return nil

}

func Process(details PaymentDetails) (string, error) {
	// Add validation
	if err := ValidatePaymentDetails(details); err != nil {
		return "", err
	}
	return "dsdsdsdd", nil
}
