package schema

// Kind represents the generic type of a field
type Kind int

const (
	KindString Kind = iota
	KindInt
	KindFloat
	KindBool
	KindObject
	KindArray
	KindTime
	KindAny
)

// Field represents a single field in a structure
type Field struct {
	Name         string // Normalized name (e.g., PascalCase)
	OriginalName string // Original input key (e.g., snake_case)
	Kind         Kind
	Notes        string  // Extra info (comments, db tags etc)
	Nested       *Struct // For KindObject or KindArray
}

// Struct represents a class/struct definition
type Struct struct {
	Name   string
	Fields []Field
}
