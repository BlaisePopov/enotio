package enotio

import (
	"testing"
)

func TestClient_Connect(t *testing.T) {
	//proxyURL, err := url.Parse("http://127.0.0.1:8888")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//client := NewClient(Config{
	//	APIKey: "test",
	//	HTTPClient: &http.Client{
	//		Transport: &http.Transport{
	//			Proxy: http.ProxyURL(proxyURL),
	//			TLSClientConfig: &tls.Config{
	//				InsecureSkipVerify: true,
	//			},
	//		},
	//	},
	//})

	client := NewClient(Config{
		APIKey: "test",
	})

	_, err := client.GetPaymentTariffs(GetPaymentTariffsRequest{ShopID: "44a90940-8206-4438-a71e-5702a9c11f5f"})

	if err == nil || err.Error() != "request failed with status: 401 Unauthorized" {
		t.Fatalf("expected 401 Unauthorized error, got %v", err)
	}
}
