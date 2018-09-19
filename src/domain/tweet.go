package domain

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Tweet interfaz
type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	GetID() *uuid.UUID
	PrintableTweet() string
	String() string
}

/******************************
 ********* Text Tweet *********
 ******************************/

// TextTweet struct
type TextTweet struct {
	User string     `json:"user"`
	Text string     `json:"text"`
	Date *time.Time `json:"date"`
	ID   *uuid.UUID `json:"id"`
}

// NewTextTweet nuevo tweet
func NewTextTweet(user, text string) *TextTweet {
	var date = time.Now()
	var uuid, _ = uuid.NewV4()
	var tweet = TextTweet{User: user, Text: text, Date: &date, ID: &uuid}
	return &tweet
}

// GetUser obtener usuario
func (t *TextTweet) GetUser() string {
	return t.User
}

// GetText obtener texto del tweet
func (t *TextTweet) GetText() string {
	return t.Text
}

// GetDate obtener fecha del tweet
func (t *TextTweet) GetDate() *time.Time {
	return t.Date
}

// GetID obtener Id del tweet
func (t *TextTweet) GetID() *uuid.UUID {
	return t.ID
}

// PrintableTweet retorna string con tweet
func (t *TextTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}

func (t *TextTweet) String() string {
	return t.PrintableTweet()
}

//Equals compare current tweet to another
func (t *TextTweet) Equals(another Tweet) bool {
	return uuid.Equal(*t.GetID(), *another.GetID())
}

/******************************
 ********* Image Tweet ********
 ******************************/

// ImageTweet struct
type ImageTweet struct {
	TextTweet
	URL string `json:"url"`
}

// NewImageTweet nuevo tweet
func NewImageTweet(user, text, url string) *ImageTweet {
	textTweet := NewTextTweet(user, text)
	imageTweet := ImageTweet{TextTweet: *textTweet, URL: url}
	return &imageTweet
}

// PrintableTweet print tweet
func (t *ImageTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s %s", t.User, t.Text, t.URL)
}

/******************************
 ********* Quote Tweet ********
 ******************************/

// QuoteTweet struct
type QuoteTweet struct {
	TextTweet
	QuotedTweet Tweet `json:"quoted_tweet"`
}

// NewQuoteTweet nuevo tweet
func NewQuoteTweet(user, text string, quotedTweet Tweet) *QuoteTweet {
	mainTweet := NewTextTweet(user, text)
	quoteTweet := QuoteTweet{TextTweet: *mainTweet, QuotedTweet: quotedTweet}
	return &quoteTweet
}

// PrintableTweet print tweet
func (qt *QuoteTweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s \"@%s: %s\"", qt.User, qt.Text, qt.QuotedTweet.GetUser(), qt.QuotedTweet.GetText())
}
