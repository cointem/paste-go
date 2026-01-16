package java_gen

import (
	"paste-go/pkg/schema"
	"strings"
	"testing"
)

func TestJavaGenerator_Generate(t *testing.T) {
	gen := NewJavaGenerator()
	s := &schema.Struct{
		Name: "TestClass",
		Fields: []schema.Field{
			{Name: "UserName", OriginalName: "user_name", Kind: schema.KindString},
		},
	}

	code, err := gen.Generate(s)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}

	if !strings.Contains(code, "public class TestClass {") {
		t.Error("Missing class definition")
	}
	if !strings.Contains(code, "@Data") {
		t.Error("Missing Lombok annotation")
	}
	// The basic implementation returns raw OriginalName or simple camelCase logic
	// In the real impl, we expect `private String user_name;` based on current logic
	if !strings.Contains(code, "private String user_name;") {
		t.Error("Missing field")
	}
}
