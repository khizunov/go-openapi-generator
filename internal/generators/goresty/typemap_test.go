package goresty_test

import (
	"testing"

	"github.com/khizunov/go-openapi-generator/internal/generators/goresty"
	"github.com/stretchr/testify/assert"
)

func TestMapType(t *testing.T) {
	cases := []struct{ typ, format, want string }{
		{"string", "", "string"},
		{"string", "date-time", "time.Time"},
		{"string", "date", "time.Time"},
		{"string", "uuid", "string"},
		{"integer", "", "int"},
		{"integer", "int32", "int32"},
		{"integer", "int64", "int64"},
		{"number", "", "float64"},
		{"number", "float", "float32"},
		{"boolean", "", "bool"},
		{"object", "", "map[string]any"},
		{"array", "", "[]any"},
		{"", "", "any"},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, goresty.MapType(tc.typ, tc.format), "%s/%s", tc.typ, tc.format)
	}
}
