package main

import (
	"context"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/goryeo-systems/gaya/pkg/binanceapi"
	"github.com/goryeo-systems/gaya/pkg/util"
)

func tickerEventHandler(event *binanceapi.TickerEvent) {
	util.Log.Info("event", "event", event)
}

func main() {
	client := binance.NewClient("", "")

	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		util.LogError(err)

		return
	}

	c := binanceapi.New()
	w, err := c.GetWallet()
	if err != nil {
		util.LogError(err)
		return
	}
	util.Log.Info("wallet", "wallet", w)

	for _, p := range prices {
		util.Log.Info("price", "price", p)
	}

	err = binanceapi.TickerStream(&binanceapi.CurrencyPair{Base: "BTC", Quote: "USDT"}, tickerEventHandler, util.LogError)
	if err != nil {
		util.Check(err)
	}
	// remove this if you do not want to be blocked here
	time.Sleep(5 * time.Second) //nolint:all
}
