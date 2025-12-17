package repository

type Repository interface {
	Insert(req interface{}) error
	FindById(out interface{}, id uint) error
	FindOne(out interface{}, query string, args ...any) error
	Update(req interface{}) error
	Delete(query string, args ...any) error
	FindMany(out interface{}, query string, args ...any) error
	Count(model interface{}, query string, args ...any) (int64, error)
}
