package ai

import "context"

// Provider defines the interface for AI backends.
// New AI services (e.g., Anthropic, DeepSeek) can be added by implementing this interface.
type Provider interface {
	// Name returns the provider identifier (e.g., "gemini", "openai")
	Name() string

	// Configure sets up the provider with API keys and other settings.
	Configure(config map[string]string) error

	// GenerateCode sends a prompt to the AI and returns the generated code.
	GenerateCode(ctx context.Context, prompt string, model string) (string, error)
}
