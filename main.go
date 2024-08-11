package main

import (
	"context"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
)

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

	wsBookTickerHandler := func(event *binance.WsBookTickerEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneChan, stopChan, err := binance.WsBookTickerServe("BTCUSDT", wsBookTickerHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(10 * time.Second)
		stopChan <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneChan

	println("Hello, Gaya!")
}
