package external

import (
	enums "github.com/khizunov/go-openapi-generator/internal/definitions/enums"
)

type GenerateHandler interface {
	Handle(filepath string, specVersion enums.SpecVersion) error
}
