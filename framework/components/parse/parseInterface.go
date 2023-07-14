package parse

import "time"

const PASE_NAME = "app:parse"

type ParseServiceInterface interface {
	IsExist(key string) (bool, error)
	GetBool(key string) (bool, error)
	GetInt(key string) (int, error)
	GetFloat64(key string) (float64, error)
	GetString(key string) (string, error)
	GetTime(key string) (time.Time, error)
}
