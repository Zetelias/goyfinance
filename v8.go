package goyfinance

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"sync"
)

func GetQuoteJSONString(ticker string, interval Interval, period Period) (string, error) {
	req := fasthttp.AcquireRequest()
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

	json_quote, err := parseJSONToJSONQuote(resp.Body())
	if err != nil {
		return Quote{}, err
	}

	return parseJSONQuoteToQuote(json_quote, ticker, period1, period2)
}

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

func GetQuoteCSV(ticker string, interval Interval, period Period) (string, error) {
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

func GetQuoteCSVBatch(tickers []string, interval Interval, period Period) ([]string, error) {
	var wg sync.WaitGroup
	var res []string
	for _, ticker := range tickers {
		wg.Add(1)
		go func(ticker string) {
			defer wg.Done()
			r, err := GetQuoteCSV(ticker, interval, period)
			if err != nil {
				return
			}
			res = append(res, r)
		}(ticker)
	}
	wg.Wait()
	return res, nil
}
