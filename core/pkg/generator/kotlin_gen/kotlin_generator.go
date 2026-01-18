package kotlin_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type KotlinGenerator struct{}

func NewKotlinGenerator() generator.Generator {
	return &KotlinGenerator{}
}

func (g *KotlinGenerator) Name() string {
	return "kotlin"
}

func (g *KotlinGenerator) Supports(lang string) bool {
	return strings.ToLower(lang) == "kotlin"
}

func (g *KotlinGenerator) Generate(s *schema.Struct) (string, error) {
	parts := make([]string, 0, len(s.Fields))
	for _, f := range s.Fields {
		ktType := "Any"
		switch f.Kind {
		case schema.KindString:
			ktType = "String"
		case schema.KindInt:
			ktType = "Int"
		case schema.KindFloat:
			ktType = "Double"
		case schema.KindBool:
			ktType = "Boolean"
		case schema.KindArray:
			ktType = "List<Any>"
		}
		parts = append(parts, fmt.Sprintf("val %s: %s", lowerFirst(f.Name), ktType))
	}
	return fmt.Sprintf("data class %s(%s)", s.Name, strings.Join(parts, ", ")), nil
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
