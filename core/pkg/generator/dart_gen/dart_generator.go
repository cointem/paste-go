package dart_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type DartGenerator struct{}

func NewDartGenerator() generator.Generator {
	return &DartGenerator{}
}

func (g *DartGenerator) Name() string {
	return "dart"
}

func (g *DartGenerator) Supports(lang string) bool {
	return strings.ToLower(lang) == "dart"
}

func (g *DartGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("class %s {\n", s.Name))
	for _, f := range s.Fields {
		dartType := "dynamic"
		switch f.Kind {
		case schema.KindString:
			dartType = "String"
		case schema.KindInt:
			dartType = "int"
		case schema.KindFloat:
			dartType = "double"
		case schema.KindBool:
			dartType = "bool"
		case schema.KindArray:
			dartType = "List<dynamic>"
		}
		sb.WriteString(fmt.Sprintf("  final %s %s;\n", dartType, lowerFirst(f.Name)))
	}

	sb.WriteString("\n  ")
	sb.WriteString(s.Name)
	sb.WriteString("({")
	for i, f := range s.Fields {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("required this.")
		sb.WriteString(lowerFirst(f.Name))
	}
	sb.WriteString("});\n")
	sb.WriteString("}\n")
	return sb.String(), nil
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
