package service

import (
	"errors"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

type BookingService struct {
	repo repository.Repository
}

func NewBookingService(repo repository.Repository) *BookingService {
	return &BookingService{repo: repo}
}

// Create booking
func (s *BookingService) CreateUserBooking(b *model.Booking) error {
	//validate slot exist
	var slot model.TimeSlot
	err := s.repo.FindById(&slot, b.SlotID)
	if err != nil {
		return errors.New("slot not found")
	}
	// check if slot is available
	if !slot.IsAvailable {
		return errors.New("slot already booked")
	}
	//create booking
	if err := s.repo.Insert(b); err != nil {
		return err
	}
	//changing the availability to false
	slot.IsAvailable = false
	return s.repo.Update(&slot)
}

// List all bookings for user
func (s *BookingService) ListUserBookings(userID uint) ([]model.Booking, error) {
	var bookings []model.Booking
	err := s.repo.FindMany(&bookings, "user_id = ?", userID)
	if err != nil {
		return []model.Booking{}, err
	}

	return bookings, nil
}

//Cancel booking

func (s *BookingService) CancelUserBooking(bookingID uint, userID uint) error {
	var booking model.Booking

	if err := s.repo.FindById(&booking, bookingID); err != nil {
		return errors.New("booking not found")
	}

	//check the userid is same
	if booking.UserID != userID {
		return errors.New("unauthorized")
	}
	//Cancel booking
	booking.Status = "cancelled"
	booking.PaymentStatus = "refunded"

	if err := s.repo.Update(&booking); err != nil {
		return err
	}
	//free the slot again
	var slot model.TimeSlot
	err := s.repo.FindById(&slot, booking.SlotID)
	if err == nil {
		slot.IsAvailable = true
		s.repo.Update(&slot)
	}
	return nil
}

// func for the payment conformation by the user
func (s *BookingService) ConfirmPayment(bookinID uint, userID uint) error {
	//fetching the booking
	var booking model.Booking
	err := s.repo.FindById(&booking, bookinID)
	if err != nil {
		return errors.New("booking not found")
	}
	//validating the user
	if booking.UserID != userID {
		return errors.New("not have a authorized to conformation")
	}
	//check payment mehtod
	if booking.PaymentMethod != "dummy" {
		return errors.New("payment conformation allowed only for dummy payment method")
	}
	//ensure not already confirmed
	if booking.PaymentStatus == "success" {
		return errors.New("payment alreadu approved by admin")
	}
	//update status
	booking.PaymentStatus = "user_confirmed"
	booking.Status = "pending"
	return s.repo.Update(&booking)
}
