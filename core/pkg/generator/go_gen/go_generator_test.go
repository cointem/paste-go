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
			{Name: "Name", OriginalName: "name", Kind: schema.KindString},
			{Name: "Age", OriginalName: "age", Kind: schema.KindInt},
			{Name: "Score", OriginalName: "score", Kind: schema.KindFloat},
			{Name: "Active", OriginalName: "active", Kind: schema.KindBool},
			{Name: "Tags", OriginalName: "tags", Kind: schema.KindArray},
			{
				Name:         "Meta",
				OriginalName: "meta",
				Kind:         schema.KindObject,
				Nested: &schema.Struct{
					Name: "Meta",
					Fields: []schema.Field{
						{Name: "Foo", OriginalName: "foo", Kind: schema.KindString},
						{Name: "Bar", OriginalName: "bar", Kind: schema.KindInt},
					},
				},
			},
			{Name: "CreatedAt", OriginalName: "created_at", Kind: schema.KindTime},
			{Name: "AnyValue", OriginalName: "any_value", Kind: schema.KindAny},
		},
	}

	code, err := gen.Generate(s)
	if err != nil {
		t.Fatalf("Generate error: %v", err)
	}

	if !strings.Contains(code, "type TestStruct struct {") {
		t.Error("Missing struct definition")
	}
	if !strings.Contains(code, "Name string `json:\"name\"`") {
		t.Error("Missing string field")
	}
	if !strings.Contains(code, "Age int64 `json:\"age\"`") {
		t.Error("Missing int field")
	}
	if !strings.Contains(code, "Score float64 `json:\"score\"`") {
		t.Error("Missing float field")
	}
	if !strings.Contains(code, "Active bool `json:\"active\"`") {
		t.Error("Missing bool field")
	}
	if !strings.Contains(code, "Tags []interface{} `json:\"tags\"`") {
		t.Error("Missing array field")
	}
	if !strings.Contains(code, "Meta struct {") {
		t.Error("Missing object field")
	}
	if !strings.Contains(code, "Foo string `json:\"foo\"`") {
		t.Error("Missing nested Foo field")
	}
	if !strings.Contains(code, "Bar int64 `json:\"bar\"`") {
		t.Error("Missing nested Bar field")
	}
	if !strings.Contains(code, "CreatedAt time.Time `json:\"created_at\"`") {
		t.Error("Missing time field")
	}
	if !strings.Contains(code, "AnyValue interface{} `json:\"any_value\"`") {
		t.Error("Missing any field")
	}
}
