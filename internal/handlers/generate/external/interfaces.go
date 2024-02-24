package external

import (
	v2 "github.com/pb33f/libopenapi/datamodel/high/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

type GeneratorAPI interface {
	GenerateFromV3(doc *v3.Document) error
	GenerateFromV2(doc *v2.Swagger) error
}

type ParserAPI interface {
	ParseV3(filename string) (doc *v3.Document, err error)
	ParseV2(filename string) (doc *v2.Swagger, err error)
}
