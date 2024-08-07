package main

import (
	"fmt"
	"github.com/mattevans/dinero"
	"time"
)

func currencyPairRate(val1, val2 string, n float64) float64 {
	appID := "b794131bdef549ea8b7991ec50a5baf3"

	client := dinero.NewClient(
		appID,
		val1,
		20*time.Minute,
	)

	rsp, _ := client.Rates.Get(val2)

	return *rsp * n

}

func main() {
	val1 := "USD"
	val2 := "EUR"
	n := 100.0
	fmt.Printf("%v %s = %v %s", n, val1, currencyPairRate(val1, val2, n), val2)
}
