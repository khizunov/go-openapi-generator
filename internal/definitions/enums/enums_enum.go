// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package definitions

import (
	"errors"
	"fmt"
)

const (
	// SpecVersionV2 is a SpecVersion of type v2.
	// swagger 2.0
	SpecVersionV2 SpecVersion = "v2"
	// SpecVersionV3 is a SpecVersion of type v3.
	// openapi 3.x
	SpecVersionV3 SpecVersion = "v3"
)

var ErrInvalidSpecVersion = errors.New("not a valid SpecVersion")

// String implements the Stringer interface.
func (x SpecVersion) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x SpecVersion) IsValid() bool {
	_, err := ParseSpecVersion(string(x))
	return err == nil
}

var _SpecVersionValue = map[string]SpecVersion{
	"v2": SpecVersionV2,
	"v3": SpecVersionV3,
}

// ParseSpecVersion attempts to convert a string to a SpecVersion.
func ParseSpecVersion(name string) (SpecVersion, error) {
	if x, ok := _SpecVersionValue[name]; ok {
		return x, nil
	}
	return SpecVersion(""), fmt.Errorf("%s is %w", name, ErrInvalidSpecVersion)
}