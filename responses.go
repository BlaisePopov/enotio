package enotio

type BaseResponse struct {
	Data        interface{} `json:"data"`
	Status      int         `json:"status"`
	StatusCheck bool        `json:"status_check"`
	Error       string      `json:"error,omitempty"`
}

type PaymentTariffsResponse struct {
	Tariffs []Tariff `json:"tariffs"`
}

type Tariff struct {
	Percent      float64 `json:"percent"`
	Fix          float64 `json:"fix"`
	MinSum       float64 `json:"min_sum"`
	MaxSum       float64 `json:"max_sum"`
	ShopPercent  float64 `json:"shop_percent"`
	UserPercent  float64 `json:"user_percent"`
	Service      string  `json:"service"`
	ServiceLabel string  `json:"service_label"`
	Currency     string  `json:"currency"`
	Status       string  `json:"status"`
}

type InvoiceResponse struct {
	ID       string `json:"id"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	URL      string `json:"url"`
	Expired  string `json:"expired"`
}

type H2HResponse struct {
	StepTwoType   string `json:"step_two_type"`
	Wallet        string `json:"wallet"`
	ReceiverName  string `json:"receiver_name"`
	Bank          string `json:"bank"`
	Amount        string `json:"amount"`
	TimeEnd       string `json:"time_end"`
	URL           string `json:"url"`
	Type          string `json:"type"`
	IsRedirect    bool   `json:"is_redirect"`
	Fingerprint   bool   `json:"fingerprint"`
	IsNeedGetInfo bool   `json:"is_need_get_info"`
}

type InvoiceInfoResponse struct {
	InvoiceID             string   `json:"invoice_id"`
	OrderID               string   `json:"order_id"`
	ShopID                string   `json:"shop_id"`
	Status                string   `json:"status"`
	InvoiceAmount         float64  `json:"invoice_amount"`
	Credited              *float64 `json:"credited"`
	Currency              string   `json:"currency"`
	PayService            *string  `json:"pay_service"`
	PayerDetails          *string  `json:"payer_details"`
	CommissionAmount      float64  `json:"commission_amount"`
	CommissionPercent     float64  `json:"commission_percent"`
	ShopCommissionAmount  float64  `json:"shop_commission_amount"`
	UserCommissionAmount  float64  `json:"user_commission_amount"`
	UserCommissionPercent float64  `json:"user_commission_percent"`
	CustomField           string   `json:"custom_field"`
	CreatedAt             string   `json:"created_at"`
	ExpiredAt             string   `json:"expired_at"`
	PaidAt                string   `json:"paid_at"`
}

type BalanceResponse struct {
	Balance       float64 `json:"balance"`
	ActiveBalance float64 `json:"active_balance"`
	FreezeBalance float64 `json:"freeze_balance"`
}

type PayoffResponse struct {
	PayoffID          string `json:"payoff_id"`
	AmountWithdrawRub string `json:"amount_withdraw_rub"`
	Balance           string `json:"balance"`
}

type PayoffInfoResponse struct {
	PayoffID          string  `json:"payoff_id"`
	Status            string  `json:"status"`
	OrderID           string  `json:"order_id"`
	Service           string  `json:"service"`
	Wallet            *string `json:"wallet"`
	Type              string  `json:"type"`
	Subtract          string  `json:"subtract"`
	Amount            float64 `json:"amount"`
	AmountWithdrawRub float64 `json:"amount_withdraw_rub"`
	CommissionRub     float64 `json:"commission_rub"`
	ReceiveCurrency   string  `json:"receive_currency"`
	AmountReceive     string  `json:"amount_receive"`
	Comment           *string `json:"comment"`
	CreatedAt         string  `json:"created_at"`
	PaidAt            string  `json:"paid_at"`
	ErrorMessage      string  `json:"error_message"`
}

type SbpBankListResponse struct {
	Banks []Bank `json:"banks"`
}

type Bank struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
