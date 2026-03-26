package codegen

import (
	"fmt"
	"sync"
)

var (
	mu       sync.RWMutex
	registry = map[string]Generator{}
)

// Register adds a generator to the global registry.
// Panics on duplicate name.
func Register(g Generator) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := registry[g.Name()]; exists {
		panic(fmt.Sprintf("codegen: generator %q already registered", g.Name()))
	}
	registry[g.Name()] = g
}

// Lookup returns the generator with the given name or an error.
func Lookup(name string) (Generator, error) {
	mu.RLock()
	defer mu.RUnlock()
	g, ok := registry[name]
	if !ok {
		return nil, fmt.Errorf("codegen: generator %q not found; registered: %v", name, names())
	}
	return g, nil
}

// All returns all registered generators.
func All() []Generator {
	mu.RLock()
	defer mu.RUnlock()
	all := make([]Generator, 0, len(registry))
	for _, g := range registry {
		all = append(all, g)
	}
	return all
}

func names() []string {
	ns := make([]string, 0, len(registry))
	for n := range registry {
		ns = append(ns, n)
	}
	return ns
}
