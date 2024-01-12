package goyfinance

import (
	"encoding/json"
	"time"
)

// ---- Helper functions ----
// These functions are not
// exported and are only used
// internally by the library

// / Gets a unix timestamp for now and for `period` days/mo/years in the past
// / The first timestamp returned is `period` days/mo/years ago and the second period is now
func getUnixTimestamps(period Period) (int64, int64) {
	var past_period time.Time
	var now time.Time

	now = time.Now()
	switch period {
	case PeriodOneDay:
		// Subtract 1 day from now
		// The pattern is the same for all cases
		past_period = now.AddDate(0, 0, -1)
	case PeriodFiveDays:
		past_period = now.AddDate(0, 0, -5)
	case PeriodOneMonth:
		past_period = now.AddDate(0, -1, 0)
	case PeriodThreeMonth:
		past_period = now.AddDate(0, -3, 0)
	case PeriodSixMonth:
		past_period = now.AddDate(0, -6, 0)
	case PeriodOneYear:
		past_period = now.AddDate(-1, 0, 0)
	case PeriodTwoYears:
		past_period = now.AddDate(-2, 0, 0)
	case PeriodFiveYear:
		past_period = now.AddDate(-5, 0, 0)
	case PeriodTenYears:
		past_period = now.AddDate(-10, 0, 0)
	case PeriodYtd:
		past_period = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	default:
		// Unreachable code but return 69, 420 because i'm so funny
		return 69, 420
	}

	// Convert the time.Time objects to unix timestamps
	return past_period.Unix(), now.Unix()
}

func parseJSONToJSONQuote(json_data []byte) (JSONQuote, error) {
	var quote JSONQuote
	err := json.Unmarshal(json_data, &quote)
	return quote, err
}

func parseJSONQuoteToQuote(quote JSONQuote, ticker string, period1 int64, period2 int64) (Quote, error) {
	var q Quote
	q.Ticker = ticker
	q.PriceRangeStart = period1
	q.PriceRangeEnd = period2
	q.Interval = Interval(quote.Chart.Result[0].Meta.DataGranularity)
	q.OpenPrice = quote.Chart.Result[0].Indicators.Quote[0].Open
	q.LowPrice = quote.Chart.Result[0].Indicators.Quote[0].Low
	q.HighPrice = quote.Chart.Result[0].Indicators.Quote[0].High
	q.ClosePrice = quote.Chart.Result[0].Indicators.Quote[0].Close
	q.Volume = quote.Chart.Result[0].Indicators.Quote[0].Volume
	return q, nil
}

// ---- Structs definitions ----
// These  structs
// are used to parse the JSON
// response from Yahoo Finance

type JSONQuote struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PreviousClose        float64 `json:"previousClose"`
				Scale                int     `json:"scale"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				TradingPeriods [][]struct {
					Timezone  string `json:"timezone"`
					Start     int    `json:"start"`
					End       int    `json:"end"`
					Gmtoffset int    `json:"gmtoffset"`
				} `json:"tradingPeriods"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp  []int `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Open   []float64 `json:"open"`
					Low    []float64 `json:"low"`
					Volume []int     `json:"volume"`
					High   []float64 `json:"high"`
					Close  []float64 `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
		Error any `json:"error"`
	} `json:"chart"`
}

type CSVQuote struct {
	Date     string  `csv:"Date"`
	Open     float64 `csv:"Open"`
	High     float64 `csv:"High"`
	Low      float64 `csv:"Low"`
	Close    float64 `csv:"Close"`
	AdjClose float64 `csv:"Adj Close"`
	Volume   int     `csv:"Volume"`
}

// Our own custom struct for convenience
type Quote struct {
	Ticker          string
	PriceRangeStart int64 // Unix timestamp of the start of the price range
	PriceRangeEnd   int64 // Unix timestamp of the end of the price range
	Interval        Interval
	OpenPrice       []float64
	LowPrice        []float64
	HighPrice       []float64
	ClosePrice      []float64
	Volume          []int
}

// ---- Enum definitions ----
// These enums are used to specify
// the interval and period of the
// stock data to be retrieved.
// It brings safety to the library
// by preventing the user from
// passing invalid values.

type Interval string

const (
	IntervalOneMinute      Interval = "1m"
	IntervalTwoMinutes     Interval = "2m"
	IntervalFiveMinutes    Interval = "5m"
	IntervalFifteenMinutes Interval = "15m"
	IntervalThirtyMinutes  Interval = "30m"
	IntervalSixtyMinutes   Interval = "60m"
	IntervalNinetyMinutes  Interval = "90m"
	IntervalOneHour        Interval = "1h"
	IntervalFourHours      Interval = "4h"
	IntervalOneDay         Interval = "1d"
	IntervalFiveDays       Interval = "5d"
	IntervalOneWeek        Interval = "1wk"
	IntervalOneMonth       Interval = "1mo"
	IntervalThreeMonths    Interval = "3mo"
)

type Period string

const (
	PeriodOneDay     Period = "1d"
	PeriodFiveDays   Period = "5d"
	PeriodOneMonth   Period = "1mo"
	PeriodThreeMonth Period = "3mo"
	PeriodSixMonth   Period = "6mo"
	PeriodOneYear    Period = "1y"
	PeriodTwoYears   Period = "2y"
	PeriodFiveYear   Period = "5y"
	PeriodTenYears   Period = "10y"
	PeriodYtd        Period = "ytd"
)
