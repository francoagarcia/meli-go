package service_test

import (
	"testing"

	"github.com/francoagarcia/meli-go/src/domain"
	"github.com/francoagarcia/meli-go/src/service"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	user := "francoagarcia"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()

	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}

	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}

}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	user := "francoagarcia"
	var text string

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet

	user := "grupoesfera"
	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweet(user, text)

	// Operation
	var err error
	err = service.PublishTweet(tweet)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}
