package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/goryeo-systems/gaya/pkg/util"
)

func StringToBigFloat(s string) (*big.Float, error) {
	if v, ok := new(big.Float).SetString(s); ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("failed to convert string to big.Float: %s", s)
	}
}

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

func binanceSymbolToCurrencyPair(symbol string) (*CurrencyPair, error) {
	switch symbol {
	case "BTCUSDT":
		return &CurrencyPair{
			Base:  "BTC",
			Quote: "USDT",
		}, nil
	default:
		return nil, fmt.Errorf("unsupported symbol: %s", symbol)
	}
}

func toTickerEvent(event *binance.WsBookTickerEvent) (*TickerEvent, error) {
	currencyPair, err := binanceSymbolToCurrencyPair(event.Symbol)
	if err != nil {
		return nil, err
	}

	bestBidPrice, err := StringToBigFloat(event.BestBidPrice)
	if err != nil {
		return nil, err
	}

	bestBidQty, err := StringToBigFloat(event.BestBidQty)
	if err != nil {
		return nil, err
	}

	bestAskPrice, err := StringToBigFloat(event.BestAskPrice)
	if err != nil {
		return nil, err
	}

	bestAskQty, err := StringToBigFloat(event.BestAskQty)
	if err != nil {
		return nil, err
	}

	return &TickerEvent{
		CurrencyPair: currencyPair,
		BestBidPrice: bestBidPrice,
		BestBidQty:   bestBidQty,
		BestAskPrice: bestAskPrice,
		BestAskQty:   bestAskQty,
	}, nil
}

var log = util.GetLogger()

func main() {
	client := binance.NewClient("", "")

	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, p := range prices {
		fmt.Println(p)
	}

	handler := func(event *TickerEvent) {
		log.Info("event", "event", event)
	}

	wsBookTickerHandler := func(event *binance.WsBookTickerEvent) {
		fmt.Println(event)
		e, err := toTickerEvent(event)
		if err != nil {
			fmt.Println(err)
			return
		}
		handler(e)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	_, _, err = binance.WsBookTickerServe("BTCUSDT", wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// remove this if you do not want to be blocked here
	time.Sleep(5 * time.Second)
	//<-doneChan

	println("Hello, Gaya!")
}
