package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	yaml "gopkg.in/yaml.v2"

	"github.com/xingjigongsi/carproject/framework/container"
)

type ParseApply struct {
	parseContainer container.InterfaceContainer
	confMaps       map[string]interface{}
	keyDelim       string
	dirPath        string
	parseLock      sync.Mutex
}

func NewParseApply(parmes ...interface{}) (interface{}, error) {
	container := parmes[0].(container.InterfaceContainer)
	dirpath := parmes[1].(string)
	return &ParseApply{
		parseContainer: container,
		confMaps:       make(map[string]interface{}),
		keyDelim:       ".",
		dirPath:        dirpath,
	}, nil
}

func (parseFile *ParseApply) readConf() (map[string]interface{}, error) {
	if _, err := os.Stat(parseFile.dirPath); os.IsNotExist(err) {
		return nil, errors.New("folder" + parseFile.dirPath + " not exist:" + err.Error())
	}
	parseFile.parseLock.Lock()
	defer parseFile.parseLock.Unlock()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	defer watcher.Close()
	err = watcher.Add(parseFile.dirPath)
	if err != nil {
		return nil, err
	}
	errs := make(chan error, 1)
	go func() {
		for {
			select {
			case ev, ok := <-watcher.Events:
				if !ok {
					return
				}
				if err != nil {
					return
				}
				path, _ := filepath.Abs(ev.Name)
				index := strings.LastIndex(path, string(os.PathSeparator))
				folder := path[:index]
				fileName := path[index+1:]
				if ev.Op&fsnotify.Create == fsnotify.Create {
					err := parseFile.loadConfigFile(folder, fileName)
					errs <- err
				}
				if ev.Op&fsnotify.Write == fsnotify.Write {
					err := parseFile.loadConfigFile(folder, fileName)
					errs <- err
				}
				if ev.Op&fsnotify.Remove == fsnotify.Remove {
					delete(parseFile.confMaps, fileName)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				errs <- err
			}
		}
	}()
	go func() {
		fmt.Println(<-errs)
	}()
	files, err := ioutil.ReadDir(parseFile.dirPath)
	for _, file := range files {
		filename := file.Name()
		err := parseFile.loadConfigFile(parseFile.dirPath, filename)
		if err == nil {
			continue
		}
	}
	return parseFile.confMaps, nil

}

func (parseFile *ParseApply) loadConfigFile(dir string, filename string) error {
	s := strings.Split(filename, ".")
	if len(s) != 2 || (s[1] != "yaml" && s[1] != "yml") {
		return errors.New(filename + "not yaml or yml")
	}
	name := s[0]
	file, err := ioutil.ReadFile(filepath.Join(dir, filename))
	if err != nil {
		return err
	}
	c := map[string]interface{}{}
	err = yaml.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	if name == "app" && parseFile.parseContainer.IsBind(container.APPKEY) {
		app := parseFile.parseContainer.MustMake(container.APPKEY).(container.AppInterface)
		app.LoadApplyConfig(cast.ToStringMapString(c["path"]))
	}
	parseFile.confMaps[name] = c
	return nil
}

func (parseFile *ParseApply) find(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}
	next, ok := source[path[0]]
	if !ok {
		return nil
	}
	if len(path) == 1 {
		return next
	}
	switch next.(type) {
	case map[interface{}]interface{}:
		return parseFile.find(cast.ToStringMap(next), path[1:])
	case map[string]interface{}:
		return parseFile.find(next.(map[string]interface{}), path[1:])
	default:
		return nil
	}
	return nil
}

func (parseFile *ParseApply) get(key string) (interface{}, error) {
	source, err := parseFile.readConf()
	if err != nil {
		return nil, err
	}
	find := parseFile.find(source, strings.Split(key, parseFile.keyDelim))
	return find, nil
}

func (parseFile *ParseApply) IsExist(key string) (bool, error) {
	data, err := parseFile.get(key)
	if err != nil {
		return false, err
	}
	return data != nil, nil
}

func (parseFile *ParseApply) GetBool(key string) (bool, error) {
	get, err := parseFile.get(key)
	if err != nil {
		return false, err
	}
	return cast.ToBool(get), nil
}

func (parseFile *ParseApply) GetInt(key string) (int, error) {
	data, err := parseFile.get(key)
	if err != nil {
		return 0, nil
	}
	return cast.ToInt(data), err
}

func (parseFile *ParseApply) GetFloat64(key string) (float64, error) {
	data, err := parseFile.get(key)
	if err != nil {
		return 0, err
	}
	return cast.ToFloat64(data), nil
}

func (parseFile *ParseApply) GetString(key string) (string, error) {
	data, err := parseFile.get(key)
	if err != nil {
		return "", err
	}
	return cast.ToString(data), nil
}

func (parseFile *ParseApply) GetTime(key string) (time.Time, error) {
	data, err := parseFile.get(key)
	if err != nil {
		return time.Now(), err
	}
	return cast.ToTime(data), nil
}
