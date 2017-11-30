package main

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
)

const TWITTER_ACCESS_TOKEN = os.Getenv("TWITTER_ACCESS_TOKEN")
const TWITTER_TOKEN_SECRET = os.Getenv("TWITTER_TOKEN_SECRET")
const TWITTER_CONSUMER_KEY = os.Getenv("TWITTER_CONSUMER_KEY")
const TWITTER_CONSUMER_SECRET = os.Getenv("TWITTER_CONSUMER_SECRET")

func main() {
	// Authenticate Twitter
	anaconda.SetConsumerKey(TWITTER_CONSUMER_KEY)
	anaconda.SetConsumerSecret(TWITTER_CONSUMER_SECRET)
	twitterApi := anaconda.NewTwitterApi(TWITTER_ACCESS_TOKEN, TWITTER_TOKEN_SECRET)

}
