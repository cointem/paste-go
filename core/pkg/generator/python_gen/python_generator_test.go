package python_gen

import (
	"paste-go/pkg/schema"
	"strings"
	"testing"
)

func TestPythonGenerator_Generate(t *testing.T) {
	gen := NewPythonGenerator()
	s := &schema.Struct{
		Name: "TestClass",
		Fields: []schema.Field{
			{Name: "MyVal", OriginalName: "my_val", Kind: schema.KindString},
		},
	}

	code, err := gen.Generate(s)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}

	if !strings.Contains(code, "class TestClass(BaseModel):") {
		t.Error("Missing class definition")
	}
	if !strings.Contains(code, "my_val: str") {
		t.Error("Missing field")
	}
}
