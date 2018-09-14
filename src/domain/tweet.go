package domain

import (
	"time"
)

// Tweet struct
type Tweet struct {
	User string
	Text string
	Date *time.Time
}

// NewTweet nuevo tweet
func NewTweet(user, text string) *Tweet {
	var date = time.Now()
	var tweet = Tweet{User: user, Text: text, Date: &date}
	return &tweet
}
