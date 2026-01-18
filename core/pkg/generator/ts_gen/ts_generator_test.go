package ts_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestTSGenerator_Generate(t *testing.T) {
	gen := NewTSGenerator()
	s := &schema.Struct{
		Name: "TestInterface",
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

	if !strings.Contains(code, "export interface TestInterface {") {
		t.Error("Missing interface definition")
	}
	if !strings.Contains(code, "name: string;") {
		t.Error("Missing name field")
	}
	if !strings.Contains(code, "age: number;") {
		t.Error("Missing age field")
	}
	if !strings.Contains(code, "score: number;") {
		t.Error("Missing score field")
	}
	if !strings.Contains(code, "active: boolean;") {
		t.Error("Missing active field")
	}
	if !strings.Contains(code, "tags: any[];") {
		t.Error("Missing tags field")
	}
	if !strings.Contains(code, "meta: object;") {
		t.Error("Missing meta field")
	}
	if !strings.Contains(code, "created_at: Date;") {
		t.Error("Missing created_at field")
	}
	if !strings.Contains(code, "any_value: any;") {
		t.Error("Missing any_value field")
	}
}
