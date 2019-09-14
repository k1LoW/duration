package duration

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var units = map[string]float64{
	"ns":      float64(time.Nanosecond),
	"us":      float64(time.Microsecond),
	"ms":      float64(time.Millisecond),
	"s":       float64(time.Second),
	"sec":     float64(time.Second),
	"second":  float64(time.Second),
	"seconds": float64(time.Second),
	"m":       float64(time.Minute),
	"min":     float64(time.Minute),
	"minute":  float64(time.Minute),
	"minutes": float64(time.Minute),
	"h":       float64(time.Hour),
	"hour":    float64(time.Hour),
	"hours":   float64(time.Hour),
	"d":       float64(time.Hour * 24),
	"day":     float64(time.Hour * 24),
	"days":    float64(time.Hour * 24),
	"w":       float64(time.Hour * 24 * 7),
	"week":    float64(time.Hour * 24 * 7),
	"weeks":   float64(time.Hour * 24 * 7),
}

// Parse parses a formatted string and returns the time.Duration value it represents.
func Parse(s string) (time.Duration, error) {
	s = strings.ToLower(s)
	duration := int64(0)
	value := []rune{}
	unit := []rune{}
	negative := false

	for _, token := range s {
		switch {
		case token == '+':
			if len(unit) > 0 {
				d, err := build(negative, value, unit)
				if err != nil {
					return 0, err
				}
				duration = duration + d
				value = []rune{}
				unit = []rune{}
			}

			negative = false
		case token == '-':
			if len(unit) > 0 {
				d, err := build(negative, value, unit)
				if err != nil {
					return 0, err
				}
				duration = duration + d
				value = []rune{}
				unit = []rune{}
			}

			negative = true
		case (token >= '0' && token <= '9'), token == '.':
			if len(unit) > 0 {
				d, err := build(negative, value, unit)
				if err != nil {
					return 0, err
				}
				duration = duration + d
				value = []rune{}
				unit = []rune{}
			}

			value = append(value, token)
		case token == ' ':
		default:
			unit = append(unit, token)
		}
	}
	d, err := build(negative, value, unit)
	if err != nil {
		return 0, err
	}
	duration = duration + d

	return time.Duration(duration), nil
}

func build(negative bool, value, unit []rune) (int64, error) {
	v, err := strconv.ParseFloat(string(value), 64)
	if err != nil {
		return int64(0), err
	}

	u, ok := units[string(unit)]
	if !ok {
		return int64(0), errors.New("invalid unit")
	}
	n := 1.0
	if negative {
		n = -1.0
	}
	return int64(n * v * u), nil
}
