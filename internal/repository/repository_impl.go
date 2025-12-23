package repository

import (
	"gorm.io/gorm"
)

type PgSQLRepository struct {
	DB *gorm.DB
}

func Newrepo(db *gorm.DB) Repository {
	return &PgSQLRepository{DB: db}
}

// insert or create
func (r *PgSQLRepository) Insert(req interface{}) error {
	return r.DB.Create(req).Error
}

// find by id
func (r *PgSQLRepository) FindById(out interface{}, id uint) error {
	return r.DB.First(out, id).Error
}

// find one using the query with conditions
func (r *PgSQLRepository) FindOne(out interface{}, query string, args ...any) error {
	return r.DB.Where(query, args...).First(out).Error
}

// for update
func (r *PgSQLRepository) Update(req interface{}) error {
	return r.DB.Save(req).Error
}

// delete by query
func (r *PgSQLRepository) Delete(model interface{}, query string, args ...any) error {
	return r.DB.Where(query, args...).Delete(model).Error
}

// func for find many
func (r *PgSQLRepository) FindMany(out interface{}, query string, args ...any) error {
	return r.DB.Where(query, args...).Find(out).Error
}

// func for calculate all count
func (r *PgSQLRepository) Count(model interface{}, query string, args ...any) (int64, error) {
	var count int64
	db := r.DB.Model(model)
	if query != "" {
		db = db.Where(query, args...)
	}
	err := db.Count(&count).Error
	return count, err
}
