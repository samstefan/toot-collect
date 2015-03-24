package routes

import (
	"encoding/json"
	// "log"
	"os"

	"../db"
	"../models"

	"github.com/erkl/robo"
	"labix.org/v2/mgo/bson"
)

var (
	dbConnection = os.Getenv("MONGODB")
)

// List current subscriptions
func getTweet(w robo.ResponseWriter, r *robo.Request) {
	var session = db.Session
	c := session.DB(dbConnection).C("tweets")

	screenName := r.Param("screenName")

	if len(screenName) > 0 {
		result := models.TweetModel{}
		err := c.Find(bson.M{"user.screenName": screenName}).Sort("-datePosted").One(&result)
		if err != nil {
			send(w, "application/json", []byte("Screen name not found"))
		} else {
			jsonData, err := json.Marshal(result)
			if err == nil {
				send(w, "application/json", jsonData)
			}
		}
	} else {
		send(w, "application/json", []byte("No account set"))
	}

}
