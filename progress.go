// Package progress implements the backend functions for handling progress bars.
package progress

import "fmt"

// Bar represents the backend of a progress bar object.
type Bar struct {
	Cur int
	Max int
}

// New returns a new Bar object initialized from the given parameters.
func New(max int) (bar *Bar, err error) {
	if max < 1 {
		return nil, fmt.Errorf("Max (%d) below 1.", max)
	}
	bar = new(Bar)
	bar.Max = max
	return bar, nil
}

// IncMax increases the Max value of bar by add.
func (bar *Bar) IncMax(add int) (err error) {
	if bar.Max+add < 1 {
		return fmt.Errorf("Max (%d) below 1.", bar.Max+add)
	}
	bar.Max += add
	return nil
}

// IncN increases the Cur value of bar by add.
func (bar *Bar) IncN(add int) (err error) {
	curNew := bar.Cur + add
	if curNew > bar.Max || curNew < 0 {
		return fmt.Errorf("invalid Cur (%d), Max (%d).", curNew, bar.Max)
	}
	bar.Cur += add
	return nil
}

// Inc increases the Cur value of bar.
func (bar *Bar) Inc() (err error) {
	return bar.IncN(1)
}

// Progress returns a float64 representing the current progress (between 0.0 and
// 1.0).
func (bar *Bar) Progress() (progress float64) {
	return float64(bar.Cur) / float64(bar.Max)
}
