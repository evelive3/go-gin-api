package admin_service

import (
	"github.com/evelive3/go-gin-api/internal/api/repository/db_repo/admin_repo"
	"github.com/evelive3/go-gin-api/internal/pkg/cache"
	"github.com/evelive3/go-gin-api/internal/pkg/core"
	"github.com/evelive3/go-gin-api/internal/pkg/password"
)

func (s *service) ResetPassword(ctx core.Context, id int32) (err error) {
	model := admin_repo.NewModel()
	model.Id = id

	data := map[string]interface{}{
		"password":     password.ResetPassword(),
		"updated_user": ctx.UserName(),
	}

	err = model.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), data)
	if err != nil {
		return err
	}

	s.cache.Del(cacheKeyPrefix+password.GenerateLoginToken(id), cache.WithTrace(ctx.Trace()))
	return
}
