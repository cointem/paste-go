package csharp_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestCSharpGenerator(t *testing.T) {
	g := NewCSharpGenerator()
	s := &schema.Struct{
		Name: "User",
		Fields: []schema.Field{
			{Name: "Name", Kind: schema.KindString},
			{Name: "Age", Kind: schema.KindInt},
			{Name: "Score", Kind: schema.KindFloat},
			{Name: "Active", Kind: schema.KindBool},
			{Name: "Tags", Kind: schema.KindArray},
			{Name: "Meta", Kind: schema.KindObject},
			{Name: "CreatedAt", Kind: schema.KindTime},
			{Name: "AnyValue", Kind: schema.KindAny},
		},
	}

	code, err := g.Generate(s)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(code, "public class User") {
		t.Fatalf("missing class declaration")
	}
	if !strings.Contains(code, "public string Name") {
		t.Fatalf("missing Name field")
	}
	if !strings.Contains(code, "public int Age") {
		t.Fatalf("missing Age field")
	}
	if !strings.Contains(code, "public double Score") {
		t.Fatalf("missing Score field")
	}
	if !strings.Contains(code, "public bool Active") {
		t.Fatalf("missing Active field")
	}
	if !strings.Contains(code, "public object Tags") {
		t.Fatalf("missing Tags field")
	}
	if !strings.Contains(code, "public object Meta") {
		t.Fatalf("missing Meta field")
	}
	if !strings.Contains(code, "public DateTime CreatedAt") {
		t.Fatalf("missing CreatedAt field")
	}
	if !strings.Contains(code, "public object AnyValue") {
		t.Fatalf("missing AnyValue field")
	}
}
