package parser

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

	v3m, v3me := d.BuildV3Model()
	if v3me != nil {
		for _, e := range v3me {
			err = multierror.Append(err, e)
		}
		return
	}

	return &v3m.Model, err
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

	model, me := d.BuildV2Model()
	if me != nil {
		for _, e := range me {
			err = multierror.Append(err, e)
		}
		return
	}

	return &model.Model, err
}
