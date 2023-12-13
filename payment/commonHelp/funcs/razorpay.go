package funcs

import (
	"github.com/razorpay/razorpay-go"
)

func CreateAccount(userid string, email, name string) (map[string]interface{}, error) {
	data := map[string]interface{}{
		"name":          name,
		"email":         email,
		"fail_existing": 0,
		"notes": map[string]interface{}{
			"notes_key_1": "Tea, Earl Grey, Hot",
			"notes_key_2": "Tea, Earl Grey… decaf.",
		},
	}

	client := razorpay.NewClient("rzp_test_siCMLqIerLB4yZ", "w3W7MyJWfOWjW4LPVDMa2nSr")
	costemer, costemerErr := client.Customer.Create(data, nil)
	if costemerErr != nil {
		return nil, costemerErr
	}

	return costemer, nil
}
