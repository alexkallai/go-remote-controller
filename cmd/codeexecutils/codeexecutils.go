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
	"logout":   "shutdown /l",
	"startcmd": "cmd",
	"calc":     "calc.exe",
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
