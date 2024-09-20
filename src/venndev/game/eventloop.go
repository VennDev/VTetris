package game

import "time"

type Loop struct {
	StartTime int64
	Delay     int64
	CallBack  func()
}

// NewLoop creates a new loop with a delay time and a callback function
// delayTime is the time in milliseconds to wait before calling the callback function
func NewLoop(delayTime int64, callBack func()) *Loop {
	return &Loop{
		StartTime: time.Now().UnixMilli(),
		Delay:     delayTime,
		CallBack:  callBack,
	}
}

func (l *Loop) Update() {
	timeNow := time.Now().UnixMilli()
	// Calculate the difference between the current time and the start time
	// Time as a Unix timestamp in nanoseconds
	diff := timeNow - l.StartTime
	if diff > l.Delay {
		l.CallBack()
		l.StartTime = timeNow
	}
}
