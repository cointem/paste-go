package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"paste-forge/pkg/ai"
	"paste-forge/pkg/ai/gemini"
	"paste-forge/pkg/ai/openai"
	gitkraken "paste-forge/pkg/generator/go_gen"
	javaGen "paste-forge/pkg/generator/java_gen"
	pyGen "paste-forge/pkg/generator/python_gen"
	rustGen "paste-forge/pkg/generator/rust_gen"
	tsGen "paste-forge/pkg/generator/ts_gen"
	jsonParser "paste-forge/pkg/parser/json"
	sqlParser "paste-forge/pkg/parser/sql"
	xmlParser "paste-forge/pkg/parser/xml"
	"paste-forge/pkg/processor"
)

func main() {
	targetLang := flag.String("lang", "go", "Target language for conversion")
	apiKey := flag.String("key", "", "API Key for AI Fallback")
	aiProviderType := flag.String("provider", "gemini", "AI Provider (gemini, openai)")
	aiModel := flag.String("model", "", "Model Name")
	aiBaseURL := flag.String("baseurl", "", "Custom Base URL for AI")
	flag.Parse()

	// 1. Initialize Processor
	proc := processor.NewProcessor()

	// 2. Register Parsers
	proc.RegisterParser(jsonParser.NewJSONParser())
	proc.RegisterParser(sqlParser.NewSQLParser())
	proc.RegisterParser(xmlParser.NewXMLParser())

	// 3. Register Generators
	proc.RegisterGenerator(gitkraken.NewGoGenerator())
	proc.RegisterGenerator(tsGen.NewTSGenerator())
	proc.RegisterGenerator(pyGen.NewPythonGenerator())
	proc.RegisterGenerator(javaGen.NewJavaGenerator())
	proc.RegisterGenerator(rustGen.NewRustGenerator())

	// 4. Configure AI (if key provided)
	if *apiKey != "" {
		var provider ai.Provider
		
		switch strings.ToLower(*aiProviderType) {
		case "openai":
			provider = openai.NewOpenAIProvider()
		case "gemini":
			// Default to gemini
			provider = gemini.NewGeminiProvider()
		default:
			// Default to gemini if unknown, or handle error
			fmt.Fprintf(os.Stderr, "Unknown AI provider: %s, defaulting to Gemini\n", *aiProviderType)
			provider = gemini.NewGeminiProvider()
		}

		config := map[string]string{
			"api_key": *apiKey,
		}
		if *aiBaseURL != "" {
			config["base_url"] = *aiBaseURL
		}

		err := provider.Configure(config)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to configure AI: %v\n", err)
		} else {
			proc.SetAIProvider(provider)
		}
	}

	// 4. Read Input from Stdin
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
		os.Exit(1)
	}
	content := string(inputBytes)

	// 6. Process
	result, err := proc.Process(context.Background(), content, *targetLang, *aiModel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1) // Exit code 1 for failure
	}

	// 6. Output Result
	fmt.Print(result)
}
