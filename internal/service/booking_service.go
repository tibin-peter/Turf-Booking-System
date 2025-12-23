package service

import (
	"errors"
	"time"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

type BookingService struct {
	repo repository.Repository
}

func NewBookingService(repo repository.Repository) *BookingService {
	return &BookingService{repo: repo}
}

//func for create booking
func (s *BookingService) CreateBooking(b *model.Booking) error {

	// fetch slot
	var slot model.TimeSlot
	if err := s.repo.FindById(&slot, b.SlotID); err != nil {
		return errors.New("slot not found")
	}

	if !slot.IsAvailable {
		return errors.New("slot already booked")
	}

	// fetch turf
	var turf model.Turf
	if err := s.repo.FindById(&turf, slot.TurfID); err != nil {
		return errors.New("turf not found")
	}

	// calculate duration
	start, _ := time.Parse("15:04", slot.StartTime)
	end, _ := time.Parse("15:04", slot.EndTime)
	hours := end.Sub(start).Hours()

	if hours <= 0 {
		return errors.New("invalid slot duration")
	}

	// Set amount
	b.TotalAmount = int(hours) * turf.PricePerHour
	b.TurfID = turf.ID
	b.Status = "pending"
	b.PaymentStatus = "pending"

	if err := s.repo.Insert(b); err != nil {
		return err
	}

	slot.IsAvailable = false
	return s.repo.Update(&slot)
}

// func for user confirmation
func (s *BookingService) ConfirmPayment(bookingID, userID uint) error {

	var booking model.Booking
	if err := s.repo.FindById(&booking, bookingID); err != nil {
		return errors.New("booking not found")
	}

	if booking.UserID != userID {
		return errors.New("unauthorized")
	}

	if booking.PaymentStatus != "pending" {
		return errors.New("payment already processed")
	}

	// Create payment record
	payment := model.Payment{
		BookingID: booking.ID,
		Amount:    booking.TotalAmount,
		Status:    "pending",
	}

	if err := s.repo.Insert(&payment); err != nil {
		return err
	}

	booking.Status = "approved"
	return s.repo.Update(&booking)
}

// func for list user booking
func (s *BookingService) ListUserBookings(userID uint) ([]model.Booking, error) {
	var bookings []model.Booking
	err := s.repo.FindMany(&bookings, "user_id = ?", userID)
	return bookings, err
}

// func for cancel the booking
func (s *BookingService) CancelBooking(bookingID, userID uint) error {

	var booking model.Booking
	if err := s.repo.FindById(&booking, bookingID); err != nil {
		return errors.New("booking not found")
	}

	if booking.UserID != userID {
		return errors.New("unauthorized")
	}

	booking.Status = "cancelled"
	booking.PaymentStatus = "refunded"

	if err := s.repo.Update(&booking); err != nil {
		return err
	}

	var slot model.TimeSlot
	if err := s.repo.FindById(&slot, booking.SlotID); err == nil {
		slot.IsAvailable = true
		_ = s.repo.Update(&slot)
	}

	return nil
}
