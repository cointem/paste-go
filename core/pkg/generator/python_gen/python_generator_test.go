package python_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestPythonGenerator_Generate(t *testing.T) {
	gen := NewPythonGenerator()
	s := &schema.Struct{
		Name: "TestClass",
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

	if !strings.Contains(code, "class TestClass(BaseModel):") {
		t.Error("Missing class definition")
	}
	if !strings.Contains(code, "name: str") {
		t.Error("Missing name field")
	}
	if !strings.Contains(code, "age: int") {
		t.Error("Missing age field")
	}
	if !strings.Contains(code, "score: float") {
		t.Error("Missing score field")
	}
	if !strings.Contains(code, "active: bool") {
		t.Error("Missing active field")
	}
	if !strings.Contains(code, "tags: List[Any]") {
		t.Error("Missing tags field")
	}
	if !strings.Contains(code, "meta: dict") {
		t.Error("Missing meta field")
	}
	if !strings.Contains(code, "created_at: str") {
		t.Error("Missing created_at field")
	}
	if !strings.Contains(code, "any_value: Any") {
		t.Error("Missing any_value field")
	}
}
