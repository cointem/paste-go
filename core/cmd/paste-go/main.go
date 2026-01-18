package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"paste-go/pkg/ai"
	_ "paste-go/pkg/ai/gemini"
	_ "paste-go/pkg/ai/openai"
	gitkraken "paste-go/pkg/generator/go_gen"
	javaGen "paste-go/pkg/generator/java_gen"
	pyGen "paste-go/pkg/generator/python_gen"
	rustGen "paste-go/pkg/generator/rust_gen"
	tsGen "paste-go/pkg/generator/ts_gen"
	cppGen "paste-go/pkg/generator/cpp_gen"
	csGen "paste-go/pkg/generator/csharp_gen"
	dartGen "paste-go/pkg/generator/dart_gen"
	kotlinGen "paste-go/pkg/generator/kotlin_gen"
	phpGen "paste-go/pkg/generator/php_gen"
	rubyGen "paste-go/pkg/generator/ruby_gen"
	scalaGen "paste-go/pkg/generator/scala_gen"
	swiftGen "paste-go/pkg/generator/swift_gen"
	jsonParser "paste-go/pkg/parser/json"
	sqlParser "paste-go/pkg/parser/sql"
	xmlParser "paste-go/pkg/parser/xml"
	"paste-go/pkg/processor"
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
	proc.RegisterGenerator(csGen.NewCSharpGenerator())
	proc.RegisterGenerator(kotlinGen.NewKotlinGenerator())
	proc.RegisterGenerator(swiftGen.NewSwiftGenerator())
	proc.RegisterGenerator(phpGen.NewPHPGenerator())
	proc.RegisterGenerator(rubyGen.NewRubyGenerator())
	proc.RegisterGenerator(dartGen.NewDartGenerator())
	proc.RegisterGenerator(cppGen.NewCppGenerator())
	proc.RegisterGenerator(scalaGen.NewScalaGenerator())

	// 4. Configure AI (if key provided)
	if *apiKey != "" {
		var provider ai.Provider
		var err error

		// Use the factory to get the provider
		provider, err = ai.GetProvider(*aiProviderType)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: %v. Defaulting to Gemini.\n", err)
			provider, _ = ai.GetProvider("gemini")
		}

		if provider == nil {
			// If even gemini fails (should not happen if registered correctly)
			fmt.Fprintf(os.Stderr, "Error: Could not initialize any AI provider.\n")
			os.Exit(1)
		}

		config := map[string]string{
			"api_key": *apiKey,
		}
		if *aiBaseURL != "" {
			config["base_url"] = *aiBaseURL
		}

		err = provider.Configure(config)
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
