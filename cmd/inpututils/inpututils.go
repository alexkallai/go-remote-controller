package inpututils

import (
	"log"

	"github.com/alexkallai/user32util"
)

// Initialize user32dll stuff
var dll, _ = user32util.LoadUser32DLL()

func MoveMouseToCoord(x int32, y int32) {
	log.Println("Mouse moved!")
	_, err := user32util.SetCursorPos(x, y, dll)
	if err != nil {
		log.Fatalf("failed to send mouse down input - %s", err)
	}
}

// Adds the supplied values to the current coordinates.
// Positive x_diff moves the pointer right.
// Positive y_diff moves the pointer up.
func MoveMouseRelative(x_diff int32, y_diff int32) {
	x_cur, y_cur, _ := user32util.GetCursorPos(dll)
	_, err := user32util.SetCursorPos(x_cur+x_diff, y_cur-y_diff, dll)
	if err != nil {
		log.Fatalf("failed to send mouse down input - %s", err)
	}
}

func SendKeys(inputStructs []user32util.KeybdInput) {
	log.Println("Sending key")
	for _, val := range inputStructs {
		err := user32util.SendKeydbInput(val, dll)
		if err != nil {
			log.Fatalf("failed to send keyboard input - %s", err)
		}
	}
}

func SendMouseInput(mouseInputStructs []user32util.MouseInput) {
	log.Println("Sending mouse button input")
	for _, val := range mouseInputStructs {
		err := user32util.SendMouseInput(val, dll)
		if err != nil {
			log.Fatalf("failed to send mouse input - %s", err)
		}
	}
}

var EventNameToMouseInputMap map[string][]user32util.MouseInput = map[string][]user32util.MouseInput{
	"leftclick":  {{DwFlags: user32util.MouseEventFLeftDown}, {DwFlags: user32util.MouseEventFLeftUp}},
	"midclick":   {{DwFlags: user32util.MouseEventFMiddleDown}, {DwFlags: user32util.MouseEventFMiddleUp}},
	"rightclick": {{DwFlags: user32util.MouseEventFRightDown}, {DwFlags: user32util.MouseEventFRightUp}},
}

var EventNameToKbdInputMap map[string][]user32util.KeybdInput = map[string][]user32util.KeybdInput{
	"enter":       {{WVK: KeyCodes["VK_RETURN"]}},
	"backspace":   {{WVK: KeyCodes["VK_BACK"]}},
	"delete":      {{WVK: KeyCodes["VK_DELETE"]}},
	"home":        {{WVK: KeyCodes["VK_HOME"]}},
	"end":         {{WVK: KeyCodes["VK_END"]}},
	"printscreen": {{WVK: KeyCodes["VK_SNAPSHOT"]}},
	"pageup":      {{WVK: KeyCodes["VK_PRIOR"]}},
	"pagedown":    {{WVK: KeyCodes["VK_NEXT"]}},
	"contextmenu": {{WVK: KeyCodes["VK_APPS"]}},
	"copy":        {{WVK: KeyCodes["VK_CONTROL"]}, {WVK: KeyCodes["C_key"]}, {WVK: KeyCodes["VK_CONTROL"], DwFlags: user32util.KeyEventFKeyUp}},
	"paste":       {{WVK: KeyCodes["VK_CONTROL"]}, {WVK: KeyCodes["V_key"]}, {WVK: KeyCodes["VK_CONTROL"], DwFlags: user32util.KeyEventFKeyUp}},
	"desktop":     {{WVK: KeyCodes["VK_LWIN"]}, {WVK: KeyCodes["D_key"]}, {WVK: KeyCodes["VK_LWIN"], DwFlags: user32util.KeyEventFKeyUp}},
	"undo":        {{WVK: KeyCodes["VK_CONTROL"]}, {WVK: KeyCodes["Z_key"]}, {WVK: KeyCodes["VK_CONTROL"], DwFlags: user32util.KeyEventFKeyUp}},
	"uparrow":     {{WVK: KeyCodes["VK_UP"]}},
	"selectall":   {{WVK: KeyCodes["VK_CONTROL"]}, {WVK: KeyCodes["A_key"]}, {WVK: KeyCodes["VK_CONTROL"], DwFlags: user32util.KeyEventFKeyUp}},
	"leftarrow":   {{WVK: KeyCodes["VK_LEFT"]}},
	"downarrow":   {{WVK: KeyCodes["VK_DOWN"]}},
	"rightarrow":  {{WVK: KeyCodes["VK_RIGHT"]}},
	//
	"playpause":     {{WVK: KeyCodes["VK_MEDIA_PLAY_PAUSE"]}},
	"previous":      {{WVK: KeyCodes["VK_MEDIA_PREV_TRACK"]}},
	"next":          {{WVK: KeyCodes["VK_MEDIA_NEXT_TRACK"]}},
	"fullscreen":    {{WVK: KeyCodes["VK_CONTROL"]}, {WVK: KeyCodes["VK_RETURN"]}, {WVK: KeyCodes["VK_CONTROL"], DwFlags: user32util.KeyEventFKeyUp}},
	"fullscreen_yt": {{WVK: KeyCodes["F_key"]}},
	"escape":        {{WVK: KeyCodes["VK_ESCAPE"]}},
	"closetab":      {{WVK: KeyCodes["VK_CONTROL"]}, {WVK: KeyCodes["W_key"]}, {WVK: KeyCodes["VK_CONTROL"], DwFlags: user32util.KeyEventFKeyUp}},
	"volumeup":      {{WVK: KeyCodes["VK_VOLUME_UP"]}},
	"mute":          {{WVK: KeyCodes["VK_VOLUME_MUTE"]}},
	"volumedown":    {{WVK: KeyCodes["VK_VOLUME_DOWN"]}},
	"open":          {{WVK: KeyCodes["VK_CONTROL"]}, {WVK: KeyCodes["O_key"]}, {WVK: KeyCodes["VK_CONTROL"], DwFlags: user32util.KeyEventFKeyUp}},
	"closeprogram":  {{WVK: KeyCodes["VK_LMENU"]}, {WVK: KeyCodes["VK_F4"]}},
	"changefocus":   {{WVK: KeyCodes["VK_LMENU"]}, {WVK: KeyCodes["VK_TAB"]}, {WVK: KeyCodes["VK_LMENU"], DwFlags: user32util.KeyEventFKeyUp}},
}
