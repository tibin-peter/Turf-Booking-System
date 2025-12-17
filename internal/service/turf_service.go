package service

import (
	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

// struct for connect the repo
type TurfService struct {
	repo repository.Repository
}

// dependecy injection
func NewTurfService(repo repository.Repository) *TurfService {
	return &TurfService{repo: repo}
}

// func for get all turfs
func (s *TurfService) ListTurfs() ([]model.Turf, error) {
	var turfs []model.Turf
	//"1 = 1" means no filtering return everything
	err := s.repo.FindMany(&turfs, "1=1")
	if err != nil {
		return nil, err
	}
	return turfs, err
}

// func for get a particular turf by id
func (s *TurfService) GetTurfByID(id uint) (model.Turf, error) {
	var turf model.Turf
	err := s.repo.FindById(&turf, id)
	if err != nil {
		//return an empty turf model
		return model.Turf{}, err
	}
	return turf, nil
}
