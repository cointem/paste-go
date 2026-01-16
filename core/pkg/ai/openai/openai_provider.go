package openai

import (
	"context"
	"fmt"
	"strings"

	"paste-go/pkg/ai"

	gopenai "github.com/sashabaranov/go-openai"
)

// init registers the OpenAI provider with the central factory.
func init() {
	ai.Register("openai", func() ai.Provider {
		return NewOpenAIProvider()
	})
}

type OpenAIProvider struct {
	client *gopenai.Client
}

func NewOpenAIProvider() *OpenAIProvider {
	return &OpenAIProvider{}
}

func (p *OpenAIProvider) Name() string {
	return "openai"
}

func (p *OpenAIProvider) Configure(config map[string]string) error {
	apiKey, ok := config["api_key"]
	if !ok || apiKey == "" {
		return fmt.Errorf("api_key is required for openai")
	}

	conf := gopenai.DefaultConfig(apiKey)
	if baseURL, ok := config["base_url"]; ok && baseURL != "" {
		conf.BaseURL = baseURL
	}

	p.client = gopenai.NewClientWithConfig(conf)
	return nil
}

func (p *OpenAIProvider) GenerateCode(ctx context.Context, prompt string, modelName string) (string, error) {
	if p.client == nil {
		return "", fmt.Errorf("openai provider not configured")
	}
	// Default model logic
	if modelName == "default-model" || modelName == "" {
		modelName = gopenai.GPT3Dot5Turbo
	}

	resp, err := p.client.CreateChatCompletion(
		ctx,
		gopenai.ChatCompletionRequest{
			Model: modelName,
			Messages: []gopenai.ChatCompletionMessage{
				{
					Role:    gopenai.ChatMessageRoleSystem,
					Content: "You are a coding assistant. Output only the requested code. No markdown, no explanations.",
				},
				{
					Role:    gopenai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from openai")
	}

	result := resp.Choices[0].Message.Content
	return cleanCodeBlock(result), nil
}

func cleanCodeBlock(content string) string {
	result := content
	// Strip Markdown code blocks if present
	if strings.HasPrefix(result, "```") {
		lines := strings.Split(result, "\n")
		// Remove first line if it starts with ```
		if len(lines) > 0 && strings.HasPrefix(strings.TrimSpace(lines[0]), "```") {
			lines = lines[1:]
		}
		// Remove last line if it is ```
		if len(lines) > 0 {
			lastIdx := len(lines) - 1
			if strings.TrimSpace(lines[lastIdx]) == "```" {
				lines = lines[:lastIdx]
			}
		}
		result = strings.Join(lines, "\n")
	}
	return result
}
