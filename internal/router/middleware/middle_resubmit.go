package middleware

import (
	"net/http"
	"time"

	"github.com/evelive3/go-gin-api/configs"
	"github.com/evelive3/go-gin-api/internal/api/code"
	"github.com/evelive3/go-gin-api/internal/pkg/cache"
	"github.com/evelive3/go-gin-api/internal/pkg/core"
	"github.com/evelive3/go-gin-api/pkg/errno"
	"github.com/evelive3/go-gin-api/pkg/token"

	"github.com/pkg/errors"
)

func (m *middleware) Resubmit() core.HandlerFunc {

	redisKeyPrefix := configs.ProjectName() + ":request-id:"

	return func(c core.Context) {
		cfg := configs.Get().URLToken

		tokenString, err := token.New(cfg.Secret).UrlSign(c.Path(), c.Method(), c.RequestInputParams())
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ResubmitError,
				code.Text(code.ResubmitError)).WithErr(err),
			)
			return
		}

		redisKey := redisKeyPrefix + tokenString
		if !m.cache.Exists(redisKey) {
			err = m.cache.Set(redisKey, "1", time.Minute*cfg.ExpireDuration)
			if err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.ResubmitError,
					code.Text(code.ResubmitError)).WithErr(err),
				)
				return
			}

			return
		}

		redisValue, err := m.cache.Get(redisKey, cache.WithTrace(c.Trace()))
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ResubmitError,
				code.Text(code.ResubmitError)).WithErr(err),
			)
			return
		}

		if redisValue == "1" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ResubmitMsg,
				code.Text(code.ResubmitMsg)).WithErr(errors.New("resubmit")),
			)
			return
		}

		return
	}
}
