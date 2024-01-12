**Operating System:** windows  
**Architecture:** amd64  
**Package:** goyfinance  
**CPU:** Intel(R) Core(TM) i9-10900K CPU @ 3.70GHz

| Benchmark                                    | Iterations | Duration/op     | Bytes/op   | Allocs/op |
|----------------------------------------------|------------|-----------------|------------|-----------|
| GetQuoteJSON-20                              | 525        | 22,151,337 ns/op | 2,543 B/op | 84 allocs/op |
| GetQuoteJSONBatch-20                         | 525        | 24,477,850 ns/op | 28,609 B/op | 874 allocs/op |
| GetQuoteJSONString-20                        | 343        | 38,779,127 ns/op | 6,377 B/op | 44 allocs/op |
| GetQuoteJSONStringBatch-20                   | 458        | 24,495,725 ns/op | 67,884 B/op | 478 allocs/op |
| GetQuote-20                                  | 357        | 38,924,958 ns/op | 2,654 B/op | 86 allocs/op |
| GetQuoteBatch-20                             | 484        | 24,895,384 ns/op | 31,124 B/op | 898 allocs/op |
| GetQuoteCSV-20                               | 133        | 93,167,141 ns/op | 3,789 B/op | 46 allocs/op |
| GetQuoteCSVBatch-20                          | 52         | 227,592,019 ns/op | 39,650 B/op | 494 allocs/op |
