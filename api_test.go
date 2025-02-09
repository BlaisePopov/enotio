package enotio

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPaymentTariffs(t *testing.T) {
	// Создаем тестовый сервер
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"tariffs":[{"percent":0,"max_sum":0,"shop_percent":0,"user_percent":0,"service":"card","service_label":"Банковская карта","currency":"RUB","status":"string"}]},"status":200,"status_check":true}`))
	}))
	defer ts.Close()

	// Создаем клиента с тестовым URL
	client := NewClient(Config{APIKey: "test_api_key", BaseURL: ts.URL})

	// Вызываем метод
	response, err := client.GetPaymentTariffs(GetPaymentTariffsRequest{ShopID: "shop_id"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Проверяем результат
	if len(response.Tariffs) == 0 {
		t.Fatalf("Expected at least one tariff, got 0")
	}
	if response.Tariffs[0].Service != "card" {
		t.Fatalf("Expected service 'card', got %v", response.Tariffs[0].Service)
	}
}

func TestCreateInvoice(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"id":"uuid","amount":100,"currency":"RUB","url":"https://enot.io/pay/uuid","expired":"2023-12-31 23:59:59"},"status":200,"status_check":true}`))
	}))
	defer ts.Close()

	client := NewClient(Config{APIKey: "test_api_key", BaseURL: ts.URL})

	response, err := client.CreateInvoice(CreateInvoiceRequest{
		Amount:   100,
		OrderID:  "order123",
		Currency: "RUB",
		ShopID:   "shop_id",
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.ID != "uuid" {
		t.Fatalf("Expected ID 'uuid', got %v", response.ID)
	}
}

func TestCreateH2H(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"step_two_type":"p2p_data","wallet":"2222222222222222","receiver_name":"Иванов Иван Иванович","bank":"Сбербанк","amount":"1005","time_end":"2024-03-15 19:09:06","url":null,"type":"p2p_data","is_redirect":false,"fingerprint":false,"is_need_get_info":false},"status":200,"status_check":true}`))
	}))
	defer ts.Close()

	client := NewClient(Config{APIKey: "test_api_key", BaseURL: ts.URL})

	response, err := client.CreateH2H(CreateH2HRequest{
		InvoiceID:   "invoice_id",
		ShopID:      "shop_id",
		IP:          "127.0.0.1",
		UserAgent:   "Mozilla/5.0",
		PaymentData: PaymentData{Bank: "sber"},
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Wallet != "2222222222222222" {
		t.Fatalf("Expected wallet '2222222222222222', got %v", response.Wallet)
	}
}

func TestGetInvoiceInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"invoice_id":"uuid","order_id":"order123","shop_id":"shop_id","status":"created","invoice_amount":100,"credited":95,"currency":"RUB","pay_service":"card","payer_details":"553691******1279","commission_amount":5,"commission_percent":5,"shop_commission_amount":3,"user_commission_amount":2,"user_commission_percent":2,"custom_field":"{}","created_at":"2023-01-01 00:00:00","expired_at":"2023-01-02 00:00:00","paid_at":"2023-01-01 12:00:00"},"status":200,"status_check":true}`))
	}))
	defer ts.Close()

	client := NewClient(Config{APIKey: "test_api_key", BaseURL: ts.URL})

	response, err := client.GetInvoiceInfo(GetInvoiceInfoRequest{
		ShopID:    "shop_id",
		InvoiceID: "uuid",
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.InvoiceID != "uuid" {
		t.Fatalf("Expected invoice ID 'uuid', got %v", response.InvoiceID)
	}
}

func TestGetBalance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"balance":1000,"active_balance":800,"freeze_balance":200},"status":200,"status_check":true}`))
	}))
	defer ts.Close()

	client := NewClient(Config{APIKey: "test_api_key", BaseURL: ts.URL})

	response, err := client.GetBalance("user_id")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.Balance != 1000 {
		t.Fatalf("Expected balance 1000, got %v", response.Balance)
	}
}

func TestCreatePayoff(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"payoff_id":"uuid","amount_withdraw_rub":100,"balance":900},"status":200,"status_check":true}`))
	}))
	defer ts.Close()

	client := NewClient(Config{APIKey: "test_api_key", BaseURL: ts.URL})

	response, err := client.CreatePayoff(CreatePayoffRequest{
		UserID:   "user_id",
		Service:  "card",
		WalletTo: "400000000000000",
		Amount:   100,
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.PayoffID != "uuid" {
		t.Fatalf("Expected payoff ID 'uuid', got %v", response.PayoffID)
	}
}

func TestGetPayoffInfo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"payoff_id":"uuid","status":"success","service":"card","wallet":"400000000000000","type":"api","subtract":"balance","amount":100,"amount_withdraw_rub":100,"commission_rub":5,"receive_currency":"RUB","amount_receive":"100","comment":"Test","created_at":"2023-01-01 00:00:00","paid_at":"2023-01-01 12:00:00"},"status":200,"status_check":true}`))
	}))
	defer ts.Close()

	client := NewClient(Config{APIKey: "test_api_key", BaseURL: ts.URL})

	response, err := client.GetPayoffInfo(GetPayoffInfoRequest{
		UserID:   "user_id",
		PayoffID: "uuid",
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if response.PayoffID != "uuid" {
		t.Fatalf("Expected payoff ID 'uuid', got %v", response.PayoffID)
	}
}

func TestGetSbpBankList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"banks":[{"id":"100000000001","name":"Газпромбанк"},{"id":"100000000003","name":"Банк Синара"}]},"status":200,"status_check":true}`))
	}))
	defer ts.Close()

	client := NewClient(Config{APIKey: "test_api_key", BaseURL: ts.URL})

	response, err := client.GetSbpBankList(GetSbpBankListRequest{UserID: "user_id"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(response.Banks) == 0 {
		t.Fatalf("Expected at least one bank, got 0")
	}
}
