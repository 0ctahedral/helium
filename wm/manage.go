package wm

import (
	"github.com/BurntSushi/xgb/xproto"
)

var (
	// ManagedFrames is a slice of all managed frames
	ManagedFrames []Frame

	// FocusQ is the order of focused frames
	FocusQ []Frame
)

// ById returns a *Frame if the id matches that of a window or it's associated
func ById(id xproto.Window) Frame {
	for _, f := range ManagedFrames {
		if f.Contains(id) {
			return f
		}
	}
	return nil
}

// AddFrame adds the given frame to the slice of Frames s
func AddFrame(f Frame, s []Frame) []Frame {
	return append([]Frame{f}, s...)
}

// RemoveFrame removes a given frame from the wm list of managed frames
func RemoveFrame(f Frame, s []Frame) []Frame {
	ret := s
	for i, f2 := range s {
		if f2.FrameId() == f.FrameId() {
			ret = append(s[:i], s[i+1:]...)
			break
		}
	}
	// fmt.Printf("Just removed frame: %+v\n", ret)
	return ret
}

// GetFocused returns the currently focused frame or nil if there are none
func GetFocused() Frame {
	if len(FocusQ) == 0 {
		return nil
	}
	return FocusQ[0]
}

// IsFocused returns if the given frame is at the front of the FocusQ
func IsFocused(f Frame) bool {
	if len(FocusQ) == 0 {
		return false
	}

	return FocusQ[0].FrameId() == f.FrameId()
}
