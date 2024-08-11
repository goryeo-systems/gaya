package main

import (
	"context"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/goryeo-systems/gaya/pkg/binanceapi"
	"github.com/goryeo-systems/gaya/pkg/util"
)

var log = util.GetLogger()

func tickerEventHandler(event *binanceapi.TickerEvent) {
	log.Info("event", "event", event)
}

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

	err = binanceapi.TickerStream(&binanceapi.CurrencyPair{Base: "BTC", Quote: "USDT"}, tickerEventHandler, util.LogError)
	if err != nil {
		util.Check(err)
	}
	// remove this if you do not want to be blocked here
	time.Sleep(5 * time.Second)
	//<-doneChan

	println("Hello, Gaya!")
}
