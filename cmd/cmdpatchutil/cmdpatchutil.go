package cmdpatchutil

import (
	"syscall"
)

var (
	kernel32           = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleMode = kernel32.NewProc("SetConsoleMode")
)

const (
	CHECK_QUICK_EDIT = 0x0040
)

// Solution based on
// https://github.com/npocmaka/batch.scripts/blob/master/hybrids/.net/c/quickEdit.bat
func DisableQuickEditMode() error {
	handle, err := syscall.GetStdHandle(syscall.STD_INPUT_HANDLE)
	if err != nil {
		return err
	}

	var mode uint32
	err = syscall.GetConsoleMode(handle, &mode)
	if err != nil {
		return err
	}

	mode &^= CHECK_QUICK_EDIT // Disable quick edit mode

	r, _, err := procSetConsoleMode.Call(uintptr(handle), uintptr(mode))
	if r == 0 {
		return err
	}

	return nil
}
