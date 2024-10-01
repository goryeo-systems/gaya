package deribitapi

import (
	"os"

	"github.com/frankrap/deribit-api"
	"github.com/frankrap/deribit-api/models"
	"github.com/goryeo-systems/gaya/pkg/exchangeclient"
	"github.com/goryeo-systems/gaya/pkg/util"
)

type DeribitClient struct {
	client *deribit.Client
}

func New() *DeribitClient {
	cfg := &deribit.Configuration{
		Addr:          deribit.RealBaseURL,
		ApiKey:        os.Getenv("DERIBIT_API_KEY"),
		SecretKey:     os.Getenv("DERIBIT_SECRET_KEY"),
		AutoReconnect: true,
		DebugMode:     true,
	}

	return &DeribitClient{
		client: deribit.New(cfg),
	}
}

func toTickerEvent(event *models.TickerNotification) (*exchangeclient.TickerEvent, error) {
	// TODO: symbol
	symbol := exchangeclient.BtcPerpetual

	return &exchangeclient.TickerEvent{
		Symbol:       symbol,
		BestBidPrice: util.FloatToBigFloat(event.BestBidPrice),
		BestBidQty:   util.FloatToBigFloat(event.BestBidAmount),
		BestAskPrice: util.FloatToBigFloat(event.BestAskPrice),
		BestAskQty:   util.FloatToBigFloat(event.BestAskAmount),
	}, nil
}

func (c *DeribitClient) TickerStream(
	s exchangeclient.Symbol,
	handler exchangeclient.TickerStreamHandler,
	errHandler exchangeclient.ErrHandler,
) error {
	// TODO: s to stream

	c.client.On("ticker.BTC-PERPETUAL.raw", func(e *models.TickerNotification) {
		tickerEvent, err := toTickerEvent(e)
		if err != nil {
			errHandler(err)

			return
		}

		handler(tickerEvent)
	})

	c.client.Subscribe([]string{
		"ticker.BTC-PERPETUAL.raw",
	})

	return nil
}

func (c *DeribitClient) GetWallet() (*exchangeclient.Wallet, error) {
	// TODO: implement this
	return nil, nil
}

/*
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
*/
