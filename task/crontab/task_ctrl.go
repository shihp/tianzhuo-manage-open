package crontab

import "sync"

var Ctrl *CronMutex

type CronMutex struct {
	Mutex *sync.Mutex
}
