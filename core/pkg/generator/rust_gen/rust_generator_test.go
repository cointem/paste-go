package rust_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestRustGenerator_Generate(t *testing.T) {
	gen := NewRustGenerator()
	s := &schema.Struct{
		Name: "TestStruct",
		Fields: []schema.Field{
			{Name: "Name", OriginalName: "name", Kind: schema.KindString},
			{Name: "Age", OriginalName: "age", Kind: schema.KindInt},
			{Name: "Score", OriginalName: "score", Kind: schema.KindFloat},
			{Name: "Active", OriginalName: "active", Kind: schema.KindBool},
			{Name: "Tags", OriginalName: "tags", Kind: schema.KindArray},
			{Name: "Meta", OriginalName: "meta", Kind: schema.KindObject},
			{Name: "CreatedAt", OriginalName: "created_at", Kind: schema.KindTime},
			{Name: "AnyValue", OriginalName: "any_value", Kind: schema.KindAny},
		},
	}

	code, err := gen.Generate(s)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}

	if !strings.Contains(code, "pub struct TestStruct {") {
		t.Error("Missing struct definition")
	}
	if !strings.Contains(code, "#[derive(Default, Debug, Clone, PartialEq, Serialize, Deserialize)]") {
		t.Error("Missing derive macros")
	}
	if !strings.Contains(code, "pub name: String,") {
		t.Error("Missing name field")
	}
	if !strings.Contains(code, "pub age: i64,") {
		t.Error("Missing age field")
	}
	if !strings.Contains(code, "pub score: f64,") {
		t.Error("Missing score field")
	}
	if !strings.Contains(code, "pub active: bool,") {
		t.Error("Missing active field")
	}
	if !strings.Contains(code, "pub tags: Vec<serde_json::Value>,") {
		t.Error("Missing tags field")
	}
	if !strings.Contains(code, "pub meta: serde_json::Value,") {
		t.Error("Missing meta field")
	}
	if !strings.Contains(code, "pub created_at: String,") {
		t.Error("Missing created_at field")
	}
	if !strings.Contains(code, "pub any_value: String,") {
		t.Error("Missing any_value field")
	}
}
