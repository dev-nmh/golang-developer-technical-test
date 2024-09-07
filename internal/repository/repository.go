package repository

import (
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) CountByWhere(db *gorm.DB, where map[string]interface{}) (int64, error) {
	var total int64
	tx := db.Model(new(T))

	for key, val := range where {
		if val != nil {
			tx.Where(key+" = ?", val)
		} else {
			tx.Where(key + " IS NULL")
		}
	}
	tx.Count(&total)
	return total, tx.Error
}

func (r *Repository[T]) FindByWhere(db *gorm.DB, entity *T, where map[string]interface{}) error {
	tx := db.Model(new(T))
	for key, val := range where {
		if val != nil {
			tx.Where(key+" = ?", val)
		} else {
			tx.Where(key + " IS NULL")
		}
	}
	return tx.Take(entity).Error
}
