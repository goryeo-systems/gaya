package exchangeclient

import "math/big"

type Symbol string

const (
	ErrSymbol    Symbol = ""
	BtcUsdt      Symbol = "BTC_USDT"
	BtcPerpetual Symbol = "BTC_PERPETUAL"
)

type TickerEvent struct {
	Symbol       Symbol
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
	TickerStream(symbol Symbol, handler TickerStreamHandler, errHandler ErrHandler) error
	GetWallet() (*Wallet, error)
}
