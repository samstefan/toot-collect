package poller

import (
	"log"
	"time"
	"os"
	"encoding/json"

	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

type properties struct {
  Accounts    []account   `json:"accounts"`
  Credentials credentials `json:"credentials"`
}

type account struct {
	User string `json:"user"`
}

type credentials struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
}

func Bootstrap() {
	properties := getProperties()
	accounts := properties.Accounts
	credentials := properties.Credentials
	pollInterval := calculateIntervalTime(len(accounts), 180, 900)
	client := authWithTwitter(credentials)

	for i := 0; i < len(accounts); i++ {
		start(client, accounts[i], pollInterval)
	}
}

func calculateIntervalTime(accounts int, requests int, seconds int) (time.Duration) {
	return time.Duration((seconds / (requests / (accounts + 1))) * 1000) * time.Millisecond
}

func authWithTwitter(credentials credentials) (client *twittergo.Client) {
	config := &oauth1a.ClientConfig{
		ConsumerKey:    credentials.ConsumerKey,
		ConsumerSecret: credentials.ConsumerSecret,
	}

	user := oauth1a.NewAuthorizedConfig(credentials.AccessToken, credentials.AccessTokenSecret)
  client = twittergo.NewClient(config, user)
	return client
}

func getProperties() (results properties) {
	// Open properties file
  propertiesFile, err := os.Open("./properties.json")

  if err != nil {
    log.Println("Error opening properties file", err.Error())
  }

	// Parse the JSON
  jsonParser := json.NewDecoder(propertiesFile)

  if err = jsonParser.Decode(&results); err != nil {
    log.Println("Error parsing properties file", err.Error())
  }

	return results
}