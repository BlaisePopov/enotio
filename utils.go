package enotio

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// VerifyWebhookSignature проверяет сигнатуру Webhook'а.
//
// Входные параметры:
//   - body: тело Webhook'а в формате JSON.
//   - signature: сигнатура из заголовка x-api-sha256-signature.
//   - secretKey: секретный ключ для проверки подписи.
//
// Возвращаемое значение: true, если сигнатура верна, иначе false.
func VerifyWebhookSignature(body []byte, signature, secretKey string) bool {
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write(body)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}