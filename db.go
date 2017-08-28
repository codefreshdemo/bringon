package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

func dbinit() *mgo.Session {
	session, err := mgo.Dial("mongo")
	log.Printf("created session")
	if err != nil {
		panic(err)
	}
	//defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	return session
}
