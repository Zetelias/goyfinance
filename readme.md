# A fast library useful to get data from yahoo finance in Go
    
## Installation
```bash
go get github.com/Zetelias/goyfinance
```

## Description
This allows you to get data from Yahoo Finance in Go.
It is centered around speed and simplicity.
You can use it to get one, or multiple quotes at once.
You can also create a channel that will send you price data with an interval of your choice.

## Benchmarks
Benchmarks against other libraries are coming soon.

## Example and usage
```go
package main

import (
    "fmt"
    "github.com/Zetelias/goyfinance"
)

func main() {
	// Define a bunch of tickers, which are the ones found on Yahoo Finance.
	tickers := []string{"AAPL", "TSLA", "MSFT", "GOOG", "AMZN"}

	// The period is the amount of days in the past to get data from.
	// It is not working days of the stock market, but actual days.
	// This means if you set the period to 5 days, you will get the data from the last 5 actual days.
	// So, the PriceHistoric of your quote will not have a length of 5, but less.
	period := goyfinance.PeriodFiveDays

	// The interval is the amount of time between each data point.
	// For example, if you set the interval to 1 day, you will get OHLCV data for each day.
	interval := goyfinance.IntervalOneDay

	// You can get a single quote like that, it's pretty simple.
	singleQuote, err := goyfinance.GetQuote(tickers[0], interval, period)
	if err != nil {
		fmt.Printf("Error getting quote: %s\n", err)
	}

	fmt.Printf("singleQuote has ticker %s\n", singleQuote.Ticker)

	// You can get a batch of quotes like that, it's pretty simple.
	// It's asynchronous, so it's as fast as getting a single quote.
	batchQuotes, err := goyfinance.GetQuoteBatch(tickers, interval, period)
	if err != nil {
		fmt.Printf("Error getting quote batch: %s\n", err)
	}

	fmt.Printf("batchQuotes has %d quotes\n", len(batchQuotes))

	// Now do something with the quotes.
	// For example calculate the average of the volume of the batch quotes.
	var totalVolume int
	for _, quote := range batchQuotes {
		totalVolume += quote.PriceHistoric[len(quote.PriceHistoric)-1].Volume
	}
	averageVolume := totalVolume / len(batchQuotes)
	fmt.Printf("Average volume of batch quotes is %d\n", averageVolume)
}
```

## Disclaimer
This uses the free, undocumented Yahoo Finance API which while being free, is not guaranteed to be stable.
The Yahoo Finance API should not be used for commercial purposes,
or for real time data.

## Contributing
Make yer PR's, issues and I'll very probably address them
if they do not break the code and the ethos of speed and simplicity.

## License
[MIT](https://choosealicense.com/licenses/mit/)