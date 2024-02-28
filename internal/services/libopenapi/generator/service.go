package libopenapi

import (
	"github.com/sirupsen/logrus"
)

type Generator struct {
	logger logrus.FieldLogger
}

func NewGenerator(logger logrus.FieldLogger) *Generator {
	return &Generator{logger: logger}
}
