package gemini

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type GeminiProvider struct {
	client *genai.Client
}

func NewGeminiProvider() *GeminiProvider {
	return &GeminiProvider{}
}

func (p *GeminiProvider) Name() string {
	return "gemini"
}

func (p *GeminiProvider) Configure(config map[string]string) error {
	apiKey, ok := config["api_key"]
	if !ok || apiKey == "" {
		return fmt.Errorf("api_key is required")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return err
	}
	p.client = client
	return nil
}

func (p *GeminiProvider) GenerateCode(ctx context.Context, prompt string, modelName string) (string, error) {
	if p.client == nil {
		return "", fmt.Errorf("gemini provider not configured")
	}
	if modelName == "default-model" {
		modelName = "gemini-2.5-flash"
	}

	model := p.client.GenerativeModel(modelName)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil {
		return "", fmt.Errorf("no response from gemini")
	}

	// Quick extraction of text
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			result := string(txt)
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
			return result, nil
		}
	}

	return "", fmt.Errorf("empty response")
}
