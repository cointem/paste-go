package ai

import (
	"fmt"
	"strings"
)

// Factory defines a function that returns a new instance of a Provider.
type Factory func() Provider

// registry stores the available provider factories.
var registry = make(map[string]Factory)

// Register adds a provider factory to the registry.
// This is typically called in the init() function of the provider package.
func Register(name string, factory Factory) {
	registry[strings.ToLower(name)] = factory
}

// GetProvider returns a new instance of the requested provider.
func GetProvider(name string) (Provider, error) {
	factory, ok := registry[strings.ToLower(name)]
	if !ok {
		return nil, fmt.Errorf("provider '%s' not found. Available providers: %s", name, listProviders())
	}
	return factory(), nil
}

func listProviders() string {
	keys := make([]string, 0, len(registry))
	for k := range registry {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}
