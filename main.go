package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/genai"
)

func main() {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GOOGLE_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Interactive loop to get user prompts
	var lastPrompt string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter prompt (ask from Gemini-2.0-flash or type 'exit' to quit): ")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		prompt := scanner.Text()
		lastPrompt = strings.ToLower(strings.TrimSpace(prompt))
		if lastPrompt == "exit" {
			break
		}
		result, err := client.Models.GenerateContent(
			ctx,
			"gemini-2.0-flash",
			genai.Text(lastPrompt),
			nil,
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.Text())

		// Ready for the next prompt
		fmt.Println("Ready for next prompt.")
	}
}
