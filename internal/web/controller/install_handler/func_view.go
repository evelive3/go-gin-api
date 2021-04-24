package install_handler

import (
	"github.com/evelive3/go-gin-api/configs"
	"github.com/evelive3/go-gin-api/internal/pkg/core"
)

func (h *handler) View() core.HandlerFunc {
	return func(c core.Context) {
		c.HTML("install_view", configs.Get())
	}
}
