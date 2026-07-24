package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

func main() {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("ANTHROPIC_API_KEY environment variable not set")
	}

	llm, err := anthropic.New(
		anthropic.WithModel("claude-haiku-4.5-20241022"),
		anthropic.WithToken(apiKey),
	)
	if err != nil {
		log.Fatal(err)
	}

	tools := []llms.Tool{
		{
			Type: "function",
			Function: &llms.FunctionDefinition{
				Name:        "add",
				Description: "Add two numbers",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"a": map[string]any{"type": "integer"},
						"b": map[string]any{"type": "integer"},
					},
					"required": []string{"a", "b"},
				},
			},
		},
		{
			Type: "function",
			Function: &llms.FunctionDefinition{
				Name:        "multiply",
				Description: "Multiply two numbers",
				Parameters: map[string]any{
					"type": "object",
					"properties": map[string]any{
						"a": map[string]any{"type": "integer"},
						"b": map[string]any{"type": "integer"},
					},
					"required": []string{"a", "b"},
				},
			},
		},
	}

	ctx := context.Background()
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, "What is 25 * 4 + 10?"),
	}

	fmt.Println("Starting agent loop...")
	fmt.Printf("User: %s\n\n", messages[0].Parts[0].(llms.TextContent).Text)

	maxIterations := 10
	for i := 0; i < maxIterations; i++ {
		response, err := llm.GenerateContent(ctx, messages, llms.WithTools(tools))
		if err != nil {
			log.Fatal(err)
		}

		if len(response.Choices) == 0 {
			log.Fatal("no response from LLM")
		}

		choice := response.Choices[0]

		// Build AI response message with text + tool calls as parts
		aiParts := []llms.ContentPart{llms.TextContent{Text: choice.Content}}
		for _, tc := range choice.ToolCalls {
			aiParts = append(aiParts, tc)
		}
		messages = append(messages, llms.MessageContent{
			Role:  llms.ChatMessageTypeAI,
			Parts: aiParts,
		})

		if len(choice.ToolCalls) == 0 {
			fmt.Printf("\nFinal Answer: %s\n", choice.Content)
			return
		}

		for _, tc := range choice.ToolCalls {
			var args map[string]any
			if err := json.Unmarshal([]byte(tc.FunctionCall.Arguments), &args); err != nil {
				log.Fatalf("failed to parse args: %v", err)
			}

			a := int(args["a"].(float64))
			b := int(args["b"].(float64))

			var result int
			switch tc.FunctionCall.Name {
			case "add":
				result = a + b
			case "multiply":
				result = a * b
			default:
				log.Fatalf("unknown tool: %s", tc.FunctionCall.Name)
			}

			fmt.Printf("Tool call: %s(%v) = %d\n", tc.FunctionCall.Name, args, result)

			messages = append(messages, llms.MessageContent{
				Role: llms.ChatMessageTypeTool,
				Parts: []llms.ContentPart{
					llms.ToolCallResponse{
						ToolCallID: tc.ID,
						Name:       tc.FunctionCall.Name,
						Content:    fmt.Sprintf("%d", result),
					},
				},
			})
		}
	}

	fmt.Printf("\nReached max iterations (%d)\n", maxIterations)
}
