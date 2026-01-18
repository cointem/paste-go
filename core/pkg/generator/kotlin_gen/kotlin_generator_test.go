package kotlin_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestKotlinGenerator(t *testing.T) {
	g := NewKotlinGenerator()
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
	if !strings.Contains(code, "data class User") {
		t.Fatalf("missing data class declaration")
	}
	if !strings.Contains(code, "val name: String") {
		t.Fatalf("missing name field")
	}
	if !strings.Contains(code, "val age: Int") {
		t.Fatalf("missing age field")
	}
	if !strings.Contains(code, "val score: Double") {
		t.Fatalf("missing score field")
	}
	if !strings.Contains(code, "val active: Boolean") {
		t.Fatalf("missing active field")
	}
	if !strings.Contains(code, "val tags: List<Any>") {
		t.Fatalf("missing tags field")
	}
	if !strings.Contains(code, "val meta: Any") {
		t.Fatalf("missing meta field")
	}
	if !strings.Contains(code, "val createdAt: Any") {
		t.Fatalf("missing createdAt field")
	}
	if !strings.Contains(code, "val anyValue: Any") {
		t.Fatalf("missing anyValue field")
	}
}
