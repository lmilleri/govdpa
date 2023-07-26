package config

import (
	"fmt"
	"path/filepath"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct {
	rootPath string // `default:""`
}

var singleInstance *Singleton

func GetInstance() *Singleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating Singleton instance now.")
			singleInstance = &Singleton{}
		} else {
			fmt.Println("Singleton instance already created.")
		}
	} else {
		fmt.Println("Singleton instance already created.")
	}

	return singleInstance
}

func (singleton *Singleton) SetRootPath(rootPath string) {
	singleton.rootPath = rootPath
}

func (singleton *Singleton) GetRootPath() string {
	return singleton.rootPath
}

func (singleton *Singleton) AdjustPath(path string) string {
	if singleton.rootPath != "" {
		return filepath.Join(singleton.rootPath, path)
	}
	return singleton.rootPath
}
