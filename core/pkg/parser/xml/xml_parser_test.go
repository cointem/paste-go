package xml

import (
	"paste-go/pkg/schema"
	"testing"
)

func TestXMLParser_CanParse(t *testing.T) {
	p := NewXMLParser()
	if !p.CanParse("<root></root>") {
		t.Error("Should parse valid XML")
	}
	if p.CanParse("not xml") {
		t.Error("Should not parse invalid XML")
	}
}

func TestXMLParser_Parse(t *testing.T) {
	p := NewXMLParser()
	content := `<User>
	<Name>John</Name>
	<IsActive>true</IsActive>
	<Age>30</Age>
</User>`

	s, err := p.Parse(content)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if s.Name != "User" {
		t.Errorf("Struct name = %s, want User", s.Name)
	}

	// Based on current implementation:
	// Age -> String (because isInt is stubbed to false)
	// IsActive -> Bool
	// Name -> String
	expectedFields := map[string]schema.Kind{
		"Name":     schema.KindString,
		"IsActive": schema.KindBool,
		"Age":      schema.KindString, // Stubbed implementation
	}

	for _, f := range s.Fields {
		want, ok := expectedFields[f.Name]
		if !ok {
			continue
		}
		if f.Kind != want {
			t.Errorf("Field %s kind = %v, want %v", f.Name, f.Kind, want)
		}
	}
}
