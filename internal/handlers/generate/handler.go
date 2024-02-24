package generate

import (
	"github.com/khizunov/go-openapi-generator/internal/handlers/generate/external"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	generator external.GeneratorAPI
	parser    external.ParserAPI

	logger logrus.FieldLogger
}

func NewHandler(
	generator external.GeneratorAPI,
	parser external.ParserAPI,
	logger logrus.FieldLogger,
) *Handler {
	return &Handler{
		generator: generator,
		parser:    parser,
		logger:    logger.WithField("handler", "generate"),
	}
}

func (h *Handler) log(method string) logrus.FieldLogger {
	return h.logger.WithField("method", method)
}
