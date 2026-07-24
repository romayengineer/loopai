package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/anthropic"
)

func init() {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		log.Fatalf("error loading .env file: %v", err)
	}
}

func main() {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Fatal("ANTHROPIC_API_KEY environment variable not set")
	}

	llm, err := anthropic.New(
		anthropic.WithModel("claude-haiku-4-5"),
		anthropic.WithToken(apiKey),
		anthropic.WithBaseURL("https://opencode.ai/zen/v1"),
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

		// Aggregate text and tool calls across all content choices
		// (Anthropic returns text + tool_use as separate choices)
		var textParts []string
		var toolCalls []llms.ToolCall
		for _, choice := range response.Choices {
			if choice.Content != "" {
				textParts = append(textParts, choice.Content)
			}
			toolCalls = append(toolCalls, choice.ToolCalls...)
		}

		if len(toolCalls) == 0 {
			answer := strings.Join(textParts, "\n")
			fmt.Printf("\nFinal Answer: %s\n", answer)
			messages = append(messages, llms.TextParts(llms.ChatMessageTypeAI, answer))
			return
		}

		if text := strings.Join(textParts, "\n"); text != "" {
			fmt.Println(text)
		}

		aiParts := make([]llms.ContentPart, len(toolCalls))
		for i, tc := range toolCalls {
			aiParts[i] = tc
		}
		messages = append(messages, llms.MessageContent{
			Role:  llms.ChatMessageTypeAI,
			Parts: aiParts,
		})

		for _, tc := range toolCalls {
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
