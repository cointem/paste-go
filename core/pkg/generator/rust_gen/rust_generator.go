package rust_gen

import (
	"fmt"
	"paste-forge/pkg/generator"
	"paste-forge/pkg/schema"
	"strings"
)

type RustGenerator struct{}

func NewRustGenerator() generator.Generator {
	return &RustGenerator{}
}

func (g *RustGenerator) Name() string {
	return "rust"
}

func (g *RustGenerator) Supports(lang string) bool {
	l := strings.ToLower(lang)
	return l == "rust" || l == "rs"
}

func (g *RustGenerator) Generate(s *schema.Struct) (string, error) {
	sb := strings.Builder{}
	sb.WriteString("#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]\n")
	sb.WriteString("#[serde(rename_all = \"camelCase\")]\n")
	sb.WriteString(fmt.Sprintf("pub struct %s {\n", s.Name))

	for _, f := range s.Fields {
		rustType := "String" // Safe default
		switch f.Kind {
		case schema.KindString:
			rustType = "String"
		case schema.KindInt:
			rustType = "i64"
		case schema.KindFloat:
			rustType = "f64"
		case schema.KindBool:
			rustType = "bool"
		case schema.KindArray:
			rustType = "Vec<serde_json::Value>"
		case schema.KindObject:
			rustType = "serde_json::Value"
		}
		
		// snake_case
		name := strings.ToLower(f.OriginalName)
		sb.WriteString(fmt.Sprintf("\tpub %s: %s,\n", name, rustType))
	}
	sb.WriteString("}")
	return sb.String(), nil
}
