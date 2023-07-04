package main

import (
	"fmt"
	"time"
)

type WebServerV2 struct {
	config *Config
	//configLock sync.RWMutex
}

func (ws *WebServerV2) reload() {
	ws.config = &Config{
		fmt.Sprintf("%d", time.Now().UnixNano()),
	}
}
func (ws *WebServerV2) ReloadWorker() {
	for {
		time.Sleep(10 * time.Second)
		ws.reload()
	}
}

func (ws *WebServerV2) Visit() string {
	return ws.config.Content
}
