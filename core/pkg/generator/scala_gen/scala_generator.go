package scala_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type ScalaGenerator struct{}

func NewScalaGenerator() generator.Generator {
	return &ScalaGenerator{}
}

func (g *ScalaGenerator) Name() string {
	return "scala"
}

func (g *ScalaGenerator) Supports(lang string) bool {
	return strings.ToLower(lang) == "scala"
}

func (g *ScalaGenerator) Generate(s *schema.Struct) (string, error) {
	parts := make([]string, 0, len(s.Fields))
	for _, f := range s.Fields {
		scType := "Any"
		switch f.Kind {
		case schema.KindString:
			scType = "String"
		case schema.KindInt:
			scType = "Int"
		case schema.KindFloat:
			scType = "Double"
		case schema.KindBool:
			scType = "Boolean"
		case schema.KindArray:
			scType = "List[Any]"
		case schema.KindObject:
			scType = "Map[String, Any]"
		}
		parts = append(parts, fmt.Sprintf("%s: %s", lowerFirst(f.Name), scType))
	}
	return fmt.Sprintf("case class %s(%s)", s.Name, strings.Join(parts, ", ")), nil
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
