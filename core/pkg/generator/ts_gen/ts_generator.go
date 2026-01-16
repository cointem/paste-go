package ts_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type TSGenerator struct{}

func NewTSGenerator() generator.Generator {
	return &TSGenerator{}
}

func (g *TSGenerator) Name() string {
	return "typescript"
}

func (g *TSGenerator) Supports(lang string) bool {
	l := strings.ToLower(lang)
	return l == "typescript" || l == "ts"
}

func (g *TSGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("export interface %s {\n", s.Name))

	for _, f := range s.Fields {
		tsType := "any"
		switch f.Kind {
		case schema.KindString:
			tsType = "string"
		case schema.KindInt, schema.KindFloat:
			tsType = "number"
		case schema.KindBool:
			tsType = "boolean"
		case schema.KindTime:
			tsType = "Date"
		case schema.KindObject:
			tsType = "object"
		case schema.KindArray:
			tsType = "any[]"
		}

		sb.WriteString(fmt.Sprintf("\t%s: %s;\n", f.OriginalName, tsType))
	}
	sb.WriteString("}")
	return sb.String(), nil
}
