package main

import (
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/francoagarcia/meli-go/src/domain"
	"github.com/francoagarcia/meli-go/src/service"
)

func main() {

	shell := ishell.New()
	shell.SetPrompt("Tweeter >> ")
	shell.Print("Type 'help' to know commands\n")

	shell.AddCmd(&ishell.Cmd{
		Name: "publishTweet",
		Help: "Publishes a tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")
			user := c.ReadLine()

			c.Print("Type your tweet: ")
			text := c.ReadLine()

			tweetManager := service.NewTweetManager()
			tweet := domain.NewTweet(user, text)

			id, err := tweetManager.PublishTweet(tweet)

			if err == nil {
				c.Printf("Tweet sent with id: %v\n", id)
			} else {
				c.Print("Error publishing tweet:", err)
			}

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweet",
		Help: "Shows the last tweet",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweetManager := service.NewTweetManager()
			tweet := tweetManager.GetTweet()

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweets",
		Help: "Shows all the tweets",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			tweetManager := service.NewTweetManager()
			tweets := tweetManager.GetTweets()

			c.Println(tweets)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "showTweetById",
		Help: "Shows the tweet with the provided id",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type the id: ")
			id, _ := strconv.Atoi(c.ReadLine())
			tweetManager := service.NewTweetManager()
			tweet := tweetManager.GetTweetByID(id)

			c.Println(tweet)

			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "register",
		Help: "Register a user",
		Func: func(c *ishell.Context) {

			defer c.ShowPrompt(true)

			c.Print("Type your username: ")
			username := c.ReadLine()

			c.Print("Type your nombre: ")
			nombre := c.ReadLine()

			c.Print("Type your mail: ")
			mail := c.ReadLine()

			c.Print("Type your contrasenia: ")
			contrasenia := c.ReadLine()

			usuario := domain.NewUsuario(username, nombre, mail, contrasenia)
			usuarioManager := service.NewUsuarioManager()
			err := usuarioManager.RegistrarUsuario(usuario)

			if err == nil {
				c.Printf("Usuario registrado: %v\n", username)
			} else {
				c.Print("Error registrando usuario:", err)
			}

			return
		},
	})

	shell.Run()

}
