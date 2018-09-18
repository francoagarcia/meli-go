package service

import (
	"errors"

	"github.com/francoagarcia/meli-go/src/domain"
)

// TweetManager tweet manager
type TweetManager struct {
	LastPublished domain.Tweet
	Tweets        []domain.Tweet
	TweetsMap     map[string][]domain.Tweet
}

// NewTweetManager inicializar slices
func NewTweetManager(tweetWriter *ChannelTweetWriter) *TweetManager {
	tweetManager := TweetManager{Tweets: make([]domain.Tweet, 0)}
	tweetManager.TweetsMap = make(map[string][]domain.Tweet)
	return &tweetManager
}

// PublishTweet publicar tweet
func (t *TweetManager) PublishTweet(tweet domain.Tweet, quit chan bool) (int, error) {
	if tweet.GetUser() == "" {
		return 0, errors.New("user is required")
	}
	if tweet.GetText() == "" {
		return 0, errors.New("text is required")
	}
	if len(tweet.GetText()) > 140 {
		return 0, errors.New("text exceeds 140 characters")
	}
	//if !service.IsUserRegistered(tweet.User) {
	//	return 0, errors.New("user not registered cannot tweet")
	//}
	tweet.SetID(len(t.Tweets) + 1)
	t.LastPublished = tweet
	t.Tweets = append(t.Tweets, tweet)
	t.TweetsMap[tweet.GetUser()] = append(t.TweetsMap[tweet.GetUser()], tweet)
	return tweet.GetID(), nil
}

// GetTweet obtener el ultimo tweet publicado
func (t *TweetManager) GetTweet() domain.Tweet {
	return t.LastPublished
}

// GetTweets obtener todos los tweets
func (t *TweetManager) GetTweets() []domain.Tweet {
	return t.Tweets
}

// GetTweetByID obtener tweet por id
func (t *TweetManager) GetTweetByID(id int) domain.Tweet {
	tweetBuscado := findBy(t.GetTweets(), func(tweet domain.Tweet) bool {
		return tweet.GetID() == id
	})
	if len(tweetBuscado) > 0 {
		return tweetBuscado[0]
	}
	return nil
}

// GetTweetsByUser filtrar tweets de usuarios
func (t *TweetManager) GetTweetsByUser(user string) (ret []domain.Tweet) {
	return t.TweetsMap[user]
}

// CountTweetsByUser contar tweets por username
func (t *TweetManager) CountTweetsByUser(user string) int {
	return len(t.GetTweetsByUser(user))
}

func findBy(tweets []domain.Tweet, criteria func(domain.Tweet) bool) (ret []domain.Tweet) {
	for _, tweet := range tweets {
		if criteria(tweet) {
			ret = append(ret, tweet)
		}
	}
	return
}
