package enotio

import "testing"

func TestVerifyWebhookSignature(t *testing.T) {
	body := []byte(`{"amount":"100.00","credited":"95.50","custom_fields":{"user":1},"invoice_id":"a3e9ff6f-c5c1-3bcd-854e-4bc995b1ae7a","order_id":"c78d8fe9-ab44-3f21-a37a-ce4ca269cb47","pay_service":"card","pay_time":"2023-04-06 16:27:59","payer_details":"553691******1279","status":"success","type":1}`)
	signature := "e582b14dd13f8111711e3cb66a982fd7bff28a0ddece8bde14a34a5bb4449136"
	secretKey := "example"

	isValid := VerifyWebhookSignature(body, signature, secretKey)
	if !isValid {
		t.Fatalf("Expected signature to be valid, got invalid")
	}
}
