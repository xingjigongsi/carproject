package back_Command

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/xingjigongsi/carproject/framework/components/parse"
	"github.com/xingjigongsi/carproject/framework/container"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Backend struct {
	backendConfig *backendConfig
	lock          sync.Mutex
}

type backendConfig struct {
	MonitorRefersTime time.Duration
	MonitorFolder     string
	Port              string
}

func initBackendConfig(c container.InterfaceContainer) *backendConfig {
	backendConfig := &backendConfig{
		MonitorRefersTime: 1,
		MonitorFolder:     "",
		Port:              "",
	}
	apply := c.MustMake(parse.PASE_NAME).(*parse.ParseApply)
	if refertime, err := apply.GetInt("app.MonitorRefersTime"); err == nil {
		backendConfig.MonitorRefersTime = time.Duration(refertime)
	}
	if folder, err := apply.GetString("app.MonitorFolder"); err == nil {
		backendConfig.MonitorFolder = folder
	}
	app := c.MustMake(container.APPKEY).(*container.AppApply)
	if backendConfig.MonitorFolder == "" {
		backendConfig.MonitorFolder = app.BaseFolder()
	}
	backendConfig.Port = Port
	return backendConfig

}

func NewBackend(c container.InterfaceContainer) *Backend {
	config := initBackendConfig(c)
	Backend := &Backend{
		backendConfig: config,
	}
	return Backend
}

func (backend *Backend) RebuildBackend() error {
	command := exec.Command("./main", "build")
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	if err := command.Start(); err == nil {
		if err := command.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func (backend *Backend) StartBackend() error {
	//port := ":" + backend.backendConfig.Port
	command := exec.Command("./main", "system", "restart")
	command.Stdout = os.NewFile(0, os.DevNull)
	command.Stderr = os.Stderr
	if err := command.Start(); err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("系统重启成功")
	return nil
}

func (backend *Backend) MoniterFolder() error {
	backend.lock.Lock()
	defer backend.lock.Unlock()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()
	folder := backend.backendConfig.MonitorFolder
	filepath.Walk(folder, func(path string, info fs.FileInfo, err error) error {
		for _, v := range []string{".git", "pid", "log"} {
			if len(path) > 1 && strings.Contains(path, v) {
				return nil
			}
		}
		if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}
		if info != nil && info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext == ".go" || ext == ".yaml" || ext == ".proto" || ext == ".yam" {
			watcher.Add(path)
		}
		return nil
	})
	tick := time.NewTimer(backend.backendConfig.MonitorRefersTime * time.Second)
	tick.Stop()
	for {
		select {
		case <-tick.C:
			if err := backend.RebuildBackend(); err != nil {
				fmt.Println(err)
			}
			if err := backend.StartBackend(); err != nil {
				fmt.Println(err)
			}
			tick.Stop()
		case _, ok := <-watcher.Events:
			if !ok {
				continue
			}
			tick.Reset(backend.backendConfig.MonitorRefersTime)
		case _, ok := <-watcher.Errors:
			if !ok {
				continue
			}
			tick.Reset(backend.backendConfig.MonitorRefersTime)
		}
	}
}
