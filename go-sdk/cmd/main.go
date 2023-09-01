// This program is used to test the payment package locally.
package main

import (
	"fmt"

	"github.com/Njunge11/open-payments-toolkit/go-sdk/payment"
)

func main() {
	fmt.Println(payment.ProcessPayment())
}
