package main

import (
	"time"

	"github.com/otomato-gh/buildinfo"
)

type Build struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	TimeStamp time.Time `json:"due"`
	info      buildinfo.BuildInfo
}

type Builds []Build
