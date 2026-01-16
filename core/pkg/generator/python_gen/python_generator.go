package python_gen

import (
	"fmt"
	"strings"

	"paste-go/pkg/generator"
	"paste-go/pkg/schema"
)

type PythonGenerator struct{}

func NewPythonGenerator() generator.Generator {
	return &PythonGenerator{}
}

func (g *PythonGenerator) Name() string {
	return "python"
}

func (g *PythonGenerator) Supports(lang string) bool {
	l := strings.ToLower(lang)
	return l == "python" || l == "py"
}

func (g *PythonGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString("from pydantic import BaseModel\n")
	sb.WriteString("from typing import Any, List, Optional\n\n")

	sb.WriteString(fmt.Sprintf("class %s(BaseModel):\n", s.Name))

	if len(s.Fields) == 0 {
		sb.WriteString("    pass\n")
		return sb.String(), nil
	}

	for _, f := range s.Fields {
		pyType := "Any"
		switch f.Kind {
		case schema.KindString:
			pyType = "str"
		case schema.KindInt:
			pyType = "int"
		case schema.KindFloat:
			pyType = "float"
		case schema.KindBool:
			pyType = "bool"
		case schema.KindTime:
			pyType = "str" // keeping simple
		case schema.KindArray:
			pyType = "List[Any]"
		case schema.KindObject:
			pyType = "dict"
		}

		// Python snake_case convention
		pyName := toSnakeCase(f.OriginalName)
		sb.WriteString(fmt.Sprintf("    %s: %s\n", pyName, pyType))
	}
	return sb.String(), nil
}

func toSnakeCase(s string) string {
	// Very simple implementation
	return strings.ToLower(s)
}
