package services

import (
	"back/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var dbconn *gorm.DB

// Template for all HTTP response
type Response struct {
	Data    []models.User `json:"data"`
	Message string        `json:"message"`
}

// HTTP GET
// Get an array of 50 users
// @params w http response; r request fields delivered by user
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n\n")
	fmt.Print(r)
	fmt.Print("\n\n")
	w.Header().Set("Content-Type", "application/json")
	var users = models.GetUsers()
	var resp Response
	// Get only the 50 elements from DB
	err := dbconn.Limit(50).Find(&users).Error
	if err == nil {
		resp.Data = users
		resp.Message = "SUCCESS"
		json.NewEncoder(w).Encode(&resp)
	} else {
		http.Error(w, err.Error(), 400)
	}
}

// HTTP GET
// Get a User given her id
// @params w http response; r request fields delivered by user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var resp Response
	var user = models.GetUser()
	err := dbconn.Where("id = ?", id).Find(&user).Error
	if err == nil {
		resp.Data = append(resp.Data, user)
		resp.Message = "SUCCESS"
		json.NewEncoder(w).Encode(&resp)
	} else {
		http.Error(w, err.Error(), 400)
	}
}

// HTTP POST
// Create a user, given the data demonstred in models/User.go
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resp Response
	var user = models.GetUser()
	json.NewDecoder(r.Body).Decode(&user)
	user.Password, _ = HashPassword(user.Password) // For sake of simplicity, we ignore error here

	err := dbconn.Create(&user).Error
	if err != nil {
		http.Error(w, "Error Creating Record", 400)
		return
	}
	resp.Message = "CREATED"
	json.NewEncoder(w).Encode(resp)
}

// HTTP PUT
// Update a User.
// User ID is given in GET. All other data is given in request body
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var resp Response
	var user = models.GetUser()
	_ = json.NewDecoder(r.Body).Decode(&user)

	id, _ := strconv.Atoi(params["id"])

	err := dbconn.Model(&user).Where("id = ?", id).Update(&user).Error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp.Message = "UPDATED"
	json.NewEncoder(w).Encode(resp)
}

// HTTP PUT
// Delete a User.
// User ID is given in GET. All other data is given in request body
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var resp Response
	var user = models.GetUser()

	id, _ := strconv.Atoi(params["id"])

	err := dbconn.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp.Message = "DELETED"
	json.NewEncoder(w).Encode(resp)
}

func HashPassword(Password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
	return string(bytes), err
}

// AutoMigrate
func SetDB(db *gorm.DB) {
	dbconn = db
	var user = models.GetUser()
	dbconn.AutoMigrate(&user)
}
