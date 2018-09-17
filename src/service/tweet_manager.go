package service

import (
	"errors"
	"sort"

	"github.com/francoagarcia/meli-go/src/domain"
)

// TweetManager tweet manager
type TweetManager struct {
	LastPublished *domain.Tweet
	Tweets        []*domain.Tweet
	TweetsMap     map[string][]*domain.Tweet
}

// NewTweetManager inicializar slices
func NewTweetManager() *TweetManager {
	tweetManager := TweetManager{Tweets: make([]*domain.Tweet, 0)}
	tweetManager.TweetsMap = make(map[string][]*domain.Tweet)
	return &tweetManager
}

// PublishTweet publicar tweet
func (t *TweetManager) PublishTweet(tweet *domain.Tweet) (int, error) {
	if tweet.User == "" {
		return 0, errors.New("user is required")
	}
	if tweet.Text == "" {
		return 0, errors.New("text is required")
	}
	if len(tweet.Text) > 140 {
		return 0, errors.New("text exceeds 140 characters")
	}
	//if !service.IsUserRegistered(tweet.User) {
	//	return 0, errors.New("user not registered cannot tweet")
	//}
	tweet.ID = len(t.Tweets) + 1
	t.LastPublished = tweet
	t.Tweets = append(t.Tweets, tweet)
	t.TweetsMap[tweet.User] = append(t.TweetsMap[tweet.User], tweet)
	return tweet.ID, nil
}

// GetTweet obtener el ultimo tweet publicado
func (t *TweetManager) GetTweet() *domain.Tweet {
	return t.LastPublished
}

// GetTweets obtener todos los tweets
func (t *TweetManager) GetTweets() []*domain.Tweet {
	return t.Tweets
}

// MapToSlice arma una lista con todos los tweets del map y los ordena por fecha
func (t *TweetManager) MapToSlice() []*domain.Tweet {
	allTweets := make([]*domain.Tweet, 0)
	for _, tweetsDeUsuario := range t.TweetsMap {
		allTweets = append(allTweets, tweetsDeUsuario...)
	}
	sort.Slice(allTweets, func(i, j int) bool {
		return allTweets[i].Date.Before(*allTweets[j].Date)
	})
	return allTweets
}

// GetTweetByID obtener tweet por id
func (t *TweetManager) GetTweetByID(id int) *domain.Tweet {
	tweetBuscado := findBy(t.GetTweets(), func(tweet *domain.Tweet) bool {
		return tweet.ID == id
	})
	if len(tweetBuscado) > 0 {
		return tweetBuscado[0]
	}
	return nil
}

// GetTweetsByUser filtrar tweets de usuarios
func (t *TweetManager) GetTweetsByUser(user string) (ret []*domain.Tweet) {
	return t.TweetsMap[user]
}

// CountTweetsByUser contar tweets por username
func (t *TweetManager) CountTweetsByUser(user string) int {
	return len(t.GetTweetsByUser(user))
}

func findBy(tweets []*domain.Tweet, criteria func(*domain.Tweet) bool) (ret []*domain.Tweet) {
	for _, tweet := range tweets {
		if criteria(tweet) {
			ret = append(ret, tweet)
		}
	}
	return
}
