package exec

import (
	"os/exec"
	"syscall"
)

func RunWithoutWindow(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return command.Run()
}

func OutputWithoutWindow(cmd string, args ...string) ([]byte, error) {
	command := exec.Command(cmd, args...)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return command.Output()
}
