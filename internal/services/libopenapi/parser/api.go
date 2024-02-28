package libopenapi

import (
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/pb33f/libopenapi"
	v2 "github.com/pb33f/libopenapi/datamodel/high/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

func (m *Parser) ParseV3(filename string) (doc *v3.Document, err error) {
	specs, fe := os.ReadFile(filename)
	if fe != nil {
		return doc, fe
	}

	d, de := libopenapi.NewDocument(specs)
	if de != nil {
		return doc, de
	}

	model, bErr := d.BuildV3Model()
	if bErr != nil {
		for _, e := range bErr {
			err = multierror.Append(err, e)
		}
		return
	}

	return &model.Model, err
}

func (m *Parser) ParseV2(filename string) (doc *v2.Swagger, err error) {
	specs, fe := os.ReadFile(filename)
	if fe != nil {
		return doc, fe
	}

	d, de := libopenapi.NewDocument(specs)
	if de != nil {
		return doc, de
	}

	model, bErr := d.BuildV2Model()
	if bErr != nil {
		for _, e := range bErr {
			err = multierror.Append(err, e)
		}
		return
	}

	return &model.Model, err
}
