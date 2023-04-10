package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Abeldlp/ai-go/option"
	openai "github.com/sashabaranov/go-openai"
)

func main() {
	chatGPT := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	askPrt := flag.Bool("ask", false, "Ask a single question")

	flag.Parse()

	if *askPrt {
		if len(os.Args) < 3 {
			fmt.Println("\nNo prompt passed")
			return
		}
		prompt := os.Args[2]
		option.Ask(prompt, chatGPT)
	} else {
		option.Chat(chatGPT)
	}
}
