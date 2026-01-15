package go_gen

import (
	"fmt"
	"paste-forge/pkg/generator"
	"paste-forge/pkg/schema"
	"strings"
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
		goType := "interface{}"
		switch f.Kind {
		case schema.KindString:
			goType = "string"
		case schema.KindInt:
			goType = "int64"
		case schema.KindFloat:
			goType = "float64"
		case schema.KindBool:
			goType = "bool"
		case schema.KindTime:
			goType = "time.Time"
		case schema.KindObject:
			goType = "struct { ... }"
		case schema.KindArray:
			goType = "[]interface{}"
		}

		sb.WriteString(fmt.Sprintf("\t%s %s `json:\"%s\"`\n", f.Name, goType, f.OriginalName))
	}
	sb.WriteString("}")
	return sb.String(), nil
}
