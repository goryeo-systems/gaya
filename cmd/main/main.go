package main

import (
	"time"

	"github.com/goryeo-systems/gaya/pkg/binanceapi"
	"github.com/goryeo-systems/gaya/pkg/exchangeclient"
	"github.com/goryeo-systems/gaya/pkg/util"
)

func tickerEventHandler(event *exchangeclient.TickerEvent) {
	util.Log.Info("event", "event", event)
}

func main() {
	c := binanceapi.New()

	w, err := c.GetWallet()
	if err != nil {
		util.LogError(err)

		return
	}

	util.Log.Info("wallet", "wallet", w)

	err = c.TickerStream(&exchangeclient.CurrencyPair{Base: "BTC", Quote: "USDT"}, tickerEventHandler, util.LogError)
	if err != nil {
		util.Check(err)
	}
	// remove this if you do not want to be blocked here
	time.Sleep(5 * time.Second) //nolint:all
}
