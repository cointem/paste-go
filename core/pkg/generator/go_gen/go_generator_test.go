package go_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestGoGenerator_Generate(t *testing.T) {
	gen := NewGoGenerator()
	s := &schema.Struct{
		Name: "TestStruct",
		Fields: []schema.Field{
			{Name: "MyString", OriginalName: "my_string", Kind: schema.KindString},
			{Name: "MyInt", OriginalName: "my_int", Kind: schema.KindInt},
		},
	}

	code, err := gen.Generate(s)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}

	if !strings.Contains(code, "type TestStruct struct {") {
		t.Error("Missing struct definition")
	}
	if !strings.Contains(code, "MyString string `json:\"my_string\"`") {
		t.Error("Missing string field")
	}
	if !strings.Contains(code, "MyInt int64 `json:\"my_int\"`") {
		t.Error("Missing int field")
	}
}
