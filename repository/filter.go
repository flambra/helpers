package repository

import (
	"github.com/flambra/helpers/types"
	"gorm.io/gorm"
)

type Filter interface {
	Apply(db *gorm.DB) *gorm.DB
}

type DefaultFilter struct {
	Name                      string     `query:"name"`
	CreatedGreaterOrEqualThan types.Date `query:"createdGreaterOrEqualThan"`
	CreatedLessOrEqualThan    types.Date `query:"createdLessOrEqualThan"`
}

func (d *DefaultFilter) Apply(db *gorm.DB) *gorm.DB {
	if !d.CreatedGreaterOrEqualThan.Default().IsZero() {
		db = db.Where("created_at >= ?", d.CreatedGreaterOrEqualThan.Default())
	}

	if !d.CreatedLessOrEqualThan.Default().IsZero() {
		db = db.Where("created_at < ?", d.CreatedLessOrEqualThan.Default())
	}

	if d.Name != "" {
		like := d.Name + "%"
		db = db.Where("name ILIKE ?", like)
	}

	return db
}
