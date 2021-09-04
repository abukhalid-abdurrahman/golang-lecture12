package main

import (
	"fmt"
	"pkg/types"
)

func FilterByCategory(payments []types.Payment, paymentType types.PaymentType) []types.Payment {
	var filtered = []types.Payment
	for _, payment := range payments {
		if payment.Type == paymentType {
			filtered = append(filtered, payment)
		}
	}

	return filtered
}

func main ()  {
	categories := map[string]float64 {
		"Faridun": 18,
		"Kamila": 9
	}
	categories["Faridun"] = 19

	for key, value := range map {
		// Some code
	}

	for key := range map {
		// Some code
	}

	delete(categories, "Faridun")

	return
}