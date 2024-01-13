package goyfinance

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func aaplList(n int) []string {
	var aaplList []string
	for i := 0; i < n; i++ {
		aaplList = append(aaplList, "AAPL")
	}
	return aaplList
}

func testAAPLQuote(quote Quote) {
	if quote.Ticker != "AAPL" {
		panic("Ticker is not AAPL")
	}
	var fields []struct {
		name string
		val  interface{}
	}
	for _, field := range fields {
		if field.val == 0 {
			panic(field.name + " is 0")
		}
	}
}

func TestGetQuote(t *testing.T) {
	quote, err := GetQuote("AAPL", IntervalOneDay, PeriodOneDay)
	if err != nil {
		t.Error(err)
	}
	testAAPLQuote(quote)
}

func TestGetQuoteBatch(t *testing.T) {
	quotes, err := GetQuoteBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
	if err != nil {
		t.Error(err)
	}
	for _, quote := range quotes {
		testAAPLQuote(quote)
	}
}

func TestContinuousPriceUpdater(t *testing.T) {
	priceChannel := make(chan PriceData)
	errorChannel := make(chan error)
	stopSignal := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		ContinuousPriceUpdater(priceChannel, errorChannel, "AAPL", IntervalOneDay, PeriodOneDay, 1, stopSignal)
	}()

	select {
	case val := <-priceChannel:
		if val.ClosePrice == 0 || val.OpenPrice == 0 || val.HighPrice == 0 || val.LowPrice == 0 || val.Volume == 0 {
			t.Error("Price data is 0")
		}
	case err := <-errorChannel:
		t.Errorf("Received an error: %v", err)
	case <-time.After(5 * time.Second): // Adjust the timeout as needed
		t.Error("Timeout: Test did not complete within the expected time")
	}

	stopSignal <- struct{}{}

	wg.Wait()
}

func TestExampleUsage(t *testing.T) {
	// Define a bunch of tickers
	tickerList := []string{"AAPL", "MSFT", "GOOG", "TSLA", "AMZN"}

	// Get quote data for all the tickers
	quotes, err := GetQuoteBatch(tickerList, IntervalOneDay, PeriodFiveDays)

	// Check for errors
	if err != nil {
		t.Error(err)
	}

	// Use the quotes
	fmt.Printf("Three days ago, %s closed at $%.2f\n", quotes[0].Ticker, quotes[0].PriceHistoric[2].ClosePrice)
	fmt.Printf("Now, the volume of %s is %d\n", quotes[1].Ticker, quotes[1].PriceHistoric[0].Volume)
}

func BenchmarkGetQuoteJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetQuoteJSON("AAPL", IntervalOneDay, PeriodOneDay)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetQuoteJSONBatch(b *testing.B) {
	b.StopTimer()
	aaplist := aaplList(10)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		_, err := GetQuoteJSONBatch(aaplist, IntervalOneDay, PeriodOneDay)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetQuoteJSONString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetQuoteJSONString("AAPL", IntervalOneDay, PeriodOneDay)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetQuoteJSONStringBatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetQuoteJSONStringBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetQuote(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetQuote("AAPL", IntervalOneDay, PeriodOneDay)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetQuoteBatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetQuoteBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetQuoteCSVString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetQuoteCSVString("AAPL", IntervalOneDay, PeriodOneDay)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetQuoteCSVStringBatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := GetQuoteCSVStringBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
		if err != nil {
			b.Error(err)
		}
	}
}
