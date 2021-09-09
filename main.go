package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"
	"log"
	"net/http"
)

// User Record
type User struct {
	Id                     string `json:"Id"`
	SecretCode             string `json:"-"` //ensuring that secretcode is not displayed
	Name                   string `json:"Name"`
	Address                string `json:"Address"`
	PhoneNumber            string `json:"PhoneNumber"`
	Email                  string `json:"Email"`
	UserType               string `json:"UserType"`
	DiseaseDescription     string `json:"DiseaseDescription"`
	RequestedUsersIds      string `json:"RequestedUsersIds"`
	PendingRequestUsersIds string `json:"PendingRequestUsersIds"`
	ConnectedUsersIds      string `json:"ConnectedUsersIds"`
}

var Users []User

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Donor and Patient Application HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

//returning all donors based on secretcode to ensure only patients can view patient data
func returnAllDonors(w http.ResponseWriter, r *http.Request) {
	if key := r.FormValue("secretcode"); key != "" {
		for _, user := range Users {
			if user.SecretCode == key {
				if user.UserType == "Patient" { //checking if secretcode belongs to a patient
					for _, retuser := range Users {
						if retuser.UserType == "Donor" {
							json.NewEncoder(w).Encode(retuser)
						}
					}
				} else {
					fmt.Fprintf(w, "Being a donor, you cannot view other donors!")
				}
			}
		}
		fmt.Println("Endpoint Hit: returnAllDonors")
	}
}

//returning all patients based on secretcode to ensure only donors can view patient data
func returnAllPatients(w http.ResponseWriter, r *http.Request) {
	if key := r.FormValue("secretcode"); key != "" {
		for _, user := range Users {
			if user.SecretCode == key {
				if user.UserType == "Donor" { //checking if secretcode belongs to a donor
					for _, retuser := range Users {
						if retuser.UserType == "Patient" {
							json.NewEncoder(w).Encode(retuser)
						}
					}
				} else {
					fmt.Fprintf(w, "Being a patient, you cannot view other patients!")
				}
			}
		}
		fmt.Println("Endpoint Hit: returnAllPatients")
	}
}

//return user based on UserID
func returnUser(w http.ResponseWriter, r *http.Request) {
	if key := r.FormValue("id"); key != "" {
		for _, user := range Users {
			if user.Id == key {
				json.NewEncoder(w).Encode(user)
			}
		}
	}
}

//creating new user
func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		reqBody, _ := ioutil.ReadAll(r.Body) //body of POST request is fetched
		var user User
		json.Unmarshal(reqBody, &user)
		Users = append(Users, user)
		json.NewEncoder(w).Encode(user)
	}
}

//deleting user based on UserID
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		if key := r.FormValue("secretcode"); key != "" {

			for index, user := range Users {
				if user.SecretCode == key {
					Users = append(Users[:index], Users[index+1:]...)
				}
			}
		}
	}
}

//updating User's contact info based on UserID
func updateUser(w http.ResponseWriter, r *http.Request) { //concurrency safe -> mutex
	if r.Method == http.MethodPut {
		if key := r.FormValue("secretcode"); key != "" {
			var updatedInfo User
			reqBody, _ := ioutil.ReadAll(r.Body)
			json.Unmarshal(reqBody, &updatedInfo)

			for i, user := range Users {
				if user.SecretCode == key {
					user.PhoneNumber = updatedInfo.PhoneNumber
					user.Email = updatedInfo.Email
					Users[i] = user
					json.NewEncoder(w).Encode(user)
				}
			}
		}
	}
}

//login using secretcode
func login(w http.ResponseWriter, r *http.Request) {
	if key := r.FormValue("secretcode"); key != "" {
		for _, user := range Users {
			if user.SecretCode == key {
				json.NewEncoder(w).Encode(user)
			}
		}
	}

}

//sendDonorRequest
func sendDonorRequest(w http.ResponseWriter, r *http.Request) {
	if key := r.FormValue("secretcode"); key != "" {
		if id := r.FormValue("id"); id != "" {
			for _, user := range Users {
				if user.SecretCode == key {
					if user.UserType == "Patient" { //checking if secretcode belongs to a donor
						requestedId := user.Id
						for _, retuser := range Users {
							if retuser.Id == id {
								retuser.RequestedUsersIds = requestedId
								fmt.Fprint(w, "Requested Donor ID", requestedId)
								json.NewEncoder(w).Encode(retuser)
							}
						}
					} else {
						fmt.Fprintf(w, "You can send requests to Donors only!")
					}
				}
			}
			fmt.Println("Endpoint Hit: sendRequests")
		}
	}
}

