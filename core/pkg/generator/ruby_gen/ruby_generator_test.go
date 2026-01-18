package ruby_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestRubyGenerator(t *testing.T) {
	g := NewRubyGenerator()
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
	if !strings.Contains(code, "attr_accessor :name, :age, :score, :active, :tags, :meta, :createdAt, :anyValue") {
		t.Fatalf("missing attr_accessor")
	}
}
