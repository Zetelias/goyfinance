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

func BenchmarkGetQuoteJSON(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteJSON("AAPL", IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuoteJSONNetHTTP(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteJSONNetHTTP("AAPL", IntervalOneDay, PeriodOneDay)
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

func BenchmarkGetQuoteJSONNetHTTPBatch(b *testing.B) {
	b.StopTimer()
	aaplist := aaplList(10)
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		GetQuoteJSONNetHTTPBatch(aaplist, IntervalOneDay, PeriodOneDay)
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
		GetQuoteCSV("AAPL", IntervalOneDay, PeriodOneDay)
	}
}

func BenchmarkGetQuoteCSVBatch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetQuoteCSVBatch(aaplList(10), IntervalOneDay, PeriodOneDay)
	}
}
