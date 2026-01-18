package php_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type PHPGenerator struct{}

func NewPHPGenerator() generator.Generator {
	return &PHPGenerator{}
}

func (g *PHPGenerator) Name() string {
	return "php"
}

func (g *PHPGenerator) Supports(lang string) bool {
	return strings.ToLower(lang) == "php"
}

func (g *PHPGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString("<?php\n")
	sb.WriteString(fmt.Sprintf("class %s\n{\n", s.Name))
	for _, f := range s.Fields {
		sb.WriteString(fmt.Sprintf("    public $%s;\n", lowerFirst(f.Name)))
	}
	sb.WriteString("}\n")
	return sb.String(), nil
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
