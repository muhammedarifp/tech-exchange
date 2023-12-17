package funcs

import (
	"fmt"

	"github.com/muhammedarifp/tech-exchange/payments/config"

	"github.com/razorpay/razorpay-go"
)

func CreateAccount(email, name string) (map[string]interface{}, error) {
	cfg := config.GetConfig()
	fmt.Println("--------> ", email+" ------- ", name)
	data := map[string]interface{}{
		"name":          name,
		"email":         email,
		"fail_existing": 0,
	}

	client := razorpay.NewClient(cfg.RAZORPAY_KEY, cfg.RAZORPAY_SEC)
	costemer, costemerErr := client.Customer.Create(data, nil)
	if costemerErr != nil {
		if costemerErr.Error() == "Customer already exists for the merchant" {

		}
		return nil, costemerErr
	}

	return costemer, nil
}
