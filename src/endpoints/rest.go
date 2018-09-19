package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Rest implements Endpoint
type Rest struct {
	manager *Manager
}

// NewRest instanciar nuevo cmd endpoint
func NewRest(manager *Manager) *Rest {
	return &Rest{manager: manager}
}

// Execute levanta server rest
func (rest *Rest) Execute() {
	r := gin.Default()

	r.POST("/tweets/text", rest.publishTextTweet)
	r.POST("/tweets/image", rest.publishImageTweet)
	r.POST("/tweets/quoted", rest.publishQuotedTweet)
	r.GET("/tweets", rest.getTweets)
	r.GET("/tweets/:id", rest.getTweetByID)
	r.GET("/users/:username/tweets", rest.getTweetsByUser)
	r.POST("/users", rest.registrarUsuario)

	r.Run() // listen and serve on 0.0.0.0:8080
}

// TweetResource resource tweets
type TweetResource struct {
	User string `json:"user"`
	Text string `json:"text"`
	URL  string `json:"url"`
	ID   int    `json:"id"`
}

func (rest *Rest) publishTextTweet(c *gin.Context) {
	var tweet TweetResource
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := rest.manager.PublishTextTweet(tweet.User, tweet.Text)

	if err == nil {
		c.JSON(http.StatusOK, struct{ ID int }{id})
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (rest *Rest) publishImageTweet(c *gin.Context) {
	var tweet TweetResource
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := rest.manager.PublishImageTweet(tweet.User, tweet.Text, tweet.URL)

	if err == nil {
		c.JSON(http.StatusOK, struct{ ID int }{id})
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (rest *Rest) publishQuotedTweet(c *gin.Context) {
	var tweet TweetResource
	if err := c.ShouldBindJSON(&tweet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := rest.manager.PublishQuotedTweet(tweet.User, tweet.Text, tweet.ID)

	if err == nil {
		c.JSON(http.StatusOK, struct{ ID int }{id})
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func (rest *Rest) getTweets(c *gin.Context) {
	lastTweet, _ := strconv.ParseBool(c.Query("last_tweet"))
	if lastTweet {
		c.JSON(http.StatusOK, rest.manager.GetTweet())
		return
	}

	c.JSON(http.StatusOK, rest.manager.GetTweets())
	return
}

func (rest *Rest) getTweetByID(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweetBuscado := rest.manager.GetTweetByID(ID)

	c.JSON(http.StatusOK, tweetBuscado)
}

func (rest *Rest) getTweetsByUser(c *gin.Context) {
	var user = c.Param("username")

	total, _ := strconv.ParseBool(c.Query("total"))
	if total {
		c.JSON(http.StatusOK, rest.manager.CountTweetsByUser(user))
		return
	}

	c.JSON(http.StatusOK, rest.manager.GetTweetsByUser(user))
	return
}

// UserResource resource usuarios
type UserResource struct {
	Username    string `json:"username"`
	Nombre      string `json:"nombre"`
	Mail        string `json:"mail"`
	Contrasenia string `json:"contrasenia"`
}

func (rest *Rest) registrarUsuario(c *gin.Context) {
	var user UserResource
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usuarioCreado, err := rest.manager.RegistrarUsuario(user.Username, user.Nombre, user.Mail, user.Contrasenia)

	if err == nil {
		c.JSON(http.StatusOK, usuarioCreado)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
