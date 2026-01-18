package scala_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestScalaGenerator(t *testing.T) {
	g := NewScalaGenerator()
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
	if !strings.Contains(code, "case class User") {
		t.Fatalf("missing case class declaration")
	}
	if !strings.Contains(code, "name: String") {
		t.Fatalf("missing name field")
	}
	if !strings.Contains(code, "age: Int") {
		t.Fatalf("missing age field")
	}
	if !strings.Contains(code, "score: Double") {
		t.Fatalf("missing score field")
	}
	if !strings.Contains(code, "active: Boolean") {
		t.Fatalf("missing active field")
	}
	if !strings.Contains(code, "tags: List[Any]") {
		t.Fatalf("missing tags field")
	}
	if !strings.Contains(code, "meta: Map[String, Any]") {
		t.Fatalf("missing meta field")
	}
	if !strings.Contains(code, "createdAt: Any") {
		t.Fatalf("missing createdAt field")
	}
	if !strings.Contains(code, "anyValue: Any") {
		t.Fatalf("missing anyValue field")
	}
}
