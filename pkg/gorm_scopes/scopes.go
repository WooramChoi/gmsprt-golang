package gorm_scopes

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Pageable struct {
	Page     int
	PageSize int
	Sort     []struct {
		Column string
		IsDESC bool
	}
}

// https://gorm.io/ko_KR/docs/scopes.html#pagination
func Paginate(pageable *Pageable) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// if page <= 0 {
		// 	page = 1
		// }

		// switch {
		// case pageSize > 100:
		// 	pageSize = 100
		// case pageSize <= 0:
		// 	pageSize = 10
		// }

		// offset := (page - 1) * pageSize
		// return db.Offset(offset).Limit(pageSize)
		if pageable.Page <= 0 {
			pageable.Page = 1
		}

		switch {
		case pageable.PageSize > 100:
			pageable.PageSize = 100
		case pageable.PageSize <= 0:
			pageable.PageSize = 10
		}

		offset := (pageable.Page - 1) * pageable.PageSize
		db.Offset(offset).Limit(pageable.PageSize)
		for _, sort := range pageable.Sort {
			db.Order(clause.OrderByColumn{Column: clause.Column{Name: sort.Column}, Desc: sort.IsDESC})
		}

		return db
	}
}

func WhereEqual(column string, value interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s = ?", column), value)
	}
}

// TODO WhereNotEqual
// TODO WhereFrom
// TODO WhereTo
// TODO WhereLike
// TODO WhereIsTrue
// TODO WhereIsFalse
