package libopenapi

import (
	v2 "github.com/pb33f/libopenapi/datamodel/high/v2"
	v3 "github.com/pb33f/libopenapi/datamodel/high/v3"
)

func (g *Generator) GenerateFromV3(doc *v3.Document) error {
	items := doc.Paths.PathItems
	for pathName := items.First(); pathName != nil; pathName = pathName.Next() {
		operations := pathName.Value().GetOperations()
		for op := operations.First(); op != nil; op = op.Next() {
			g.logger.Infof("Processing operation %s of path %s", op.Key(), pathName.Key())
		}
	}
	return nil
}

func (g *Generator) GenerateFromV2(doc *v2.Swagger) error {
	return nil
}
