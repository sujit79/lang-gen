package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cdipaolo/sentiment"
)

func Sentiment(text string) int {
	model, err := sentiment.Restore()
	if err != nil {
		panic(err)
	}

	analysis := model.SentimentAnalysis(text, sentiment.English)
	return int(analysis.Score)
}

func main() {

	args := os.Args

	content, error := os.ReadFile(args[1])
	if error != nil {

		log.Fatal(error)
	}

	text := string(content)
	score := Sentiment(text)
	f, _ := os.Create(args[2])
	f.WriteString(fmt.Sprintf("%d", score))
	f.Close()
}
