package db

import (
	"labix.org/v2/mgo"
)

var (
	Session *mgo.Session
)

func Connect(url string) {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}

	Session = session
}
