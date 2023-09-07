// This program is used to test the payment package locally.
package main

import (
	"fmt"

	"github.com/Njunge11/open-payments-toolkit/go-sdk/payment"
)

func main() {
	details := payment.PaymentDetails{
		TransactionID:        "",
		PaymentMethod:        "",
		CountryCode:          "KE",
		Amount:               0,
		CustomerMobileNumber: "+254726325093",
		PaymentMethodProperties: map[string]interface{}{
			"key": "value",
		},
	}

	result, err := payment.Process(details)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(result)
}
