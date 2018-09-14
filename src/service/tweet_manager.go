package service

import (
	"errors"

	"github.com/francoagarcia/meli-go/src/domain"
)

// Tweet string
var Tweet *domain.Tweet

// PublishTweet publicar tweet
func PublishTweet(tweet *domain.Tweet) error {
	if tweet.User == "" {
		return errors.New("user is required")
	}
	if tweet.Text == "" {
		return errors.New("text is required")
	}
	if len(tweet.Text) > 140 {
		return errors.New("text exceeds 140 characters")
	}
	Tweet = tweet
	return nil
}

// GetTweet getter tweet
func GetTweet() *domain.Tweet {
	return Tweet
}
