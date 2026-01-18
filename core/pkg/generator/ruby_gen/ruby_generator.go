package ruby_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type RubyGenerator struct{}

func NewRubyGenerator() generator.Generator {
	return &RubyGenerator{}
}

func (g *RubyGenerator) Name() string {
	return "ruby"
}

func (g *RubyGenerator) Supports(lang string) bool {
	return strings.ToLower(lang) == "ruby"
}

func (g *RubyGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("class %s\n", s.Name))
	if len(s.Fields) > 0 {
		fields := make([]string, 0, len(s.Fields))
		for _, f := range s.Fields {
			fields = append(fields, lowerFirst(f.Name))
		}
		sb.WriteString(fmt.Sprintf("  attr_accessor :%s\n", strings.Join(fields, ", :")))
	}
	sb.WriteString("end\n")
	return sb.String(), nil
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
