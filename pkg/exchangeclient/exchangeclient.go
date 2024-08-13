package exchangeclient

import "math/big"

type CurrencyPair struct {
	Base  string
	Quote string
}

type TickerEvent struct {
	CurrencyPair *CurrencyPair
	BestBidPrice *big.Float
	BestBidQty   *big.Float
	BestAskPrice *big.Float
	BestAskQty   *big.Float
}

type Wallet struct {
	Available map[string]*big.Float
}

type TickerStreamHandler func(event *TickerEvent)
type ErrHandler func(err error)

type ExchangeClient interface {
	TickerStream(currencyPair *CurrencyPair, handler TickerStreamHandler, errHandler ErrHandler) error
	GetWallet() (*Wallet, error)
}
