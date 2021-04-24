package gencode_handler

import (
	"github.com/evelive3/go-gin-api/internal/pkg/core"
)

func (h *handler) HandlerView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("gencode_handler", nil)
	}
}
