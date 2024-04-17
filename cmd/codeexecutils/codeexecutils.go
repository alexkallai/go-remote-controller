package codeexecutils

import (
	"os"
	"os/exec"
	"syscall"
)

func ExecuteCommand(command string) {
	cmd := exec.Command("cmd")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c " + os.ExpandEnv(command), HideWindow: true}
	cmd.Env = os.Environ()
	cmd.Start()
}

var EventNameToCommandMap map[string]string = map[string]string{
	"logout":       "shutdown /l",
	"shutdown":     "shutdown /s /f /t 0",
	"restart":      "shutdown /r /f /t 0",
	"sleep":        "powercfg -h off && rundll32.exe powrprof.dll,SetSuspendState 0,1,0",
	"controlpanel": "control",
	"chrome":       "start chrome",
	"firefox":      "start firefox",
	"lock":         "rundll32.exe user32.dll,LockWorkStation",
	"taskmgr":      "taskmgr",
	"notepad":      "start notepad",
	"calc":         "start calc",
	"cmd":          "start cmd",
}

func IsEventStringInCommandMap(eventString string) bool {
	found := false
	for key := range EventNameToCommandMap {
		if key == eventString {
			found = true
			break
		}
	}
	return found
}
