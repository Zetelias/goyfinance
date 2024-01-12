package goyfinance

import (
	"testing"
)

func aaplList(n int) []string {
	var aapl_list []string
	for i := 0; i < n; i++ {
		aapl_list = append(aapl_list, "AAPL")
	}
	return aapl_list
}

func TestGetQuote(t *testing.T) {
	quote, err := GetQuote("AAPL", IntervalOneDay, PeriodOneDay)
	if err != nil {
		t.Error(err)
	}
	if quote.Ticker != "AAPL" {
		t.Error("Ticker is not AAPL")
	}
}

func TestGetQuoteBatch(t *testing.T) {
	quotes, err := GetQuoteBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
	if err != nil {
		t.Error(err)
	}
	if len(quotes) != 10 {
		t.Error("Quotes length is not 10")
	}
}

func BenchmarkGetQuoteJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteJSON("AAPL", IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuoteJSONBatch(b *testing.B) {
	b.StopTimer()
	aaplist := aaplList(10)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		GetQuoteJSONBatch(aaplist, IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuoteJSONString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteJSONString("AAPL", IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuoteJSONStringBatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteJSONStringBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuote(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuote("AAPL", IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuoteBatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuoteCSV(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteCSVString("AAPL", IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuoteCSVBatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteCSVStringBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
	}
}
