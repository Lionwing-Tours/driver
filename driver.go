package main

import (
	//"fmt"
	"log"
	//"os"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fvbock/endless"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"driver/handler"
)

var (
	MySqlDB *sql.DB
	uid     string
	jres    string
)

func Register(res http.ResponseWriter, req *http.Request) {
	fname := req.FormValue("fname")
	lname := req.FormValue("lname")
	email := req.FormValue("email")
	password := req.FormValue("password")
	phone := req.FormValue("phone")

	err := data.RegisterDriverProto(fname, lname, email, password, phone)
	if err != nil {
		log.Println("ERROR registering driver: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusFound)
}

func Login(res http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")
	err := handler.AuthenticateUser(email, password)
	if err != nil {
		log.Println("Invalid Username or password")
		w.Write([]byte("ERROR User does not exist"))
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusFound)
}

func ChangePassword(res http.ResponseWriter, req *http.Request) {

}

func ForgotPassword(res http.ResponseWriter, req *http.Request) {

}

func UpdateVehicleProfile(w http.ResponseWriter, r *http.Request) {
	//tour/booking charges
	Rate1PerDay := r.FormValue("rate1-per-day")
	Rate2PerDay := r.FormValue("rate2-per-day")
	Rate3PerDay := r.FormValue("rate3-per-day")
	Rate1MileageLimit := r.FormValue("rate1-mileage-limit")
	Rate2MileageLimit := r.FormValue("rate2-mileage-limit")
	Rate3MileageLimit := r.FormValue("rate3-mileage-limit")
	Rate1ExtraKmCharge := r.FormValue("rate1-extra-km-charge")
	Rate2ExtraKmCharge := r.FormValue("rate2-extra-km-charge")
	Rate3ExtraKmCharge := r.FormValue("rate3-extra-km-charge")

	err := data.UpdateVehicleProfile(Rate1PerDay, Rate2PerDay, Rate3PerDay, Rate1MileageLimit,
		Rate2MileageLimit, Rate3MileageLimit, Rate1PerDayKmLimit,
		Rate1ExtraKmCharge, Rate2ExtraKmCharge, Rate3ExtraKmCharge)

	if err != nil {
		log.Println("ERROR updating driver profile: ", err)
		return
	}
	w.WriteHeader(http.StatusFound)
}

func UpdateInfo(res http.ResponseWriter, req *http.Request) {

}

func ViewBookings(w http.ResponseWriter, r *http.Request) {
	driverId, _ := strconv.Atoi(r.FormValue("driver_id"))
	bookings, err := handler.ViewBookingsByDriverId(driverId)
	if err != nil {
		w.Write([]byte(err))
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, err := json.Marshal(bookings)
	if err != nil {
		w.Write([]byte(err))
		return
	}
	w.Write([]byte(string(result)))
	w.WriteHeader(http.StatusFound)
}

func ViewFeedback(w http.ResponseWriter, r *http.Request) {
	driverId, _ := strconv.Atoi(r.FormValue("driverId"))
}

func ViewInvoices(w http.ResponseWriter, r *http.Request) {
	driverId, _ := strconv.Atoi(r.FormValue("driverId"))

}

func CompleteBookingStatus(w http.ResponseWriter, r *http.Request) {
	bookingId, _ := strconv.Atoi(r.FormValue("bookingId"))
	err := handler.CompleteBookingStatus(bookingId)
	if err != nil {
		log.Println(err)
		return
	}
}

func UnassignFromBooking(r http.ResponseWriter, r *http.Request) {
	//make request to booking server
	driverId, _ := strconv.Atoi(r.FormValue("driverId"))
	bookingId, _ := strconv.Atoi(r.FormValue("bookingId"))

	err := handler.UnAssignDriverFromBooking(bookingId, driverId)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/register", Register)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/changePassword", ChangePassword)
	http.HandleFunc("/forgotPassword", ForgotPassword)
	http.HandleFunc("/updateVehicleProfile", UpdateVehicleProfile)
	http.HandleFunc("/completeBookingStatus", CompleteBookingStatus)
	http.HandleFunc("/UpdateInfo", UpdateInfo)
	http.HandleFunc("/updateAvailability", UpdateAvailability)
	http.HandleFunc("/viewBookings", ViewBookings)
	http.HandleFunc("/viewFeedback", ViewFeedback)
	http.HandleFunc("/viewInvoices", ViewInvoices)
	http.HandleFunc("/unassign", unassignFromBooking)

	log.Fatal(endless.ListenAndServe("localhost:8081", nil))
}

func init() {
	var err error

	mysqlDb, err := sql.Open("mysql", "root:rootroot@/TestDb")
	MySqlDB = mysqlDb
	MySqlDB.SetMaxOpenConns(10)
	MySqlDB.SetMaxIdleConns(0)
	if err != nil {
		panic(err.Error())
	}

}
