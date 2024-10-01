package main

import (
	"time"

	"github.com/goryeo-systems/gaya/pkg/binanceapi"
	"github.com/goryeo-systems/gaya/pkg/deribitapi"
	"github.com/goryeo-systems/gaya/pkg/exchangeclient"
	"github.com/goryeo-systems/gaya/pkg/util"
)

func tickerEventHandler(event *exchangeclient.TickerEvent) {
	util.Log.Info("BINANCE", "event", event)
}

func deribitTickerEventHandler(event *exchangeclient.TickerEvent) {
	util.Log.Info("DERIBIT", "event", event)
}

func main() {
	deribitClient := deribitapi.New()
	err := deribitClient.TickerStream(exchangeclient.BtcPerpetual, deribitTickerEventHandler, util.LogError)
	if err != nil {
		util.Check(err)
	}

	c := binanceapi.New()

	w, err := c.GetWallet()
	if err != nil {
		util.LogError(err)

		return
	}

	util.Log.Info("wallet", "wallet", w)

	err = c.TickerStream(exchangeclient.BtcUsdt, tickerEventHandler, util.LogError)
	if err != nil {
		util.Check(err)
	}
	// remove this if you do not want to be blocked here
	time.Sleep(5 * time.Second) //nolint:all
}
