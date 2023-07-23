package util

import (
	"os"
	"syscall"
)

func PathIsExist(folder string) bool {
	_, err := os.Stat(folder)
	if err != nil {
		exist := os.IsExist(err)
		if exist {
			return true
		}
		return false
	}
	return true
}

func IsExistProcess(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}
	return true
}
