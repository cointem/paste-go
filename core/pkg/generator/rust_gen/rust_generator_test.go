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
			{Name: "IsActive", OriginalName: "isActive", Kind: schema.KindBool},
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
	if !strings.Contains(code, "pub isactive: bool,") { // toLower logic in generator
		t.Error("Missing field")
	}
}