//sendPatientRequest
func sendPatientRequest(w http.ResponseWriter, r *http.Request) {
	if key := r.FormValue("secretcode"); key != "" {
		if id := r.FormValue("id"); id != "" {
			for _, user := range Users {
				if user.SecretCode == key {
					if user.UserType == "Donor" { //checking if secretcode belongs to a donor
						requestedId := user.Id
						for _, retuser := range Users {
							if retuser.Id == id {
								json.NewEncoder(w).Encode(retuser)
								retuser.RequestedUsersIds = requestedId
								fmt.Fprint(w, "Requested Patient ID", requestedId)
							}
						}
					} else {
						fmt.Fprintf(w, "You can send requests to Patients only!")
					}
				}
			}
			fmt.Println("Endpoint Hit: sendRequests")
		}
	}
}

//acceptPatientRequest
func acceptPatientRequest(w http.ResponseWriter, r *http.Request) {
	if key := r.FormValue("secretcode"); key != "" {
		if id := r.FormValue("id"); id != "" {
			for _, user := range Users {
				if user.SecretCode == key {
					if user.UserType == "Donor" { //checking if secretcode belongs to a donor
						requestedId := user.Id
						for _, retuser := range Users {
							if retuser.Id == id {
								json.NewEncoder(w).Encode(retuser)
								retuser.RequestedUsersIds = requestedId
								fmt.Fprint(w, "Requested Patient ID", requestedId)
							}
						}
					} else {
						fmt.Fprintf(w, "You can send requests to Patients only!")
					}
				}
			}
			fmt.Println("Endpoint Hit: sendRequests")
		}
	}
}

//acceptdonor
func acceptDonorRequest(w http.ResponseWriter, r *http.Request) {
	if key := r.FormValue("secretcode"); key != "" {
		if id := r.FormValue("id"); id != "" {
			for _, user := range Users {
				if user.SecretCode == key {
					if user.UserType == "Patient" { //checking if secretcode belongs to a patient
						requestedId := user.Id
						for _, retuser := range Users {
							if retuser.Id == id {
								json.NewEncoder(w).Encode(retuser)
								retuser.RequestedUsersIds = requestedId
								fmt.Fprint(w, "Accepted Patient ID", requestedId)
							}
						}
					} else {
						fmt.Fprintf(w, "You can send requests to Patients only!")
					}
				}
			}
			fmt.Println("Endpoint Hit: sendRequests")
		}
	}
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login/", login)                     //login using secretcode
	http.HandleFunc("/createUser", createUser)            //.Methods("POST")                       //user will have to submit a HTTP POST request
	http.HandleFunc("/deleteUser/", deleteUser)           //.Methods("DELETE")        //to ensure users can delete only their data
	http.HandleFunc("/updateContactDetails/", updateUser) //.Methods("PUT") //to ensure users can update only their contact info
	http.HandleFunc("/getUser/", returnUser)
	http.HandleFunc("/getAllDonors/", returnAllDonors)     //to ensure only patients can read donor data
	http.HandleFunc("/getAllPatients/", returnAllPatients) //to ensure only donors can read patient data
	http.HandleFunc("/sendRequestToDonor", sendDonorRequest)
	http.HandleFunc("/sendRequestToPatient", sendPatientRequest)

	http.HandleFunc("/acceptDonorRequest", acceptDonorRequest)
	http.HandleFunc("/acceptPatientRequest", acceptPatientRequest)

	http.HandleFunc("/cancelConnection", homePage)
	http.HandleFunc("/cancelRequest", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil)) //remove
}

func main() {
	Users = []User{
		{Id: "1", SecretCode: "P1", Name: "Ana", Address: "Bhilai", Email: "ana@gmail.com", UserType: "Patient", DiseaseDescription: "pneumonia", RequestedUsersIds: "0", PendingRequestUsersIds: "0", ConnectedUsersIds: "0"},
		{Id: "2", SecretCode: "D1", Name: "Avani", Address: "Bhilai", Email: "avani.opjs@gmail.com", UserType: "Donor", DiseaseDescription: "N/A", RequestedUsersIds: "0", PendingRequestUsersIds: "0", ConnectedUsersIds: "0"},
		{Id: "3", SecretCode: "D2", Name: "Avi", Address: "Delhi", Email: "avi@gmail.com", UserType: "Donor", DiseaseDescription: "N/A", RequestedUsersIds: "0", PendingRequestUsersIds: "1", ConnectedUsersIds: "0"},
		{Id: "4", SecretCode: "P2", Name: "John", Address: "Bangalore", Email: "john@gmail.com", UserType: "Patient", DiseaseDescription: "cancer", RequestedUsersIds: "0", PendingRequestUsersIds: "0", ConnectedUsersIds: "0"},
	} //dummy data

	//generate ids and username in createuser
	//statemngmnt in flutter
	handleRequests()
}
