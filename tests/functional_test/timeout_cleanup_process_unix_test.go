//go:build !windows

package functional_test

import "syscall"

func timeoutCleanupProcessRunning(pid int) bool {
	err := syscall.Kill(pid, 0)
	return err == nil || err == syscall.EPERM
}

func timeoutCleanupTerminateProcess(pid int) {
	_ = syscall.Kill(pid, syscall.SIGKILL)
}
