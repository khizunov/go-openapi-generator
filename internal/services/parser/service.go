package parser

import (
	"github.com/sirupsen/logrus"
)

type Parser struct {
	logger logrus.FieldLogger
}

func NewParser(logger logrus.FieldLogger) *Parser {
	return &Parser{logger: logger}
}
