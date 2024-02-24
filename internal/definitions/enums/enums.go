package definitions

//go:generate go-enum
// SpecVersion version of an API spec (v2 swagger or v3 openapi).
/* ENUM(
v2 // swagger 2.0
v3 // openapi 3.x
)*/
type SpecVersion string
