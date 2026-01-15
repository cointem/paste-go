package parser

import "paste-forge/pkg/schema"

type Parser interface {
	Name() string
	CanParse(content string) bool
	Parse(content string) (*schema.Struct, error)
}
