package bringon

import (
	"time"

	"github.com/otomato-gh/buildinfo"
)

type Build struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	TimeStamp time.Time `json:"timestamp"`
	Info      buildinfo.BuildInfo `json:"info"` 
}

type Builds []Build
