package handler

import (
	"driver/data"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ViewBookings(driverId int) {

}

func AuthenticateUser(email, password string) error {
	log.Println("Authenticating User")
	pwd, err := data.AuthenticateUser(email, password)
	if err != nil {
		log.Println("Error Authenticating User")
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(password)); err != nil {
		log.Println("Invalid Passwod: ", err)
		return err
	}
	log.Println("Redirect User to home page")
	return nil

}

func ViewBookingsByDriverId(driverId int) (bookings []data.Booking, err error) {
	bookings, err = data.GetBookingsForDriverId(driverId)
	if err != nil {
		return bookings, err
	}
	return bookings, nil
}

func UnAssignDriverFromBooking(bookingId, driverId int) (err error) {
	// make api call to Bookings api
	return err
}

func CompleteBookingStatus(bookingId int) (err error) {
	//make api call to Boikings api
	return err
}
