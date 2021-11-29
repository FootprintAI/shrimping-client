package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

func NewLocalFileConfigure(localfile string) *LocalFileConfigure {
	return &LocalFileConfigure{
		localfile: localfile,
		Data:      &ConfigureData{},
	}
}

type ConfigureData struct {
	Token         string    `json:"token"`
	LastLoginTime time.Time `json:"lastLoginTime"`
}

type LocalFileConfigure struct {
	localfile string
	Data      *ConfigureData
}

func (l *LocalFileConfigure) Load() error {
	if !exists(l.localfile) {
		return errors.New("missing localfile")
	}
	byteValue, err := ioutil.ReadFile(l.localfile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(byteValue, l.Data); err != nil {
		return err
	}
	return nil
}

func (l *LocalFileConfigure) Save() error {
	byteValue, err := json.Marshal(l.Data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(l.localfile, byteValue, 0644)
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
