package enotio

type PaymentMethod string

const (
	PaymentMethodCard         PaymentMethod = "card"
	PaymentMethodMirCard      PaymentMethod = "mir_card"
	PaymentMethodQiwi         PaymentMethod = "qiwi"
	PaymentMethodPerfectMoney PaymentMethod = "perfect_money"
	PaymentMethodYoomoney     PaymentMethod = "yoomoney"
	PaymentMethodSbp          PaymentMethod = "sbp"
	PaymentMethodZcash        PaymentMethod = "zcash"
	PaymentMethodAdvcash      PaymentMethod = "advcash"
	PaymentMethodWebmoney     PaymentMethod = "webmoney"
	PaymentMethodGooglePay    PaymentMethod = "google_pay"
	PaymentMethodApplePay     PaymentMethod = "apple_pay"
	PaymentMethodBitcoin      PaymentMethod = "bitcoin"
	PaymentMethodEthereum     PaymentMethod = "ethereum"
	PaymentMethodDash         PaymentMethod = "dash"
	PaymentMethodLitecoin     PaymentMethod = "litecoin"
	PaymentMethodUsdtTrc20    PaymentMethod = "usdt_trc20"
	PaymentMethodUsdtErc20    PaymentMethod = "usdt_erc20"
	PaymentMethodTrx          PaymentMethod = "trx"
	PaymentMethodTon          PaymentMethod = "ton"
	PaymentMethodBitcoinCash  PaymentMethod = "bitcoin_cash"
	PaymentMethodP2PCard      PaymentMethod = "p2p_card"
	PaymentMethodP2PSbp       PaymentMethod = "p2p_sbp"
)

type Currency string

const (
	CurrencyRUB        Currency = "RUB"
	CurrencyUSD        Currency = "USD"
	CurrencyEUR        Currency = "EUR"
	CurrencyUAH        Currency = "UAH"
	CurrencyKZT        Currency = "KZT"
	CurrencyBTC        Currency = "BTC"
	CurrencyLTC        Currency = "LTC"
	CurrencyUSDT_TRC20 Currency = "USDT_TRC20"
	CurrencyUSDT_ERC20 Currency = "USDT_ERC20"
	CurrencyTRX        Currency = "TRX"
	CurrencyTON        Currency = "TON"
	CurrencyDASH       Currency = "DASH"
	CurrencyETH        Currency = "ETH"
	CurrencyZCASH      Currency = "ZCASH"
	CurrencyBTC_CASH   Currency = "BTC_CASH"
)
