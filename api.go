package enotio

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// GetPaymentTariffs получает список доступных методов оплаты. [https://docs.enot.io/e/new/payments-methods-new#request]
//
// Входные параметры:
//   - request: структура GetPaymentTariffsRequest.
//
// Возвращаемое значение: структура PaymentTariffsResponse и ошибка, если она возникла.
func (c *Client) GetPaymentTariffs(request GetPaymentTariffsRequest) (*PaymentTariffsResponse, error) {
	paymentTariffsResponse := &PaymentTariffsResponse{}

	body, err := c.sendRequest("GET", "/shops/"+request.ShopID+"/payment-tariffs", nil)
	if err != nil {
		return nil, err
	}

	baseResponse := &BaseResponse{Data: paymentTariffsResponse}
	err = json.Unmarshal(body, baseResponse)
	if err != nil {
		return nil, err
	}

	if !baseResponse.StatusCheck {
		return nil, fmt.Errorf("status check: %s", baseResponse.Error)
	}

	return paymentTariffsResponse, nil
}

// CreateInvoice создает новый инвойс. [https://docs.enot.io/e/new/create-invoice#request]
//
// Входные параметры:
//   - request: структура CreateInvoiceRequest.
//
// Возвращаемое значение: структура InvoiceResponse и ошибка, если она возникла.
func (c *Client) CreateInvoice(request CreateInvoiceRequest) (*InvoiceResponse, error) {
	invoiceResponse := &InvoiceResponse{}

	params, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest("POST", "/invoice/create", params)
	if err != nil {
		return nil, err
	}

	baseResponse := &BaseResponse{Data: invoiceResponse}
	err = json.Unmarshal(body, baseResponse)
	if err != nil {
		return nil, err
	}

	if !baseResponse.StatusCheck {
		return nil, fmt.Errorf("status check: %s", baseResponse.Error)
	}

	return invoiceResponse, nil
}

// CreateH2H создает платеж h2h. [https://docs.enot.io/e/new/create-invoice-h2h#info]
//
// Входные параметры:
//   - request: структура CreateH2HRequest.
//
// Возвращаемое значение: структура H2HResponse и ошибка, если она возникла.
func (c *Client) CreateH2H(request CreateH2HRequest) (*H2HResponse, error) {
	h2hResponse := &H2HResponse{}

	params, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest("POST", "/invoice/h2h", params)
	if err != nil {
		return nil, err
	}

	baseResponse := &BaseResponse{Data: h2hResponse}
	err = json.Unmarshal(body, baseResponse)
	if err != nil {
		return nil, err
	}

	if !baseResponse.StatusCheck {
		return nil, fmt.Errorf("status check: %s", baseResponse.Error)
	}

	return h2hResponse, nil
}

