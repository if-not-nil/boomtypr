package typing

import "time"

// Stats tracks typing test statistics.
type Stats struct {
	StartTime    time.Time
	EndTime      time.Time
	TotalChars   int
	CorrectChars int
	TotalWords   int
}

// NewStats creates a new Stats instance.
func NewStats() *Stats {
	return &Stats{}
}

// Start begins tracking stats.
func (s *Stats) Start() {
	// TODO: implement
}

// Stop ends tracking stats.
func (s *Stats) Stop() {
	// TODO: implement
}

// WPM calculates words per minute.
func (s *Stats) WPM() float64 {
	// TODO: implement
	return 0
}

// Accuracy calculates typing accuracy percentage.
func (s *Stats) Accuracy() float64 {
	// TODO: implement
	return 0
}

// Reset resets all stats.
func (s *Stats) Reset() {
	// TODO: implement
}
