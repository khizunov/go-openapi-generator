package goresty_test

import (
	"testing"

	"github.com/khizunov/go-openapi-generator/internal/generators/goresty"
	"github.com/stretchr/testify/assert"
)

func TestToModelName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Pet", "Pet"},
		{"pet", "Pet"},
		{"pet_category", "PetCategory"},
		{"petCategory", "PetCategory"},
		{"HTTPClient", "HTTPClient"},
		{"", ""},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, goresty.ToModelName(tc.in), tc.in)
	}
}

func TestToOperationName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"listPets", "ListPets"},
		{"list_pets", "ListPets"},
		{"ListPets", "ListPets"},
		{"", ""},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, goresty.ToOperationName(tc.in), tc.in)
	}
}

func TestToParamName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"pet_id", "petID"},
		{"userId", "userID"},
		{"limit", "limit"},
		{"X-Request-ID", "xRequestID"},
		{"", ""},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.want, goresty.ToParamName(tc.in), tc.in)
	}
}
