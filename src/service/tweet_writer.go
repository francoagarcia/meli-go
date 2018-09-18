package service

import (
	"fmt"

	"github.com/francoagarcia/meli-go/src/domain"
)

// TweetWriter interfaz del tweet writer
type TweetWriter interface {
	Write(tweet domain.Tweet)
}

/**************************************
 ********* Memory Tweet Writer ********
 **************************************/

// MemoryTweetWriter canal de tweet writer en memoria
type MemoryTweetWriter struct {
	Tweets []domain.Tweet
}

// NewMemoryTweetWriter crear MemoryTweetWriter
func NewMemoryTweetWriter() *MemoryTweetWriter {
	return &MemoryTweetWriter{Tweets: make([]domain.Tweet, 0)}
}

// WriteTweet escribir tweet en canal de tweets. Finaliza cuando recibe un true en el canal quit
func (writer *MemoryTweetWriter) Write(tweet domain.Tweet) {
	writer.Tweets = append(writer.Tweets, tweet)
}

/**************************************
 ********** File Tweet Writer *********
 **************************************/

// FileTweetWriter file tweet writer
type FileTweetWriter struct {
	Tweets []domain.Tweet
}

// NewFileTweetWriter crear NewFileTweetWriter
func NewFileTweetWriter() *FileTweetWriter {
	return &FileTweetWriter{Tweets: make([]domain.Tweet, 0)}
}

// WriteTweet escribir tweet en canal
func (writer *FileTweetWriter) Write(tweet domain.Tweet) {

}

/*****************************************
 ********** Channel Tweet Writer *********
 *****************************************/

// ChannelTweetWriter wrapper del TweetWriter
type ChannelTweetWriter struct {
	Writer TweetWriter
}

// NewChannelTweetWriter crear ChannelTweetWriter
func NewChannelTweetWriter(tweetWriter TweetWriter) *ChannelTweetWriter {
	return &ChannelTweetWriter{Writer: tweetWriter}
}

// WriteTweet escribir tweet en canal
func (channelWrapper *ChannelTweetWriter) WriteTweet(tweetsToWrite chan domain.Tweet, quit chan bool) {

	for tweet := range tweetsToWrite {
		fmt.Sprintln("tweet to write")
		channelWrapper.Writer.Write(tweet)
	}

	quit <- true
}
