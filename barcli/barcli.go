// Package barcli implements a cli frontend for progress bars.
package barcli

import "fmt"

import "github.com/0xC3/progress"

// Bar represents a cli frontend of a progress.Bar object.
type Bar struct {
	backend  *progress.Bar
	hasRun   bool
	barCount int
	percent  int
}

// New returns a new Bar object initialized from the given parameters and prints
// the bar.
func New(max int) (bar *Bar, err error) {
	bar = new(Bar)
	bar.backend, err = progress.New(max)
	if err != nil {
		return nil, err
	}
	err = bar.Print()
	if err != nil {
		return nil, err
	}
	return bar, nil
}

// IncMax increases the Max value of bar by add and prints the bar.
func (bar *Bar) IncMax(add int) (err error) {
	err = bar.backend.IncMax(add)
	if err != nil {
		return err
	}
	return bar.Print()
}

// IncN increases the Cur value of bar by add and prints the bar.
func (bar *Bar) IncN(add int) (err error) {
	err = bar.backend.IncN(add)
	if err != nil {
		return err
	}
	return bar.Print()
}

// Inc increases the Cur value of bar by one and prints the bar.
func (bar *Bar) Inc() (err error) {
	err = bar.backend.Inc()
	if err != nil {
		return err
	}
	return bar.Print()
}

// Print prints the current progress bar.
//
// Note: the terminal has to be at least 14 characters in width.
func (bar *Bar) Print() (err error) {
	const prettyFmt = "[%s] %3d%% done"

	//    '%s' -> ''  = -2
	//    '%%' -> '%' = -1
	const prettySize = len(prettyFmt) - 3
	ws := getWinSize()
	barSize := int(ws.Col) - prettySize
	if barSize < 2 {
		return fmt.Errorf("terminal too small (%d) to display progressbar.", ws.Col)
	}
	part := bar.backend.Progress()
	barCount := int(part * float64(barSize))
	percent := int(part * 100)
	if bar.hasRun == true && barCount == bar.barCount && percent == bar.percent {
		return
	}
	bar.hasRun = true
	bar.barCount = barCount
	bar.percent = percent
	slots := make([]byte, barSize)
	for i := 0; i < barSize; i++ {
		if i < barCount {
			slots[i] = byte('=')
		} else {
			slots[i] = byte(' ')
		}
	}
	fmt.Printf("\r")
	fmt.Printf(prettyFmt, string(slots), percent)
	if percent == 100 {
		fmt.Println()
	}
	return nil
}
