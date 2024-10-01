package main

import (
	"os"
	"time"

	"github.com/frankrap/deribit-api"
	"github.com/frankrap/deribit-api/models"
	"github.com/goryeo-systems/gaya/pkg/binanceapi"
	"github.com/goryeo-systems/gaya/pkg/exchangeclient"
	"github.com/goryeo-systems/gaya/pkg/util"
)

func tickerEventHandler(event *exchangeclient.TickerEvent) {
	util.Log.Info("BINANCE", "event", event)
}

func main() {
	cfg := &deribit.Configuration{
		Addr:          deribit.RealBaseURL,
		ApiKey:        os.Getenv("DERIBIT_API_KEY"),
		SecretKey:     os.Getenv("DERIBIT_SECRET_KEY"),
		AutoReconnect: true,
		DebugMode:     true,
	}
	client := deribit.New(cfg)

	client.On("ticker.BTC-PERPETUAL.raw", func(e *models.TickerNotification) {
		util.Log.Info("DERIBIT", "event", e)
	})

	client.Subscribe([]string{
		"ticker.BTC-PERPETUAL.raw",
	})

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
