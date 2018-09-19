package endpoints

import (
	"strconv"

	"github.com/abiosoft/ishell"
)

// Cmd implements Endpoint
type Cmd struct {
	manager *Manager
}

// NewCmd instanciar nuevo cmd endpoint
func NewCmd(manager *Manager) *Cmd {
	return &Cmd{manager: manager}
}

// Execute levanta front con linea de comandos
func (cmd *Cmd) Execute() {
	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: cmd.publishTweet,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishImageTweet",
		Help: "Publishes a tweet with an image",
		Func: cmd.publishImageTweet,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "publishQuoteTweet",
		Help: "Publishes a tweet with a quote",
		Func: cmd.publishQuotedTweet,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows the last tweet",
		Func: cmd.showTweet,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: cmd.showTweets,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows the tweet with the provided id",
		Func: cmd.showTweetByID,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "countTweetsByUser",
		Help: "Counts the tweets published by the user",
		Func: cmd.countTweetsByUser,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetsByUser",
		Help: "Shows the tweets published by the user",
		Func: cmd.showTweetsByUser,
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Register a user",
		Func: cmd.register,
	})

	shell.Run()
}

func (cmd *Cmd) publishTweet(c *ishell.Context) {
	defer c.ShowPrompt(true)

	c.Print("Type your username: ")
	user := c.ReadLine()

	c.Print("Type your tweet: ")
	text := c.ReadLine()

	id, err := cmd.manager.PublishTextTweet(user, text)

	if err == nil {
		c.Printf("Tweet sent with id: %v\n", id)
	} else {
		c.Print("Error publishing tweet:", err)
	}

	return
}

func (cmd *Cmd) publishImageTweet(c *ishell.Context) {

	defer c.ShowPrompt(true)

	c.Print("Type your username: ")
	user := c.ReadLine()

	c.Print("Type your tweet: ")
	text := c.ReadLine()

	c.Print("Type the url of your image: ")
	url := c.ReadLine()

	id, err := cmd.manager.PublishImageTweet(user, text, url)

	if err == nil {
		c.Printf("Tweet sent with id: %v\n", id)
	} else {
		c.Print("Error publishing tweet:", err)
	}

	return
}

func (cmd *Cmd) publishQuotedTweet(c *ishell.Context) {

	defer c.ShowPrompt(true)

	c.Print("Type your username: ")
	user := c.ReadLine()

	c.Print("Type your tweet: ")
	text := c.ReadLine()

	c.Print("Type the id of the tweet you want to quote: ")
	id, _ := strconv.Atoi(c.ReadLine())

	id, err := cmd.manager.PublishQuotedTweet(user, text, id)

	if err == nil {
		c.Printf("Tweet sent with id: %v\n", id)
	} else {
		c.Print("Error publishing tweet:", err)
	}

	return
}

func (cmd *Cmd) showTweet(c *ishell.Context) {

	defer c.ShowPrompt(true)

	tweet := cmd.manager.GetTweet()

	c.Println(tweet)

	return
}

func (cmd *Cmd) showTweets(c *ishell.Context) {

	defer c.ShowPrompt(true)

	tweets := cmd.manager.GetTweets()

	c.Println(tweets)

	return
}

func (cmd *Cmd) showTweetByID(c *ishell.Context) {

	defer c.ShowPrompt(true)

	c.Print("Type the id: ")

	id, _ := strconv.Atoi(c.ReadLine())

	tweet := cmd.manager.GetTweetByID(id)

	c.Println(tweet)

	return
}

func (cmd *Cmd) countTweetsByUser(c *ishell.Context) {

	defer c.ShowPrompt(true)

	c.Print("Type the user: ")

	user := c.ReadLine()

	count := cmd.manager.CountTweetsByUser(user)

	c.Println(count)

	return
}

func (cmd *Cmd) showTweetsByUser(c *ishell.Context) {

	defer c.ShowPrompt(true)

	c.Print("Type the user: ")

	user := c.ReadLine()

	tweets := cmd.manager.GetTweetsByUser(user)

	c.Println(tweets)

	return
}

func (cmd *Cmd) register(c *ishell.Context) {

	defer c.ShowPrompt(true)

	c.Print("Type your username: ")
	username := c.ReadLine()

	c.Print("Type your nombre: ")
	nombre := c.ReadLine()

	c.Print("Type your mail: ")
	mail := c.ReadLine()

	c.Print("Type your contrasenia: ")
	contrasenia := c.ReadLine()

	_, err := cmd.manager.RegistrarUsuario(username, nombre, mail, contrasenia)

	if err == nil {
		c.Printf("Usuario registrado: %v\n", username)
	} else {
		c.Print("Error registrando usuario:", err)
	}

	return
}
