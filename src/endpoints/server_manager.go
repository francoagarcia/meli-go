package endpoints

import (
	"fmt"

	"github.com/satori/go.uuid"

	"github.com/francoagarcia/meli-go/src/domain"
	"github.com/francoagarcia/meli-go/src/service"
)

// Endpoint interface de app
type Endpoint interface {
	Execute()
}

// Manager server manager
type Manager struct {
	tweetManager *service.TweetManager
}

// NewManager instancia nuevo server manager
func NewManager(tweetManager *service.TweetManager) *Manager {
	return &Manager{tweetManager: tweetManager}
}

// PublishTextTweet publish new text tweet
func (manager *Manager) PublishTextTweet(user, text string) (id *uuid.UUID, err error) {
	tweet := domain.NewTextTweet(user, text)
	id, err = manager.tweetManager.PublishTweet(tweet)
	return
}

// PublishImageTweet publish new image tweet
func (manager *Manager) PublishImageTweet(user, text, url string) (id *uuid.UUID, err error) {
	tweet := domain.NewImageTweet(user, text, url)
	id, err = manager.tweetManager.PublishTweet(tweet)
	return
}

// PublishQuotedTweet publish new quoted tweet
func (manager *Manager) PublishQuotedTweet(user, text string, idQuoted *uuid.UUID) (id *uuid.UUID, err error) {
	quoteTweet := manager.tweetManager.GetTweetByID(idQuoted)
	if quoteTweet == nil {
		return nil, fmt.Errorf("PublishQuotedTweet - tweet not found - id [%s]", idQuoted.String())
	}
	tweet := domain.NewQuoteTweet(user, text, quoteTweet)
	id, err = manager.tweetManager.PublishTweet(tweet)
	return
}

// GetTweet obtener ultimo tweet publicado
func (manager *Manager) GetTweet() domain.Tweet {
	return manager.tweetManager.GetTweet()
}

// GetTweets obtener ultimo tweet publicado
func (manager *Manager) GetTweets() []domain.Tweet {
	return manager.tweetManager.GetTweets()
}

// GetTweetByID obtener tweet por id
func (manager *Manager) GetTweetByID(id *uuid.UUID) domain.Tweet {
	return manager.tweetManager.GetTweetByID(id)
}

// CountTweetsByUser obtener cantidad de tweets publicados por un usuario
func (manager *Manager) CountTweetsByUser(user string) int {
	return manager.tweetManager.CountTweetsByUser(user)
}

// GetTweetsByUser obtener tweets de un usuario
func (manager *Manager) GetTweetsByUser(user string) []domain.Tweet {
	return manager.tweetManager.GetTweetsByUser(user)
}

// RegistrarUsuario registrar usuario
func (manager *Manager) RegistrarUsuario(username, nombre, mail, contrasenia string) (user *domain.Usuario, err error) {
	usuarioRequest := domain.NewUsuario(username, nombre, mail, contrasenia)
	usuarioManager := service.NewUsuarioManager()
	user, err = usuarioManager.RegistrarUsuario(usuarioRequest)
	return
}
