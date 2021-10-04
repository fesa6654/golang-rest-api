package admin

import (
	"encoding/json"
	"fmt"
	models "golang-rest-api/api/models/admin"
	"golang-rest-api/api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewAdmins models.Admin

func GetAdmins(w http.ResponseWriter, r *http.Request) {

	newAdmin := models.GetAllAdmins()

	res, _ := json.Marshal(newAdmin)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetAdminById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	adminId := vars["adminId"]
	ID, err := strconv.ParseInt(adminId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	adminDetails, _ := models.GetAdminById(ID)

	res, _ := json.Marshal(adminDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateAdmin(w http.ResponseWriter, r *http.Request) {

	createAdmin := &models.Admin{}

	utils.ParseBody(r, createAdmin)

	b := createAdmin.CreateNewAdmin()

	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAdminById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	adminId := vars["adminId"]
	ID, err := strconv.ParseInt(adminId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	admin := models.DeleteAdminById(ID)

	res, _ := json.Marshal(admin)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateAdmin(w http.ResponseWriter, r *http.Request) {

	var updateAdmin = &models.Admin{}

	utils.ParseBody(r, updateAdmin)

	vars := mux.Vars(r)

	adminId := vars["adminId"]

	ID, err := strconv.ParseInt(adminId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	adminsDetails, db := models.GetAdminById(ID)

	if updateAdmin.Name != "" {
		adminsDetails.Name = updateAdmin.Name
	}

	if updateAdmin.Surname != "" {
		adminsDetails.Surname = updateAdmin.Surname
	}

	db.Save(&adminsDetails)

	res, _ := json.Marshal(adminsDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
