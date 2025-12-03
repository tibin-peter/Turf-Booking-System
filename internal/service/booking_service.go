package service

import (
	"errors"

	"github.com/tibin-peter/Turf-Booking-System/internal/model"
	"github.com/tibin-peter/Turf-Booking-System/internal/repository"
)

// Create booking
func CreteUserBooking(b *model.Booking) error {
	//validate slot exist
	slot, err := repository.GetSlotByID(b.SlotID)
	if err != nil {
		return errors.New("slot not found")
	}
	// check if slot is available
	if !slot.IsAvailable {
		return errors.New("slot already booked")
	}
	//create booking
	if err := repository.CreateBooking(b); err != nil {
		return err
	}
	//changing the availability to false
	slot.IsAvailable = false
	return repository.UpdateSlot(&slot)
}

// List all bookings for user
func ListUserBookings(userID uint) ([]model.Booking, error) {
	return repository.GetUserBookings(userID)
}

//Cancel booking

func CancelUserBooking(bookinID uint, userID uint) error {
	booking, err := repository.GetBookingByID(bookinID)
	if err != nil {
		return errors.New("booking not found")
	}

	//check the userid is same
	if booking.UserID != userID {
		return errors.New("unauthorized")
	}
	//Cancel booking
	booking.Status = "cancelled"
	booking.PaymentStatus = "refunded"

	if err := repository.UpdateBooking(&booking); err != nil {
		return err
	}
	//free the slot again
	slot, err := repository.GetSlotByID(booking.SlotID)
	if err == nil {
		slot.IsAvailable = true
		repository.UpdateSlot(&slot)
	}
	return nil
}
