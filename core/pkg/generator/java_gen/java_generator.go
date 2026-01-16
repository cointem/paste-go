package java_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type JavaGenerator struct{}

func NewJavaGenerator() generator.Generator {
	return &JavaGenerator{}
}

func (g *JavaGenerator) Name() string {
	return "java"
}

func (g *JavaGenerator) Supports(lang string) bool {
	l := strings.ToLower(lang)
	return l == "java"
}

func (g *JavaGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString("import lombok.Data;\n\n")
	sb.WriteString("@Data\n")
	sb.WriteString(fmt.Sprintf("public class %s {\n", s.Name))

	for _, f := range s.Fields {
		javaType := "Object"
		switch f.Kind {
		case schema.KindString:
			javaType = "String"
		case schema.KindInt:
			javaType = "Integer"
		case schema.KindFloat:
			javaType = "Double"
		case schema.KindBool:
			javaType = "Boolean"
		case schema.KindTime:
			javaType = "String"
		case schema.KindArray:
			javaType = "java.util.List<Object>"
		}

		// camelCase
		name := toCamelCase(f.OriginalName)
		sb.WriteString(fmt.Sprintf("\tprivate %s %s;\n", javaType, name))
	}
	sb.WriteString("}")
	return sb.String(), nil
}

func toCamelCase(s string) string {
	// basic implementation
	return s
}
