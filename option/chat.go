package option

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func Chat(c *openai.Client) {

	ctx := context.Background()

	fmt.Println("")
	fmt.Println("To exit the program type: exit")

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{},
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("")
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')

		if text == "quit\n" || text == "exit\n" {
			break
		}

		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		resp, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		fmt.Printf("\n")
		for {
			response, err := resp.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Printf("\n")
				break
			}

			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				break
			}

			fmt.Printf(response.Choices[0].Delta.Content)
		}
	}
}
