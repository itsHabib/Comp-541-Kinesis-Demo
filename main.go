package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

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

	// Authenticate Twitter
	// var twitterData []string
	TWITTER_ACCESS_TOKEN := os.Getenv("TWITTER_ACCESS_TOKEN")
	TWITTER_TOKEN_SECRET := os.Getenv("TWITTER_TOKEN_SECRET")
	TWITTER_CONSUMER_KEY := os.Getenv("TWITTER_CONSUMER_KEY")
	TWITTER_CONSUMER_SECRET := os.Getenv("TWITTER_CONSUMER_SECRET")
	anaconda.SetConsumerKey(TWITTER_CONSUMER_KEY)
	anaconda.SetConsumerSecret(TWITTER_CONSUMER_SECRET)
	twitterAPI := anaconda.NewTwitterApi(TWITTER_ACCESS_TOKEN, TWITTER_TOKEN_SECRET)

	kinesisClient := kinesis.New(awsSession)

	v := url.Values{}
	twitterStreamHandler := twitterAPI.PublicStreamSample(v)
	statusCount := 0
	for statusCount < 10 {
		item := <-twitterStreamHandler.C
		switch status := item.(type) {
		case anaconda.Tweet:
			partitionKey := fmt.Sprintf("partitionKey-%d", statusCount)
			_, err := kinesisClient.PutRecord(&kinesis.PutRecordInput{
				Data:         []byte(strconv.Itoa(status.RetweetCount)),
				StreamName:   &streamName,
				PartitionKey: aws.String(partitionKey),
			})
			if err != nil {
				panic(err)
			}

			statusCount++
		}
	}

}
