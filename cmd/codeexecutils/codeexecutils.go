package codeexecutils

import (
	"os"
	"os/exec"
	"syscall"
)

func ExecuteCommand(command string) (string, string, int) {
	cmd := exec.Command("cmd")
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: "/c " + os.ExpandEnv(command), HideWindow: true}
	cmd.Env = os.Environ()

	// Capture stdout and stderr
	output, err := cmd.CombinedOutput()

	var stdout, stderr string
	if output != nil {
		stdout = string(output)
	}

	if err != nil {
		stderr = err.Error()
	}

	// Get the exit code
	exitCode := 0
	if exitError, ok := err.(*exec.ExitError); ok {
		exitCode = exitError.ExitCode()
	}

	return stdout, stderr, exitCode
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
