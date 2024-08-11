package main

import (
	"context"
	"fmt"

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

	println("Hello, Gaya!")
}
