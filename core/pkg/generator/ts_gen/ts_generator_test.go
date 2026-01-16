package ts_gen

import (
	"paste-go/pkg/schema"
	"strings"
	"testing"
)

func TestTSGenerator_Generate(t *testing.T) {
	gen := NewTSGenerator()
	s := &schema.Struct{
		Name: "TestInterface",
		Fields: []schema.Field{
			{Name: "Name", OriginalName: "name", Kind: schema.KindString},
			{Name: "Count", OriginalName: "count", Kind: schema.KindInt},
		},
	}

	code, err := gen.Generate(s)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}

	if !strings.Contains(code, "export interface TestInterface {") {
		t.Error("Missing interface definition")
	}
	if !strings.Contains(code, "name: string;") {
		t.Error("Missing name field")
	}
	if !strings.Contains(code, "count: number;") {
		t.Error("Missing count field")
	}
}
