package swift_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestSwiftGenerator(t *testing.T) {
	g := NewSwiftGenerator()
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
	if !strings.Contains(code, "struct User: Codable") {
		t.Fatalf("missing struct declaration")
	}
	if !strings.Contains(code, "var name: String") {
		t.Fatalf("missing name field")
	}
	if !strings.Contains(code, "var age: Int") {
		t.Fatalf("missing age field")
	}
	if !strings.Contains(code, "var score: Double") {
		t.Fatalf("missing score field")
	}
	if !strings.Contains(code, "var active: Bool") {
		t.Fatalf("missing active field")
	}
	if !strings.Contains(code, "var tags: [Any]") {
		t.Fatalf("missing tags field")
	}
	if !strings.Contains(code, "var meta: Any") {
		t.Fatalf("missing meta field")
	}
	if !strings.Contains(code, "var createdAt: Any") {
		t.Fatalf("missing createdAt field")
	}
	if !strings.Contains(code, "var anyValue: Any") {
		t.Fatalf("missing anyValue field")
	}
}
