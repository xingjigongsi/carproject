package util

import "os"

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
