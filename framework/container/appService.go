package container

const APPKEY = "app:con"
const SYSTEFOLDER = "carproject"

type AppInterface interface {
	AppID() string
	Version() string
	BaseFolder() string
	ApplyConfig() string
	LogerFolder() string
	LoadApplyConfig(kv map[string]string)
}
