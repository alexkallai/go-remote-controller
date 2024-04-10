package httpserverutils

import (
	"go-remote-controller/cmd/inpututils"
	"log"
	"time"
)

var HOLD_FOR_RIGHT_CLICK_TIMEOUT int32 = 1000 //ms
var CLICK_TOUCHEND_TIMEOUT int32 = 200        //ms
var SENSITIVITY_MULTIPLIER float64 = 2.0
var mouseStateMachine MouseFSM = MouseFSM{
	State: Idle,
}

type MouseState int

// enum in Go
const (
	Idle MouseState = iota
	LeftClick
	RightClick
	MouseMove
	ScrollMove
)

type MouseFSM struct {
	State                MouseState
	StartTime            time.Time
	LastEventX           float64
	LastEventY           float64
	TouchCount           int
	IsLeftClickPossible  bool
	IsRightClickPossible bool
	MouseMoved           bool
	FirstX               float64
	FirstY               float64
}

func ResetStateMachine() {
	mouseStateMachine = MouseFSM{
		State: Idle,
	}
}

func SetSensitivityFromCli(multiplier float64) {
	SENSITIVITY_MULTIPLIER = multiplier
}

// Placed here because of circular import issue
func HandleTouchEvent(e Event) {
	mouseStateMachine.ProcessEvent(e)
}

func (m *MouseFSM) ProcessEvent(e Event) {
	switch e.EventType {
	case "touchmove":
		log.Println("touchmove case handling")
		if m.TouchCount == 1 {
			if m.State == Idle {
				m.State = MouseMove
			}
			cur_x := e.XCoordinateArray[0]
			cur_y := e.YCoordinateArray[0]
			diff_y := (m.LastEventY - cur_y) * SENSITIVITY_MULTIPLIER
			diff_x := (cur_x - m.LastEventX) * SENSITIVITY_MULTIPLIER
			inpututils.MoveMouseRelative(int32(diff_x), int32(diff_y))
			m.LastEventX = cur_x
			m.LastEventY = cur_y
		} else if m.TouchCount == 2 {
			m.State = ScrollMove
			// Calculate scroll movement here based on X and Y coordinates
			log.Println("Scroll move:", e.XCoordinateArray, e.YCoordinateArray)
		}
	case "touchstart":
		// Normal touchstart
		log.Println("touchstart case handling")
		if m.State == Idle && len(e.XCoordinateArray) == 1 {
			m.StartTime = time.Now()
			m.TouchCount = 1
			//m.State = MouseMove
			m.LastEventX = e.XCoordinateArray[0]
			m.LastEventY = e.YCoordinateArray[0]
		} else {
			//ResetStateMachine()
		}
	case "touchend":
		log.Println("touchend case handling")
		duration := time.Since(m.StartTime)
		if duration < time.Millisecond*200 { // Short touch interval for left click
			log.Println("Left click")
			inpututils.SendMouseInput(inpututils.EventNameToMouseInputMap["leftclick"])
		} else if duration < time.Millisecond*1000 && m.TouchCount == 1 { // Longer interval for right click
			log.Println("Right click")
		}
		ResetStateMachine()
	}
}
