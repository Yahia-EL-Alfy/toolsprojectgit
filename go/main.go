package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "yahia2002"
	dbname   = "toolsproject"
)

var userid int
var db *sql.DB

type User struct {
	ID        int    `json:"user_id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsDoctor  bool   `json:"is_doctor"`
}

type Slot struct {
	Date  string `json:"date"`
	Hour  string `json:"hour"`
	Empty bool   `json:"empty"`
}

type Appointment struct {
	AP              int    `json:"AP"`
	DID             int    `json:"DID"`
	Date            string `json:"Date"`
	DoctorFirstName string `json:"DoctorFirstName"`
	DoctorLastName  string `json:"DoctorLastName"`
	Hour            string `json:"Hour"`
	PID             int    `json:"PID"`
	SID             int    `json:"SID"`
}

type SlotResponse struct {
	SlotID int    `json:"slot_id"`
	Date   string `json:"date"`
	Hour   string `json:"hour"`
	Empty  bool   `json:"empty"`
}

func initDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")
}

var secretKey = []byte("yayamaya")

func generateToken(userID int, isDoctor bool) string {
	claims := jwt.MapClaims{
		"user_id":   userID,
		"is_doctor": isDoctor,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}

func authenticateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Token not provided", http.StatusUnauthorized)
			return
		}

		// Remove the "Bearer " prefix
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("yayamaya"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			fmt.Println("Invalid token:", err)
			return
		}

		// Print the actual token value
		fmt.Println("Token:", tokenString)

		// Pass the user ID and role to the next handler
		r = r.WithContext(context.WithValue(r.Context(), "user_id", int(claims["user_id"].(float64))))
		r = r.WithContext(context.WithValue(r.Context(), "is_doctor", claims["is_doctor"].(bool)))

		next.ServeHTTP(w, r)
	})
}

func AddSlot(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	var slot Slot
	err := json.NewDecoder(r.Body).Decode(&slot)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		fmt.Println("Raw Request Body:", r.Body)
		return
	}

	// Log the raw request body

	_, err = db.Exec("INSERT INTO slots (user_id, date, hour) VALUES ($1, $2, $3)", userID, slot.Date, slot.Hour)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Slot added successfully")
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO users (password, email, first_name, last_name, is_doctor) VALUES ($1, $2, $3, $4, $5)",
		user.Password, user.Email, user.FirstName, user.LastName, user.IsDoctor)
	if err != nil {
		log.Fatal(err) // This will print the detailed error message to the console
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User %s signed up successfully", user.FirstName)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received SignIn request")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	var dbUser User
	err = db.QueryRow("SELECT id, email, password, is_doctor FROM users WHERE email = $1 AND password = $2",
		user.Email, user.Password).Scan(&userid, &dbUser.Email, &user.Password, &dbUser.IsDoctor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Send a JSON response with the is_doctor field
	jsonResponse := map[string]interface{}{
		"is_doctor": dbUser.IsDoctor,
	}

	token := generateToken(userid, dbUser.IsDoctor)
	// Update the values of jsonResponse within the if block
	jsonResponse["token"] = token

	json.NewEncoder(w).Encode(jsonResponse)
	print(token)
}

func GetAllDoctors(w http.ResponseWriter, r *http.Request) {
	// Query the database to get all doctors

	rows, err := db.Query("SELECT id, email, first_name, last_name FROM users WHERE is_doctor = true")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var doctors []User
	for rows.Next() {
		var doctor User
		err := rows.Scan(&doctor.ID, &doctor.Email, &doctor.FirstName, &doctor.LastName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		doctors = append(doctors, doctor)
	}

	// Return the list of doctors as JSON response
	jsonResponse, err := json.Marshal(doctors)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetDoctorSlots(w http.ResponseWriter, r *http.Request) {
	// Get the doctor ID from the query parameters
	doctorID := r.URL.Query().Get("doctor_id")
	if doctorID == "" {
		http.Error(w, "Doctor ID is required", http.StatusBadRequest)
		return
	}

	// Query the database to get all empty slots for the specified doctor
	rows, err := db.Query("SELECT slot_id, date, hour, empty FROM slots WHERE user_id = $1 AND empty = true", doctorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var slots []SlotResponse
	for rows.Next() {
		var slot SlotResponse
		err := rows.Scan(&slot.SlotID, &slot.Date, &slot.Hour, &slot.Empty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		slots = append(slots, slot)
	}

	// Return the list of slots as JSON response
	jsonResponse, err := json.Marshal(slots)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func MakeAppointment(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		SlotID int `json:"slot_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Get the doctor ID for the given slot ID
	var doctorID int
	err = db.QueryRow("SELECT user_id FROM slots WHERE slot_id = $1", requestData.SlotID).Scan(&doctorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert the appointment into the appointments table
	_, err = db.Exec("INSERT INTO appointments (doctor_id, patient_id, slot_id) VALUES ($1, $2, $3)", doctorID, userid, requestData.SlotID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the corresponding slot to mark it as not empty
	_, err = db.Exec("UPDATE slots SET empty = false WHERE slot_id = $1", requestData.SlotID)
	if err != nil {
		http.Error(w, "Error updating slot status", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Appointment made successfully")
}

func GetUserAppointments(w http.ResponseWriter, r *http.Request) {
	// Query the database to get all appointments for the specified patient (user)
	rows, err := db.Query(`
		SELECT a.appointment_id as AP, a.doctor_id as DID, a.patient_id as PID, a.slot_id as SID, s.date as Date, s.hour as Hour, u.first_name as DoctorFirstName, u.last_name as DoctorLastName
		FROM appointments a
		JOIN slots s ON a.slot_id = s.slot_id
		JOIN users u ON a.doctor_id = u.id
		WHERE a.patient_id = $1
	`, userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var appointments []Appointment
	for rows.Next() {
		var appointment Appointment
		err := rows.Scan(
			&appointment.AP,
			&appointment.DID,
			&appointment.PID,
			&appointment.SID,
			&appointment.Date,
			&appointment.Hour,
			&appointment.DoctorFirstName,
			&appointment.DoctorLastName,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		appointments = append(appointments, appointment)
	}

	// Return the list of appointments as JSON response
	jsonResponse, err := json.Marshal(appointments)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func CancelAppointment(w http.ResponseWriter, r *http.Request) {
	// Use the existing userid variable
	userID := userid

	// Get the appointment and slot IDs for the given user ID
	var appointmentID, slotID int
	err := db.QueryRow("SELECT appointment_id, slot_id FROM appointments WHERE patient_id = $1", userID).Scan(&appointmentID, &slotID)
	if err != nil {
		http.Error(w, "Error retrieving appointment and slot IDs for the user", http.StatusInternalServerError)
		return
	}

	// Delete the appointment
	_, err = db.Exec("DELETE FROM appointments WHERE appointment_id = $1", appointmentID)
	if err != nil {
		http.Error(w, "Error deleting appointment", http.StatusInternalServerError)
		return
	}

	// Update the corresponding slot to mark it as empty (true) again
	_, err = db.Exec("UPDATE slots SET empty = true WHERE slot_id = $1", slotID)
	if err != nil {
		http.Error(w, "Error updating slot status", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Appointment canceled successfully")
}

func GetDoctorSlotsByID(w http.ResponseWriter, r *http.Request) {
	// Use the existing userid variable as the doctor ID

	// Query the database to get all empty slots for the specified doctor
	rows, err := db.Query("SELECT date, hour, empty FROM slots WHERE user_id = $1", userid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var slots []SlotResponse
	for rows.Next() {
		var slot SlotResponse
		err := rows.Scan(&slot.Date, &slot.Hour, &slot.Empty)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		slots = append(slots, slot)
	}

	// Return the list of empty slots as JSON response
	jsonResponse, err := json.Marshal(slots)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	initDB()

	r := mux.NewRouter()
	r.Handle("/addslot", authenticateToken(http.HandlerFunc(AddSlot))).Methods("POST")
	r.Handle("/signin", http.HandlerFunc(SignIn)).Methods("POST")
	r.Handle("/signup", http.HandlerFunc(SignUp)).Methods("POST")
	r.Handle("/getdoctors", http.HandlerFunc(GetAllDoctors)).Methods("GET")
	r.Handle("/getdoctorslots", http.HandlerFunc(GetDoctorSlots)).Methods("GET")
	r.Handle("/makeappointment", authenticateToken(http.HandlerFunc(MakeAppointment))).Methods("POST")
	r.Handle("/getuserappointments", authenticateToken(http.HandlerFunc(GetUserAppointments))).Methods("GET")
	r.Handle("/cancelappointment", authenticateToken(http.HandlerFunc(CancelAppointment))).Methods("POST")
	r.Handle("/getdoctorslotsbyid", authenticateToken(http.HandlerFunc(GetDoctorSlotsByID))).Methods("GET")

	// Enable CORS
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:4200"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.ExposedHeaders([]string{"Content-Length"}),
		handlers.AllowCredentials(),
	)

	// Wrap the router with CORS middleware
	http.Handle("/", corsMiddleware(r))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
