package enotio

type GetPaymentTariffsRequest struct {
	ShopID string `json:"shop_id"`
}

type CreateInvoiceRequest struct {
	Amount         float64         `json:"amount"`
	OrderID        string          `json:"order_id"`
	Currency       Currency        `json:"currency"`
	ShopID         string          `json:"shop_id"`
	HookURL        string          `json:"hook_url,omitempty"`
	CustomFields   string          `json:"custom_fields,omitempty"`
	Comment        string          `json:"comment,omitempty"`
	FailURL        string          `json:"fail_url,omitempty"`
	SuccessURL     string          `json:"success_url,omitempty"`
	Expire         int             `json:"expire,omitempty"`
	IncludeService []PaymentMethod `json:"include_service,omitempty"`
	ExcludeService []PaymentMethod `json:"exclude_service,omitempty"`
}

type CreateH2HRequest struct {
	InvoiceID   string      `json:"invoice_id"`
	ShopID      string      `json:"shop_id"`
	IP          string      `json:"ip"`
	UserAgent   string      `json:"user_agent"`
	PaymentData PaymentData `json:"payment_data"`
}

type PaymentData struct {
	Bank string `json:"bank"`
}

type GetInvoiceInfoRequest struct {
	ShopID    string `json:"shop_id"`
	InvoiceID string `json:"invoice_id,omitempty"`
	OrderID   string `json:"order_id,omitempty"`
}

type GetPayoffInfoRequest struct {
	UserID   string `json:"user_id"`
	PayoffID string `json:"id,omitempty"`
	OrderID  string `json:"order_id,omitempty"`
}

type CreatePayoffRequest struct {
	UserID     string     `json:"user_id"`
	Service    string     `json:"service"`
	WalletTo   string     `json:"wallet_to"`
	Amount     float64    `json:"amount"`
	OrderID    string     `json:"order_id,omitempty"`
	Comment    string     `json:"comment,omitempty"`
	HookURL    string     `json:"hook_url,omitempty"`
	Subtract   int        `json:"subtract,omitempty"`
	PayoffData PayoffData `json:"payoff_data,omitempty"`
}

type PayoffData struct {
	SbpBankType string `json:"sbp_bank_type,omitempty"`
}

type GetSbpBankListRequest struct {
	UserID string `json:"user_id"`
}
