# Go Client для API Enot.io

[![Go Reference](https://pkg.go.dev/badge/github.com/BlaisePopov/cloudpayments.svg)](https://pkg.go.dev/github.com/BlaisePopov/enotio)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/BlaisePopov/enotio/blob/main/LICENSE)

Go-клиент для взаимодействия с API платежной системы [Enot.io](https://enot.io). Клиент поддерживает все методы API, включая создание платежей, вывод средств, проверку баланса и другие.

## Установка

Для установки пакета выполните:

```bash
go get github.com/BlaisePopov/enotio
```

## Примеры использования

### Инициализация клиента

```go
package main

import (
    "fmt"
    "github.com/BlaisePopov/enotio"
)

func main() {
    client := enotio.NewClient(Config{
      APIKey: "test",
    })
}
```

### Создание платежа

```go
invoice, err := client.CreateInvoice(enot.CreateInvoiceRequest{
    Amount:   100.0,
    OrderID:  "order_123",
    Currency: enot.CurrencyRUB,
    ShopID:   "your_shop_id",
    Expire:   360, // 6 часов
})
if err != nil {
    panic(err)
}
fmt.Printf("Ссылка для оплаты: %s\n", invoice.URL)
```

### Получение баланса

```go
balance, err := client.GetBalance("your_user_id")
if err != nil {
    panic(err)
}
fmt.Printf("Баланс: %.2f RUB\n", balance.Balance)
```

### Создание выплаты

```go
payoff, err := client.CreatePayoff(enot.CreatePayoffRequest{
    UserID:   "your_user_id",
    Service:  enotio.PaymentMethodCard,
    WalletTo: "4000000000000000",
    Amount:   500.0,
})
if err != nil {
    panic(err)
}
fmt.Printf("ID выплаты: %s\n", payoff.PayoffID)
```

## Документация API
Официальная документация Enot.io:  
[https://docs.enot.io](https://docs.enot.io)

## Лицензия
Этот проект распространяется под лицензией MIT. Подробности см. в файле [LICENSE](LICENSE).