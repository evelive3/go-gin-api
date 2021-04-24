package authorized_handler

import "github.com/evelive3/go-gin-api/internal/pkg/core"

func (h *handler) DemoView() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("authorized_demo", nil)
	}
}
