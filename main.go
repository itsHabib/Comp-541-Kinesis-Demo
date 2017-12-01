package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func main() {

	// Authenticate AWS
	region := "us-east-1"
	awsSession := session.New(&aws.Config{Region: &region})
	streamName := "541-Demo-Stream"

	// Twitter Credentials
	TWITTER_ACCESS_TOKEN := os.Getenv("TWITTER_ACCESS_TOKEN")
	TWITTER_TOKEN_SECRET := os.Getenv("TWITTER_TOKEN_SECRET")
	TWITTER_CONSUMER_KEY := os.Getenv("TWITTER_CONSUMER_KEY")
	TWITTER_CONSUMER_SECRET := os.Getenv("TWITTER_CONSUMER_SECRET")

	// Authenticate Twitter
	anaconda.SetConsumerKey(TWITTER_CONSUMER_KEY)
	anaconda.SetConsumerSecret(TWITTER_CONSUMER_SECRET)
	twitterAPI := anaconda.NewTwitterApi(TWITTER_ACCESS_TOKEN, TWITTER_TOKEN_SECRET)

	// Kinesis Client used for Putting records, reading records, etc..
	kinesisClient := kinesis.New(awsSession)

	v := url.Values{}
	twitterStreamHandler := twitterAPI.PublicStreamSample(v)
	statusCount := 0
	for statusCount < 10 {
		// tweet data comes in through a channel
		item := <-twitterStreamHandler.C
		switch status := item.(type) {
		case anaconda.Tweet:
			partitionKey := fmt.Sprintf("partitionKey-%d", statusCount)
			_, err := kinesisClient.PutRecord(&kinesis.PutRecordInput{
				Data:         []byte(status.Text),
				StreamName:   &streamName,
				PartitionKey: aws.String(partitionKey),
			})
			if err != nil {
				panic(err)
			}
			log.Println(status.Text)
			statusCount++
		}
	}

}
