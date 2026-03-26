package codegen

// Property represents a single field on a model.
type Property struct {
	Name          string
	OriginalName  string
	Type          string
	OpenAPIType   string
	Format        string
	IsArray       bool
	ArrayItemType string
	IsObject      bool
	IsEnum        bool
	EnumValues    []string
	Required      bool
	Nullable      bool
	Description   string
}

// Model represents a schema component (struct, enum, or type alias).
type Model struct {
	Name          string
	OriginalName  string
	Description   string
	Properties    []Property
	IsEnum        bool
	EnumValues    []string
	IsArray       bool
	ArrayItemType string
}
