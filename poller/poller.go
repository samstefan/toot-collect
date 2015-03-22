package poller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"../db"
	"../models"

	"github.com/kurrik/twittergo"
	"labix.org/v2/mgo/bson"
)

const (
	requestUrl     = "/1.1/statuses/user_timeline.json?%v"
	count      int = 1
)

var (
	dbConnection = os.Getenv("MONGODB")
	query        url.Values
	response     *twittergo.APIResponse
	req          *http.Request
	err          error
	results      *twittergo.Timeline
)

type tweets struct {
	TweetId    string `json:"id_str"`
	User       user   `json:"user"`
	DatePosted string `json:"created_at"`
	Text       string `json:"text"`
}

type user struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

func start(client *twittergo.Client, account account, interval time.Duration) {
	log.Println("Subscribing to: " + account.User)

	go func() {
		ticker := time.NewTicker(interval)
		for _ = range ticker.C {
			poll(client, account)
		}
	}()
}

func poll(client *twittergo.Client, account account) {
	var session = db.Session
	c := session.DB(dbConnection).C("tweets")

	query = url.Values{}
	query.Set("count", fmt.Sprintf("%v", count))
	query.Set("screen_name", account.User)

	endpoint := fmt.Sprintf(requestUrl, query.Encode())
	if req, err = http.NewRequest("GET", endpoint, nil); err != nil {
		fmt.Printf("Could not parse request: %v\n", err)
	} else {
		if response, err = client.SendRequest(req); err != nil {
			fmt.Printf("Could not send request: %v\n", err)
		} else {
			decoder := json.NewDecoder(response.Body)
			var decocedResponse []tweets
			err = decoder.Decode(&decocedResponse)

			if err != nil {
				log.Println(err)
			}

			var tweetToSave models.TweetModel
			tweetToSave.TweetId = decocedResponse[0].TweetId
			tweetToSave.User.Name = decocedResponse[0].User.Name
			tweetToSave.User.ScreenName = decocedResponse[0].User.ScreenName
			tweetToSave.DatePosted, _ = time.Parse(time.RubyDate, decocedResponse[0].DatePosted)
			tweetToSave.Text = decocedResponse[0].Text

			// Upsert the document
			_, err := c.Upsert(bson.M{"tweetId": tweetToSave.TweetId}, bson.M{"$set": tweetToSave})

			if err != nil {
				log.Println(err)
			}
		}

		defer response.Body.Close()
	}
}
