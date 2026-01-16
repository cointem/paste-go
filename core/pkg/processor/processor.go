package processor

import (
	"context"
	"fmt"
	"os"

	"paste-go/pkg/ai"
	"paste-go/pkg/generator"
	"paste-go/pkg/parser"
)

// ForgeProcessor orchestrates the conversion process.
type ForgeProcessor struct {
	parsers    []parser.Parser
	generators []generator.Generator
	aiProvider ai.Provider
}

// NewProcessor creates a new processor instance.
func NewProcessor() *ForgeProcessor {
	return &ForgeProcessor{
		parsers:    make([]parser.Parser, 0),
		generators: make([]generator.Generator, 0),
	}
}

// RegisterParser adds a new parser to the registry.
func (p *ForgeProcessor) RegisterParser(pr parser.Parser) {
	p.parsers = append(p.parsers, pr)
}

// RegisterGenerator adds a new generator to the registry.
func (p *ForgeProcessor) RegisterGenerator(g generator.Generator) {
	p.generators = append(p.generators, g)
}

// SetAIProvider sets the fallback AI provider.
func (p *ForgeProcessor) SetAIProvider(provider ai.Provider) {
	p.aiProvider = provider
}

// Process attempts to convert the code locally using Parser -> Schema -> Generator pipeline, falling back to AI if needed.
func (p *ForgeProcessor) Process(ctx context.Context, content string, targetLang string, modelName string) (string, error) {
	// 1. Try local pipeline first
	var matchedParser parser.Parser
	for _, pr := range p.parsers {
		if pr.CanParse(content) {
			matchedParser = pr
			break
		}
	}

	if matchedParser != nil {
		fmt.Fprintf(os.Stderr, "Matched parser: %s\n", matchedParser.Name())
		ir, err := matchedParser.Parse(content)
		if err == nil {
			// Find suitable generator
			for _, g := range p.generators {
				if g.Supports(targetLang) {
					result, err := g.Generate(ir)
					if err == nil {
						return result, nil
					}
					fmt.Fprintf(os.Stderr, "Generator %s failed: %v\n", g.Name(), err)
				}
			}
			fmt.Fprintf(os.Stderr, "No generator found for lang: %s\n", targetLang)
		} else {
			fmt.Fprintf(os.Stderr, "Parser %s failed: %v\n", matchedParser.Name(), err)
		}
	}

	// 2. Fallback to AI
	if p.aiProvider != nil {
		fmt.Fprintln(os.Stderr, "Fallback to AI...")
		// Construct a generic prompt
		prompt := fmt.Sprintf("Convert the following code/text into a %s struct/class definition. Return only the code.\n\n%s", targetLang, content)

		// Use provided model name, or let provider decide
		if modelName == "" {
			modelName = "default-model"
		}

		return p.aiProvider.GenerateCode(ctx, prompt, modelName)
	}

	return "", fmt.Errorf("no suitable converter found and no AI provider configured")
}
