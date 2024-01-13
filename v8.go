package goyfinance

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"sync"
	"time"
)

// GetQuoteJSONString returns a JSON string from Yahoo Finance.
// If an error occurs, the JSON string will be empty.
func GetQuoteJSONString(ticker string, interval Interval, period Period) (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	period1, period2 := getUnixTimestamps(period)

	req.SetRequestURI(fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=%s&period1=%d&period2=%d", ticker, interval, period1, period2))
	req.Header.SetMethod("GET")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	err := fasthttp.Do(req, resp)

	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}

// GetQuoteJSONStringBatch returns a slice of JSON strings from Yahoo Finance.
// The order of the slice is the same as the order of the tickers slice.
// If an error occurs, the JSON string will be empty.
func GetQuoteJSONStringBatch(tickers []string, interval Interval, period Period) ([]string, error) {
	var wg sync.WaitGroup
	var res []string
	for _, ticker := range tickers {
		wg.Add(1)
		go func(ticker string) {
			defer wg.Done()
			r, err := GetQuoteJSONString(ticker, interval, period)
			if err != nil {
				return
			}
			res = append(res, r)
		}(ticker)
	}
	wg.Wait()
	return res, nil
}

// GetQuoteJSON returns a JSONQuote struct from Yahoo Finance.
// If an error occurs, the JSONQuote struct will be empty.
func GetQuoteJSON(ticker string, interval Interval, period Period) (JSONQuote, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	period1, period2 := getUnixTimestamps(period)
	req.SetRequestURI(fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=%s&period1=%d&period2=%d", ticker, interval, period1, period2))
	req.Header.SetMethod("GET")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return JSONQuote{}, err
	}

	return parseJSONToJSONQuote(resp.Body())
}

// GetQuoteJSONBatch returns a slice of JSONQuote structs from Yahoo Finance.
// The order of the slice is the same as the order of the tickers slice.
// If an error occurs, the JSONQuote struct will be empty.
func GetQuoteJSONBatch(tickers []string, interval Interval, period Period) ([]JSONQuote, error) {
	var wg sync.WaitGroup
	var res []JSONQuote
	for _, ticker := range tickers {
		wg.Add(1)
		go func(ticker string) {
			defer wg.Done()
			r, err := GetQuoteJSON(ticker, interval, period)
			if err != nil {
				return
			}
			res = append(res, r)
		}(ticker)
	}
	wg.Wait()
	return res, nil
}

// GetQuote returns a Quote struct from Yahoo Finance.
// If an error occurs, the Quote struct will be empty.
// This function is (surprisingly) around the same speed as GetQuoteJSON.
// and a tad faster than GetQuoteJSONString.
func GetQuote(ticker string, interval Interval, period Period) (Quote, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	period1, period2 := getUnixTimestamps(period)
	req.SetRequestURI(fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=%s&period1=%d&period2=%d", ticker, interval, period1, period2))
	req.Header.SetMethod("GET")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		return Quote{}, err
	}

	quote, err := parseJSONtoQuote(resp.Body(), ticker, period1, period2)
	if err != nil {
		return Quote{}, err
	}

	return quote, nil
}

// GetQuoteBatch returns a slice of Quote structs from Yahoo Finance.
// The order of the slice is the same as the order of the tickers slice.
// If an error occurs, the Quote struct will be empty.
// This function is (surprisingly) around the same speed as GetQuoteJSONBatch.
// and a tad faster than GetQuoteJSONStringBatch.
func GetQuoteBatch(tickers []string, interval Interval, period Period) ([]Quote, error) {
	var wg sync.WaitGroup
	var res []Quote
	for _, ticker := range tickers {
		wg.Add(1)
		go func(ticker string) {
			defer wg.Done()
			r, err := GetQuote(ticker, interval, period)
			if err != nil {
				return
			}
			res = append(res, r)
		}(ticker)
	}
	wg.Wait()
	return res, nil
}

// GetQuoteCSVString returns a CSV string with OHLCV data from Yahoo Finance.
// If an error occurs, the CSV string will be empty.
func GetQuoteCSVString(ticker string, interval Interval, period Period) (string, error) {
	req := fasthttp.AcquireRequest()
	period1, period2 := getUnixTimestamps(period)

	req.SetRequestURI(fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/download/%s?interval=%s&period1=%d&period2=%d&events=history", ticker, interval, period1, period2))
	req.Header.SetMethod("GET")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:86.0) Gecko/20100101 Firefox/86.0")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	err := fasthttp.Do(req, resp)

	if err != nil {
		return "", err
	}

	return string(resp.Body()), nil
}

// GetQuoteCSVStringBatch returns a slice of CSV strings with OHLCV data from Yahoo Finance.
// The order of the slice is the same as the order of the tickers slice.
// If an error occurs, the CSV string will be empty.
func GetQuoteCSVStringBatch(tickers []string, interval Interval, period Period) ([]string, error) {
	var wg sync.WaitGroup
	var res []string
	for _, ticker := range tickers {
		wg.Add(1)
		go func(ticker string) {
			defer wg.Done()
			r, err := GetQuoteCSVString(ticker, interval, period)
			if err != nil {
				return
			}
			res = append(res, r)
		}(ticker)
	}
	wg.Wait()
	return res, nil
}

// ContinuousPriceUpdater updates a channel with the latest price data.
// The channel will be updated every updateIntervalSeconds seconds.
// The stopSignal channel is used to stop the function.
// The errorChannel is used to send errors to the caller.
// The priceChannel is used to send the latest price data as a PriceData struct to the caller.
// ticker, interval, period are used to fetch the price data from GetQuote.
func ContinuousPriceUpdater(priceChannel chan PriceData, errorChannel chan error, ticker string, interval Interval, period Period, updateIntervalSeconds float64, stopSignal chan struct{}) {
	for {
		select {
		case <-stopSignal:
			return // We chose not to close all the channels because we don't know if they are used elsewhere
		default:
			price, err := GetQuote(ticker, interval, period)
			if err != nil {
				errorChannel <- err
			} else if len(price.PriceHistoric) > 0 {
				priceChannel <- price.PriceHistoric[len(price.PriceHistoric)-1]
			}
			time.Sleep(time.Duration(updateIntervalSeconds) * time.Second)
		}
	}
}
