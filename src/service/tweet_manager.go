package service

import (
	"errors"

	"github.com/francoagarcia/meli-go/src/domain"
)

// Tweets slice de tweets
var Tweets []*domain.Tweet

// InitializeService inicializar slices
func InitializeService() {
	Tweets = make([]*domain.Tweet, 0)
	Usuarios = make([]*domain.Usuario, 0)
}

// PublishTweet publicar tweet
func PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, errors.New("user is required")
	}
	if tweet.Text == "" {
		return 0, errors.New("text is required")
	}
	if len(tweet.Text) > 140 {
		return 0, errors.New("text exceeds 140 characters")
	}
	tweet.ID = len(Tweets) + 1
	Tweets = append(Tweets, tweet)
	return tweet.ID, nil
}

// GetTweet obtener el ultimo tweet publicado
func GetTweet() *domain.Tweet {
	return Tweets[len(Tweets)-1]
}

// GetTweets obtener todos los tweets
func GetTweets() []*domain.Tweet {
	return Tweets
}

// GetTweetByID obtener tweet por id
func GetTweetByID(id int) *domain.Tweet {
	for i := 0; i < len(Tweets); i++ {
		if Tweets[i].ID == id {
			return Tweets[i]
		}
	}
	return nil
}
