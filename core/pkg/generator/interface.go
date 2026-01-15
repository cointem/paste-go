package generator

import "paste-forge/pkg/schema"

type Generator interface {
	Name() string
	Supports(lang string) bool
	Generate(s *schema.Struct) (string, error)
}
