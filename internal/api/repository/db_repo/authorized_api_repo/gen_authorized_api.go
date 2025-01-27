///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package authorized_api_repo

import (
	"fmt"
	"time"

	"github.com/evelive3/go-gin-api/internal/api/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *AuthorizedApi {
	return new(AuthorizedApi)
}

func NewQueryBuilder() *authorizedApiRepoQueryBuilder {
	return new(authorizedApiRepoQueryBuilder)
}

func (t *AuthorizedApi) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

func (t *AuthorizedApi) Delete(db *gorm.DB) (err error) {
	if err = db.Delete(t).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (t *AuthorizedApi) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	if err = db.Model(&AuthorizedApi{}).Where("id = ?", t.Id).Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

type authorizedApiRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *authorizedApiRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *authorizedApiRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&AuthorizedApi{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *authorizedApiRepoQueryBuilder) First(db *gorm.DB) (*AuthorizedApi, error) {
	ret := &AuthorizedApi{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *authorizedApiRepoQueryBuilder) QueryOne(db *gorm.DB) (*AuthorizedApi, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *authorizedApiRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*AuthorizedApi, error) {
	var ret []*AuthorizedApi
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *authorizedApiRepoQueryBuilder) Limit(limit int) *authorizedApiRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) Offset(offset int) *authorizedApiRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereIdIn(value []int32) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereIdNotIn(value []int32) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderById(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereBusinessKey(p db_repo.Predicate, value string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereBusinessKeyIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereBusinessKeyNotIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderByBusinessKey(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_key "+order)
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereMethod(p db_repo.Predicate, value string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereMethodIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereMethodNotIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "method", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderByMethod(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "method "+order)
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereApi(p db_repo.Predicate, value string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereApiIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereApiNotIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "api", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderByApi(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "api "+order)
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereIsDeleted(p db_repo.Predicate, value int32) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereIsDeletedIn(value []int32) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereIsDeletedNotIn(value []int32) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderByIsDeleted(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereCreatedAt(p db_repo.Predicate, value time.Time) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereCreatedAtIn(value []time.Time) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderByCreatedAt(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereCreatedUser(p db_repo.Predicate, value string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereCreatedUserIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereCreatedUserNotIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderByCreatedUser(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereUpdatedAt(p db_repo.Predicate, value time.Time) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderByUpdatedAt(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereUpdatedUser(p db_repo.Predicate, value string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereUpdatedUserIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) WhereUpdatedUserNotIn(value []string) *authorizedApiRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedApiRepoQueryBuilder) OrderByUpdatedUser(asc bool) *authorizedApiRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
