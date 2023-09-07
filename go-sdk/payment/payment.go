package payment

import (
	"github.com/Njunge11/open-payments-toolkit/go-sdk/validation"
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
	ProcessPayment(details PaymentDetails) (string, error)
	Validate(details PaymentDetails) error
}

func CheckForMissingPaymentDetails(details PaymentDetails) error {
	var missingPaymentDetails []validation.ValidationError
	paymentDetailsValidations := map[string]func(PaymentDetails) bool{
		"TransactionID":           func(d PaymentDetails) bool { return d.TransactionID != "" },
		"PaymentMethod":           func(d PaymentDetails) bool { return d.PaymentMethod != "" },
		"CountryCode":             func(d PaymentDetails) bool { return d.CountryCode != "" },
		"Amount":                  func(d PaymentDetails) bool { return d.Amount > 0 },
		"CustomerMobileNumber":    func(d PaymentDetails) bool { return d.CustomerMobileNumber != "" },
		"PaymentMethodProperties": func(d PaymentDetails) bool { return d.PaymentMethodProperties != nil },
	}

	for paymentDetailProperty, validateFunc := range paymentDetailsValidations {
		if !validateFunc(details) {
			missingPaymentDetails = append(missingPaymentDetails, validation.ValidationError{
				Name:   paymentDetailProperty,
				Reason: "is empty",
			})
		}
	}

	if len(missingPaymentDetails) > 0 {
		return &validation.ValidationErrors{
			Type:          "https://example.net/validation-error",
			Title:         "Your request parameters didn't validate.",
			InvalidParams: missingPaymentDetails,
		}
	}
	return nil
}

func Process(details PaymentDetails) (string, error) {
	if err := CheckForMissingPaymentDetails(details); err != nil {
		return "", err
	}

	// Perform actual payment processing here
	// For demonstration, returning a success message
	return "Payment processed successfully.", nil
}
