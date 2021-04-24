package admin_handler

import "github.com/evelive3/go-gin-api/internal/pkg/core"

func (h *handler) AddView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("admin_add", nil)
	}
}
