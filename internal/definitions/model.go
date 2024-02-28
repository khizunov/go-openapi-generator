package definitions

type Model struct {
	IncludeModel       bool
	IsExported         bool
	Name               string
	ExtraSchemas       []Model
	IsBaseType         bool
	DiscriminatorField string
}
