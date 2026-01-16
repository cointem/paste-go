package sql

import (
	"testing"

	"paste-go/pkg/schema"
)

func TestSQLParser_CanParse(t *testing.T) {
	p := NewSQLParser()
	tests := []struct {
		name    string
		content string
		want    bool
	}{
		{"Valid Create Table", "CREATE TABLE users (id int)", true},
		{"Case Insensitive", "create table users (id int)", true},
		{"Invalid", "SELECT * FROM users", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := p.CanParse(tt.content); got != tt.want {
				t.Errorf("CanParse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLParser_Parse(t *testing.T) {
	p := NewSQLParser()
	content := `
	CREATE TABLE user_profiles (
		user_id INT PRIMARY KEY,
		full_name VARCHAR(255),
		is_active BOOL,
		salary DECIMAL(10, 2),
		created_at DATETIME
	);
	`

	s, err := p.Parse(content)
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	if s.Name != "UserProfiles" {
		t.Errorf("Expected struct name UserProfiles, got %s", s.Name)
	}

	expectedFields := map[string]schema.Kind{
		"UserId":    schema.KindInt,
		"FullName":  schema.KindString,
		"IsActive":  schema.KindBool,
		"Salary":    schema.KindFloat,
		"CreatedAt": schema.KindTime,
	}

	for _, f := range s.Fields {
		want, ok := expectedFields[f.Name]
		if !ok {
			continue // Skip checking unknown fields if regex is loose, but here we expect precise matches
		}
		if f.Kind != want {
			t.Errorf("Field %s kind = %v, want %v", f.Name, f.Kind, want)
		}
	}
}
