package dart_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestDartGenerator(t *testing.T) {
	g := NewDartGenerator()
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
	if !strings.Contains(code, "class User") {
		t.Fatalf("missing class declaration")
	}
	if !strings.Contains(code, "final String name") {
		t.Fatalf("missing name field")
	}
	if !strings.Contains(code, "final int age") {
		t.Fatalf("missing age field")
	}
	if !strings.Contains(code, "final double score") {
		t.Fatalf("missing score field")
	}
	if !strings.Contains(code, "final bool active") {
		t.Fatalf("missing active field")
	}
	if !strings.Contains(code, "final List<dynamic> tags") {
		t.Fatalf("missing tags field")
	}
	if !strings.Contains(code, "final dynamic meta") {
		t.Fatalf("missing meta field")
	}
	if !strings.Contains(code, "final dynamic createdAt") {
		t.Fatalf("missing createdAt field")
	}
	if !strings.Contains(code, "final dynamic anyValue") {
		t.Fatalf("missing anyValue field")
	}
	if !strings.Contains(code, "required this.name") {
		t.Fatalf("missing constructor field")
	}
}
