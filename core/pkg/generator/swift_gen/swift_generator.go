package swift_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type SwiftGenerator struct{}

func NewSwiftGenerator() generator.Generator {
	return &SwiftGenerator{}
}

func (g *SwiftGenerator) Name() string {
	return "swift"
}

func (g *SwiftGenerator) Supports(lang string) bool {
	return strings.ToLower(lang) == "swift"
}

func (g *SwiftGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("struct %s: Codable {\n", s.Name))
	for _, f := range s.Fields {
		swiftType := "Any"
		switch f.Kind {
		case schema.KindString:
			swiftType = "String"
		case schema.KindInt:
			swiftType = "Int"
		case schema.KindFloat:
			swiftType = "Double"
		case schema.KindBool:
			swiftType = "Bool"
		case schema.KindArray:
			swiftType = "[Any]"
		}
		sb.WriteString(fmt.Sprintf("    var %s: %s\n", lowerFirst(f.Name), swiftType))
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
