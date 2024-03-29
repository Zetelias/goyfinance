package goyfinance

import (
	"github.com/mailru/easyjson"
	"time"
)

// ---- Helper functions ----
// These functions are not
// exported and are only used
// internally by the library

// / Gets a unix timestamp for now and for `period` days/mo/years in the past
// / The first timestamp returned is `period` days/mo/years ago and the second period is now
func getUnixTimestamps(period Period) (int64, int64) {
	var pastPeriod time.Time
	var now time.Time

	now = time.Now()
	switch period {
	case PeriodFiveDays:
		pastPeriod = now.AddDate(0, 0, -5)
	case PeriodOneMonth:
		pastPeriod = now.AddDate(0, -1, 0)
	case PeriodThreeMonth:
		pastPeriod = now.AddDate(0, -3, 0)
	case PeriodSixMonth:
		pastPeriod = now.AddDate(0, -6, 0)
	case PeriodOneYear:
		pastPeriod = now.AddDate(-1, 0, 0)
	case PeriodTwoYears:
		pastPeriod = now.AddDate(-2, 0, 0)
	case PeriodFiveYear:
		pastPeriod = now.AddDate(-5, 0, 0)
	case PeriodTenYears:
		pastPeriod = now.AddDate(-10, 0, 0)
	case PeriodYtd:
		pastPeriod = time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	default:
		// Unreachable code but return 69, 420 because I'm so funny
		return 69, 420
	}

	// Convert the time.Time objects to unix timestamps
	return pastPeriod.Unix(), now.Unix()
}

func parseJSONToJSONQuote(jsonData []byte) (JSONQuote, error) {
	var quote JSONQuote
	err := easyjson.Unmarshal(jsonData, &quote)
	return quote, err
}

func parseJSONtoQuote(jsonData []byte, ticker string, period1 int64, period2 int64) (Quote, error) {
	var jsonQuote JSONQuote
	err := easyjson.Unmarshal(jsonData, &jsonQuote)
	if err != nil {
		return Quote{}, err
	}
	return parseJSONQuoteToQuote(jsonQuote, ticker, period1, period2)
}

func parseJSONQuoteToQuote(jsonQuote JSONQuote, ticker string, period1 int64, period2 int64) (Quote, error) {
	var quote Quote
	quote.Ticker = ticker
	quote.PriceRangeStart = period1
	quote.PriceRangeEnd = period2
	quote.Interval = Interval(jsonQuote.Chart.Result[0].Meta.DataGranularity)
	for i := 0; i < len(jsonQuote.Chart.Result[0].Timestamp); i++ {
		var priceData PriceData
		priceData.OpenPrice = jsonQuote.Chart.Result[0].Indicators.Quote[0].Open[i]
		priceData.LowPrice = jsonQuote.Chart.Result[0].Indicators.Quote[0].Low[i]
		priceData.HighPrice = jsonQuote.Chart.Result[0].Indicators.Quote[0].High[i]
		priceData.ClosePrice = jsonQuote.Chart.Result[0].Indicators.Quote[0].Close[i]
		priceData.Volume = jsonQuote.Chart.Result[0].Indicators.Quote[0].Volume[i]
		quote.PriceHistoric = append(quote.PriceHistoric, priceData)
	}
	return quote, nil
}

// ---- Structs definitions ----
// Structs are used to parse the
// JSON data returned by the
// Yahoo Finance API into
// a more usable format

// An auto-generated struct from https://mholt.github.io/json-to-go/
// Although it was made as a middle between the JSON data and the Quote struct,
// you could use it yourself because it's more complete than the Quote struct.
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

// One interval of price data
// for a ticker.
// Contains OHLVC data
type PriceData struct {
	OpenPrice  float64
	LowPrice   float64
	HighPrice  float64
	ClosePrice float64
	Volume     int
}

// Quote is a single quote for a ticker
// Contains the ticker, the
// price range, the interval
// and the price data
type Quote struct {
	Ticker          string
	PriceRangeStart int64 // Unix timestamp of the start of the price range
	PriceRangeEnd   int64 // Unix timestamp of the end of the price range
	Interval        Interval
	PriceHistoric   []PriceData
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
