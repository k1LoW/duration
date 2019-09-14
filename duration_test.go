package duration

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in      string
		wantErr bool
		want    time.Duration
	}{
		{"3 hour", false, time.Duration(3 * time.Hour)},
		{"+ 3 hour", false, time.Duration(3 * time.Hour)},
		{"- 3 hour", false, time.Duration(-3 * time.Hour)},
		{"2 days", false, time.Duration(2 * 24 * time.Hour)},
		{"4min", false, time.Duration(4 * time.Minute)},
		{"5 hour 6 min", false, time.Duration(5*time.Hour + 6*time.Minute)},
		{"7 hour - 9sec", false, time.Duration(7*time.Hour - 9*time.Second)},
	}

	for _, tt := range tests {
		got, err := Parse(tt.in)
		if !tt.wantErr && err != nil {
			t.Fatalf("Parse(\"%s\") error: %v", tt.in, err)
		}
		if got != tt.want {
			t.Errorf("Parse(\"%s\"): got %d want %d", tt.in, got, tt.want)
		}
	}
}
