package barcli

import "syscall"
import "unsafe"

// winSize contains the dimentions of a terminal.
type winSize struct {
	Row, Col       uint16
	XPixel, YPixel uint16
}

// getWinSize returns a winSize for the current terminal.
func getWinSize() winSize {
	ws := winSize{}
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
	return ws
}
