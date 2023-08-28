package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func getGPT3Response(prompt string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY")) // Replace with your OpenAI API Key

	completion, err := client.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	})

	if err != nil {
		fmt.Println("Error getting response:", err)
		return "Error"
	}

	return completion.Choices[0].Message.Content
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("You: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1] // remove newline character

		if text == "exit" {
			break
		}

		response := getGPT3Response(text)
		fmt.Println("GPT-3:", response)
	}
}
