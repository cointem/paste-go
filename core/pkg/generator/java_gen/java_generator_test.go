package java_gen

import (
	"strings"
	"testing"

	"paste-go/pkg/schema"
)

func TestJavaGenerator_Generate(t *testing.T) {
	gen := NewJavaGenerator()
	s := &schema.Struct{
		Name: "TestClass",
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

	if !strings.Contains(code, "public class TestClass {") {
		t.Error("Missing class definition")
	}
	if !strings.Contains(code, "@Data") {
		t.Error("Missing Lombok annotation")
	}
	if !strings.Contains(code, "private String name;") {
		t.Error("Missing name field")
	}
	if !strings.Contains(code, "private Integer age;") {
		t.Error("Missing age field")
	}
	if !strings.Contains(code, "private Double score;") {
		t.Error("Missing score field")
	}
	if !strings.Contains(code, "private Boolean active;") {
		t.Error("Missing active field")
	}
	if !strings.Contains(code, "private java.util.List<Object> tags;") {
		t.Error("Missing tags field")
	}
	if !strings.Contains(code, "private Object meta;") {
		t.Error("Missing meta field")
	}
	if !strings.Contains(code, "private String created_at;") {
		t.Error("Missing created_at field")
	}
	if !strings.Contains(code, "private Object any_value;") {
		t.Error("Missing any_value field")
	}
}
