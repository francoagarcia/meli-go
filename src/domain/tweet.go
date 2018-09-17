package domain

import (
	"time"
)

// Tweet struct
type Tweet struct {
	User string
	Text string
	Date *time.Time
	ID   int
}

// NewTweet nuevo tweet
func NewTweet(user, text string) *Tweet {
	var date = time.Now()
	var tweet = Tweet{User: user, Text: text, Date: &date}
	return &tweet
}

// PrintableTweet retorna string con tweet
func (t *Tweet) PrintableTweet() string {
	return "@" + t.User + ": " + t.Text
}

//Equals compare current tweet to another
func (t *Tweet) Equals(another *Tweet) bool {
	return t.ID == another.ID
}
