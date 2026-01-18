package go_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type GoGenerator struct{}

func NewGoGenerator() generator.Generator {
	return &GoGenerator{}
}

func (g *GoGenerator) Name() string {
	return "go"
}

func (g *GoGenerator) Supports(lang string) bool {
	return strings.ToLower(lang) == "go"
}

func (g *GoGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("type %s struct {\n", s.Name))

	for _, f := range s.Fields {
		goType := goTypeForField(f, 1)
		sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", f.Name, goType, f.OriginalName))
	}
	sb.WriteString("}")
	return sb.String(), nil
}

func goTypeForField(f schema.Field, indentLevel int) string {
	baseIndent := strings.Repeat("\t", indentLevel)

	switch f.Kind {
	case schema.KindString:
		return "string"
	case schema.KindInt:
		return "int64"
	case schema.KindFloat:
		return "float64"
	case schema.KindBool:
		return "bool"
	case schema.KindTime:
		return "time.Time"
	case schema.KindObject:
		return inlineStructType(f.Nested, indentLevel, baseIndent)
	case schema.KindArray:
		if f.Nested != nil {
			return "[]" + inlineStructType(f.Nested, indentLevel, baseIndent)
		}
		return "[]interface{}"
	default:
		return "interface{}"
	}
}

func inlineStructType(nested *schema.Struct, indentLevel int, baseIndent string) string {
	if nested == nil || len(nested.Fields) == 0 {
		return "struct {}"
	}
	var sb strings.Builder
	sb.WriteString("struct {\n")
	for _, nf := range nested.Fields {
		nestedType := goTypeForField(nf, indentLevel+1)
		sb.WriteString(fmt.Sprintf("%s%s %s `json:\"%s\"`\n", baseIndent+"\t", nf.Name, nestedType, nf.OriginalName))
	}
	sb.WriteString(baseIndent + "}")
	return sb.String()
}
