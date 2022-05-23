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
		{"3hour", false, time.Duration(3 * time.Hour)},
		{"3h", false, time.Duration(3 * time.Hour)},
		{"3Hour", false, time.Duration(3 * time.Hour)},
		{"3H", false, time.Duration(3 * time.Hour)},
		{"+ 3 hour", false, time.Duration(3 * time.Hour)},
		{"- 3 hour", false, time.Duration(-3 * time.Hour)},
		{"12 days", false, time.Duration(12 * 24 * time.Hour)},
		{"4min", false, time.Duration(4 * time.Minute)},
		{"5 hour 6 min", false, time.Duration(5*time.Hour + 6*time.Minute)},
		{"7 hour - 9sec", false, time.Duration(7*time.Hour - 9*time.Second)},
		{"7day+4days-3days", false, time.Duration(8 * 24 * time.Hour)},
		{"0.5 hour", false, time.Duration(30 * time.Minute)},
		{"0.5 sec", false, time.Duration(500 * time.Millisecond)},
		{"0.05 sec", false, time.Duration(50 * time.Millisecond)},
		{"0.005 sec", false, time.Duration(5 * time.Millisecond)},
		// want error
		{"2 years", true, time.Duration(0)},
		{"7", true, time.Duration(0)},
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
