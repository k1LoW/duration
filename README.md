# duration [![GitHub Actions](https://action-badges.now.sh/k1LoW/duration)](https://github.com/k1LoW/duration/actions) [![codecov](https://codecov.io/gh/k1LoW/duration/branch/master/graph/badge.svg)](https://codecov.io/gh/k1LoW/duration) [![GoDoc](https://godoc.org/github.com/k1LoW/duration?status.svg)](https://godoc.org/github.com/k1LoW/duration)

`duration.Parse()` parses a formatted string and returns the time.Duration value it represents.

## Supported unit of time

| Unit of time | value |
| --- | --- |
| ns | time.Nanosecond |
| us | time.Microsecond |
| ms | time.Millisecond |
| s | time.Second |
| sec | time.Second |
| second | time.Second |
| seconds | time.Second |
| m | time.Minute |
| min | time.Minute |
| minute | time.Minute |
| minutes | time.Minute |
| h | time.Hour |
| hour | time.Hour |
| hours | time.Hour |
| d | time.Hour * 24 |
| day | time.Hour * 24 |
| days | time.Hour * 24 |
| w | time.Hour * 24 * 7 |
| week | time.Hour * 24 * 7 |
| weeks | time.Hour * 24 * 7 |
