package models

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type TweetModel struct {
	Id         bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	TweetId    string        `bson:"tweetId" json:"tweetId"`
	User       user          `bson:"user" json:"user"`
	DatePosted time.Time     `bson:"datePosted" json:"datePosted"`
	Text       string        `bson:"text" json:"text"`
}

type user struct {
	Name       string `bson:"name" json:"name"`
	ScreenName string `bson:"screenName" json:"screenName"`
}
