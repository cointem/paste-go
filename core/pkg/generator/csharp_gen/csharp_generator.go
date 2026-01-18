package csharp_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type CSharpGenerator struct{}

func NewCSharpGenerator() generator.Generator {
	return &CSharpGenerator{}
}

func (g *CSharpGenerator) Name() string {
	return "csharp"
}

func (g *CSharpGenerator) Supports(lang string) bool {
	l := strings.ToLower(lang)
	return l == "c#" || l == "csharp" || l == "cs"
}

func (g *CSharpGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("public class %s\n{\n", s.Name))

	for _, f := range s.Fields {
		csType := "object"
		switch f.Kind {
		case schema.KindString:
			csType = "string"
		case schema.KindInt:
			csType = "int"
		case schema.KindFloat:
			csType = "double"
		case schema.KindBool:
			csType = "bool"
		case schema.KindTime:
			csType = "DateTime"
		}
		sb.WriteString(fmt.Sprintf("    public %s %s { get; set; }\n", csType, f.Name))
	}
	sb.WriteString("}\n")
	return sb.String(), nil
}
