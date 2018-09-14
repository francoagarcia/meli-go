package service

// Tweet string
var Tweet string

// PublishTweet publicar tweet
func PublishTweet(tweet string) string {
	Tweet = tweet
	return Tweet
}
