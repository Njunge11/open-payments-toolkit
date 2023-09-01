// Package payment provides functionality for processing payments.
package payment

// ProcessPayment returns a processed payment message.
/**
- transactionId
- paymentMethod
- countryCode
- amount
- customerMobileNumber
- paymentMethodProperties - each payment method may have unique properites
**/

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

func Process(details PaymentDetails) PaymentDetails {
	return details
}
