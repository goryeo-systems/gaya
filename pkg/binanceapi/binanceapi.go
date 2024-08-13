package binanceapi

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/adshao/go-binance/v2"
	"github.com/goryeo-systems/gaya/pkg/exchangeclient"
	"github.com/goryeo-systems/gaya/pkg/util"
)

func binanceSymbolToCurrencyPair(symbol string) (*exchangeclient.CurrencyPair, error) {
	switch symbol {
	case "BTCUSDT":
		return &exchangeclient.CurrencyPair{
			Base:  "BTC",
			Quote: "USDT",
		}, nil
	default:
		return nil, fmt.Errorf("unsupported symbol: %s", symbol)
	}
}

func currencyPairToBinanceSymbol(currencyPair *exchangeclient.CurrencyPair) (string, error) {
	if currencyPair.Base == "BTC" && currencyPair.Quote == "USDT" {
		return "BTCUSDT", nil
	}

	return "", fmt.Errorf("unsupported currency pair: %v", currencyPair)
}

func toTickerEvent(event *binance.WsBookTickerEvent) (*exchangeclient.TickerEvent, error) {
	currencyPair, err := binanceSymbolToCurrencyPair(event.Symbol)
	if err != nil {
		return nil, err
	}

	bestBidPrice, err := util.StringToBigFloat(event.BestBidPrice)
	if err != nil {
		return nil, err
	}

	bestBidQty, err := util.StringToBigFloat(event.BestBidQty)
	if err != nil {
		return nil, err
	}

	bestAskPrice, err := util.StringToBigFloat(event.BestAskPrice)
	if err != nil {
		return nil, err
	}

	bestAskQty, err := util.StringToBigFloat(event.BestAskQty)
	if err != nil {
		return nil, err
	}

	return &exchangeclient.TickerEvent{
		CurrencyPair: currencyPair,
		BestBidPrice: bestBidPrice,
		BestBidQty:   bestBidQty,
		BestAskPrice: bestAskPrice,
		BestAskQty:   bestAskQty,
	}, nil
}

// TickerStream subscribes to the ticker stream for the given currency pair.
func (c *BinanceClient) TickerStream(
	currencyPair *exchangeclient.CurrencyPair,
	handler exchangeclient.TickerStreamHandler,
	errHandler exchangeclient.ErrHandler,
) error {
	symbol, err := currencyPairToBinanceSymbol(currencyPair)
	if err != nil {
		return err
	}

	_, _, err = binance.WsBookTickerServe(
		symbol,
		func(event *binance.WsBookTickerEvent) {
			tickerEvent, err := toTickerEvent(event)
			if err != nil {
				errHandler(err)

				return
			}

			handler(tickerEvent)
		},
		func(err error) {
			errHandler(err)
		},
	)

	return err
}

type BinanceClient struct {
	client *binance.Client
}

// New creates a new Binance client.
func New() *BinanceClient {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")

	return &BinanceClient{
		client: binance.NewClient(apiKey, secretKey),
	}
}

var zeroBigFloat = big.NewFloat(0) //nolint:gochecknoglobals

// GetWallet returns the wallet of the user.
func (c *BinanceClient) GetWallet() (*exchangeclient.Wallet, error) {
	account, err := c.client.NewGetAccountService().Do(context.Background())
	if err != nil {
		return nil, err
	}

	wallet := &exchangeclient.Wallet{
		Available: make(map[string]*big.Float),
	}

	for _, balance := range account.Balances {
		available, err := util.StringToBigFloat(balance.Free)
		if err != nil {
			return nil, err
		}

		if available.Cmp(zeroBigFloat) == 0 {
			continue
		}

		wallet.Available[balance.Asset] = available
	}

	return wallet, nil
}
