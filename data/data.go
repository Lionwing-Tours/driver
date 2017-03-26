package data

import (
	"database/sql"
	"driver/db"
	"errors"
	"log"
)

var (
	ErrUserDoesNotExist = errors.New("User Does Not Exist")
	ErrInvalidPassword  = errors.New("Invalid Password")
	ErrNoRows           = errors.New("No Bookings Available")
)

const (
	getBookingsForDriverIdQuery = `SELECT passengers.first_name, 
											passengers.last_name,
											bookings.pickup_location,
											bookings.pickup_address,
											bookings.start_date,
											bookings.end_date,
											bookings.no_passengers,
											bookings.estimated_cost
									FROM passengers
										INNER JOIN bookings
											ON passengers.id=bookings.passenger_id
											WHERE bookings.booking_id=?`

	registerDriverQuery = `INSERT INTO driver(name, lastname, email, sec_password, phone) VALUES(?,?,?,?,?)`

	UpdateVehicleProfileQuery = ` `
)

type Booking struct {
	PassengerFirstName string
	PassengerLastName  string
	PickupLocation     string
	StartDate          string
	EndDate            string
	NoOfPassengers     int
	EstimatedCost      float64
}

func GetBookingsForDriverId(driverId int) (bookings []Booking, err error) {
	rows, err := db.MySqlDB.Query(getBookingsForDriverIdQuery, driverId)
	if err != nil {
		log.Println("ERROR in getting live drivers: ", err.Error())
		return bookings, err
	}
	if rows == nil {
		log.Println("ERROR in getting live drivers: nil rows")
		return bookings, ErrNoRows
	}
	defer rows.Close()

	for rows.Next() {
		b := Booking{}
		//driver.driver_id,driver.status,driver.latitude,driver.longitude,driver.update_date,people.name,company.company_name,taxi.taxi_model
		err = rows.Scan(&b.PassengerFirstName, &b.PassengerLastName, &b.PickupLocation, &b.StartDate, &b.EndDate, &b.NoOfPassengers, &b.EstimatedCost)
		if err != nil {
			log.Println("ERROR parsing rows in live drivers: ", err.Error())
			return bookings, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func AuthenticateUser(email, password string) (string, error) {
	var Email string
	var Password string

	row := db.MySqlDB.QueryRow("SELECT email FROM passengers WHERE email = ?", email)
	err := row.Scan(&Email)
	if err != nil {
		if err == sql.ErrNoRows || len(Email) == 0 {
			log.Println("ERROR fetching passenger device details from DB:", err)
			return "", ErrUserDoesNotExist
		}
	}

	row = db.MySqlDB.QueryRow("SELECT password FROM passengers WHERE passengers.email = ?", email)
	err = row.Scan(&password)
	if err != nil {
		log.Println("Undefined Error")
		return "", err
	}
	return Password, nil
}

func RegisterDriverProto(fname, lname, email, password, phone string) error {
	result, err := MySqlDB.Exec(registerDriverQuery, fname, lname, email, password, phone)
	if err != nil {
		log.Println(result, err)
		return err
	}

	log.Println("Table updated", result)
	return nil
}

func UpdateVehicleProfile(Rate1PerDay, Rate2PerDay, Rate3PerDay, Rate1MileageLimit, Rate2MileageLimit, Rate3MileageLimit, Rate1ExtraKmCharge, Rate2ExtraKmCharge, Rate3ExtraKmCharge string) error {
	// might have to use redis locks
}
