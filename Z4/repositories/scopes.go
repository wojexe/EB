package repositories

import (
	"gorm.io/gorm"
)

// Pagination adds limit and offset to queries
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

// OrderBy adds sorting to queries
func OrderBy(column string, direction string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(column + " " + direction)
	}
}

// Search provides a generic search capability
func Search(column, query string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if query != "" {
			return db.Where(column+" LIKE ?", "%"+query+"%")
		}
		return db
	}
}
