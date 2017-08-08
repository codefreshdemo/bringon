package main

import "time"

type Build struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	TimeStamp time.Time `json:"due"`
}

type Builds []Build
