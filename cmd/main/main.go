package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/goryeo-systems/gaya/pkg/binanceapi"
	"github.com/goryeo-systems/gaya/pkg/deribitapi"
	"github.com/goryeo-systems/gaya/pkg/exchangeclient"
	"github.com/goryeo-systems/gaya/pkg/util"
)

var (
	// Variables to store the current bid/ask for Binance and Deribit
	binanceBestBid *big.Float
	binanceBestAsk *big.Float
	deribitBestBid *big.Float
	deribitBestAsk *big.Float
)

func init() {
	// Initialize big.Float variables with 0
	binanceBestBid = new(big.Float).SetFloat64(0)
	binanceBestAsk = new(big.Float).SetFloat64(0)
	deribitBestBid = new(big.Float).SetFloat64(0)
	deribitBestAsk = new(big.Float).SetFloat64(0)
}

// calculateAndPrintSpread calculates the spread and logs it
func calculateAndPrintSpread() {
	// Ensure both exchanges have bid and ask data
	if binanceBestBid.Sign() > 0 && binanceBestAsk.Sign() > 0 && deribitBestBid.Sign() > 0 && deribitBestAsk.Sign() > 0 {
		// Calculate the bid and ask spreads between Binance and Deribit
		bidSpread := new(big.Float).Sub(binanceBestBid, deribitBestAsk) // Binance bid - Deribit ask
		askSpread := new(big.Float).Sub(deribitBestBid, binanceBestAsk) // Deribit bid - Binance ask

		// Log the spreads
		fmt.Println("spread", bidSpread, askSpread)
		/*
			util.Log.Info("SPREAD",
				"BinanceBid-DeribitAsk", bidSpread,
				"DeribitBid-BinanceAsk", askSpread,
			)
		*/
	}
}

// tickerEventHandler handles Binance ticker events
func tickerEventHandler(event *exchangeclient.TickerEvent) {
	// Update the Binance bid/ask prices
	binanceBestBid = event.BestBidPrice
	binanceBestAsk = event.BestAskPrice

	//util.Log.Info("BINANCE", "event", event)

	// Calculate and print the spread
	calculateAndPrintSpread()
}

// deribitTickerEventHandler handles Deribit ticker events
func deribitTickerEventHandler(event *exchangeclient.TickerEvent) {
	// Update the Deribit bid/ask prices
	deribitBestBid = event.BestBidPrice
	deribitBestAsk = event.BestAskPrice

	//util.Log.Info("DERIBIT", "event", event)

	// Calculate and print the spread
	calculateAndPrintSpread()
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
	time.Sleep(30 * time.Second) //nolint:all
}
