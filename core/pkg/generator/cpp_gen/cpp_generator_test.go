package cpp_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestCppGenerator(t *testing.T) {
	g := NewCppGenerator()
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
	if !strings.Contains(code, "struct User") {
		t.Fatalf("missing struct declaration")
	}
	if !strings.Contains(code, "std::string name") {
		t.Fatalf("missing name field")
	}
	if !strings.Contains(code, "int age") {
		t.Fatalf("missing age field")
	}
	if !strings.Contains(code, "double score") {
		t.Fatalf("missing score field")
	}
	if !strings.Contains(code, "bool active") {
		t.Fatalf("missing active field")
	}
	if !strings.Contains(code, "std::vector<std::string> tags") {
		t.Fatalf("missing tags field")
	}
	if !strings.Contains(code, "std::map<std::string, std::string> meta") {
		t.Fatalf("missing meta field")
	}
	if !strings.Contains(code, "auto createdAt") {
		t.Fatalf("missing createdAt field")
	}
	if !strings.Contains(code, "auto anyValue") {
		t.Fatalf("missing anyValue field")
	}
}
