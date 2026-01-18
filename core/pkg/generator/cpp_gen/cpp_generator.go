package cpp_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type CppGenerator struct{}

func NewCppGenerator() generator.Generator {
	return &CppGenerator{}
}

func (g *CppGenerator) Name() string {
	return "cpp"
}

func (g *CppGenerator) Supports(lang string) bool {
	l := strings.ToLower(lang)
	return l == "cpp" || l == "c++" || l == "c"
}

func (g *CppGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("struct %s {\n", s.Name))
	for _, f := range s.Fields {
		cppType := "auto"
		switch f.Kind {
		case schema.KindString:
			cppType = "std::string"
		case schema.KindInt:
			cppType = "int"
		case schema.KindFloat:
			cppType = "double"
		case schema.KindBool:
			cppType = "bool"
		case schema.KindArray:
			cppType = "std::vector<std::string>"
		case schema.KindObject:
			cppType = "std::map<std::string, std::string>"
		}
		sb.WriteString(fmt.Sprintf("    %s %s;\n", cppType, lowerFirst(f.Name)))
	}
	sb.WriteString("};\n")
	return sb.String(), nil
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
