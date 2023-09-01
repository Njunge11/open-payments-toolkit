// This program is used to test the payment package locally.
package main

import (
	"fmt"

	"github.com/Njunge11/open-payments-toolkit/go-sdk/payment"
)

func main() {
	details := payment.PaymentDetails{
		TransactionID:        "123",
		PaymentMethod:        "Mpesa",
		CountryCode:          "KE",
		Amount:               1000,
		CustomerMobileNumber: "+254726325093",
		PaymentMethodProperties: map[string]interface{}{
			"key": "value",
		},
	}

	result := payment.Process(details)
	fmt.Println(result)
}