// GetInvoiceInfo получает информацию о платеже. [https://docs.enot.io/e/new/invoice-info-new#request]
//
// Входные параметры:
//   - request: структура GetInvoiceInfoRequest.
//
// Возвращаемое значение: структура InvoiceInfoResponse и ошибка, если она возникла.
func (c *Client) GetInvoiceInfo(request GetInvoiceInfoRequest) (*InvoiceInfoResponse, error) {
	invoiceInfoResponse := &InvoiceInfoResponse{}

	params := url.Values{}
	params.Add("shop_id", request.ShopID)
	if request.InvoiceID != "" {
		params.Add("invoice_id", request.InvoiceID)
	}
	if request.OrderID != "" {
		params.Add("order_id", request.OrderID)
	}

	body, err := c.sendRequest("GET", "/invoice/info?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	baseResponse := &BaseResponse{Data: invoiceInfoResponse}
	err = json.Unmarshal(body, baseResponse)
	if err != nil {
		return nil, err
	}

	if !baseResponse.StatusCheck {
		return nil, fmt.Errorf("status check: %s", baseResponse.Error)
	}

	return invoiceInfoResponse, nil
}

// GetBalance получает баланс пользователя. [https://docs.enot.io/e/new/new-get-balance#request]
//
// Входные параметры:
//   - userID: идентификатор пользователя.
//
// Возвращаемое значение: структура BalanceResponse и ошибка, если она возникла.
func (c *Client) GetBalance(userID string) (*BalanceResponse, error) {
	balanceResponse := &BalanceResponse{}

	body, err := c.sendRequest("GET", "/account/users/"+userID+"/balance", nil)
	if err != nil {
		return nil, err
	}

	baseResponse := &BaseResponse{Data: balanceResponse}
	err = json.Unmarshal(body, baseResponse)
	if err != nil {
		return nil, err
	}

	if !baseResponse.StatusCheck {
		return nil, fmt.Errorf("status check: %s", baseResponse.Error)
	}

	return balanceResponse, nil
}

// CreatePayoff создает заявку на вывод средств. [https://docs.enot.io/e/new/new-create-withdraw#payoff-request]
//
// Входные параметры:
//   - request: структура CreatePayoffRequest.
//
// Возвращаемое значение: структура PayoffResponse и ошибка, если она возникла.
func (c *Client) CreatePayoff(request CreatePayoffRequest) (*PayoffResponse, error) {
	payoffResponse := &PayoffResponse{}

	params, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	body, err := c.sendRequest("POST", "/payoff/create", params)
	if err != nil {
		return nil, err
	}

	baseResponse := &BaseResponse{Data: payoffResponse}
	err = json.Unmarshal(body, baseResponse)
	if err != nil {
		return nil, err
	}

	if !baseResponse.StatusCheck {
		return nil, fmt.Errorf("status check: %s", baseResponse.Error)
	}

	return payoffResponse, nil
}

// GetPayoffInfo получает информацию о выводе. [https://docs.enot.io/e/new/new-payment-info#get-payoff-info]
//
// Входные параметры:
//   - request: структура GetPayoffInfoRequest.
//
// Возвращаемое значение: структура PayoffInfoResponse и ошибка, если она возникла.
func (c *Client) GetPayoffInfo(request GetPayoffInfoRequest) (*PayoffInfoResponse, error) {
	payoffInfoResponse := &PayoffInfoResponse{}

	params := url.Values{}
	params.Add("user_id", request.UserID)
	if request.PayoffID != "" {
		params.Add("id", request.PayoffID)
	}
	if request.OrderID != "" {
		params.Add("order_id", request.OrderID)
	}

	body, err := c.sendRequest("GET", "/payoff/info?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	baseResponse := &BaseResponse{Data: payoffInfoResponse}
	err = json.Unmarshal(body, baseResponse)
	if err != nil {
		return nil, err
	}

	if !baseResponse.StatusCheck {
		return nil, fmt.Errorf("status check: %s", baseResponse.Error)
	}

	return payoffInfoResponse, nil
}

// GetSbpBankList получает список банков для СБП. [https://docs.enot.io/e/new/get-sbp-bank-list#get-list-bank]
//
// Входные параметры:
//   - request: структура GetSbpBankListRequest.
//
// Возвращаемое значение: структура SbpBankListResponse и ошибка, если она возникла.
func (c *Client) GetSbpBankList(request GetSbpBankListRequest) (*SbpBankListResponse, error) {
	sbpBankListResponse := &SbpBankListResponse{}

	params := url.Values{}
	params.Add("user_id", request.UserID)

	body, err := c.sendRequest("GET", "/payoff/get-sbp-bank-list?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	baseResponse := &BaseResponse{Data: sbpBankListResponse}
	err = json.Unmarshal(body, baseResponse)
	if err != nil {
		return nil, err
	}

	if !baseResponse.StatusCheck {
		return nil, fmt.Errorf("status check: %s", baseResponse.Error)
	}

	return sbpBankListResponse, nil
}
