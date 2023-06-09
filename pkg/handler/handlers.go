package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("myName")
	role := r.FormValue("myRole")
	email := r.FormValue("myEmail")
	gender := ""
	checkboxValue := ""

	getGender := r.FormValue("myGender")
	if getGender == "on" {
		gender = "Male"
	} else {
		gender = "Female"
	}

	getCheckboxValue := r.FormValue("Checkbox")
	if getCheckboxValue == "on" {
		checkboxValue = "yes"
	} else {
		checkboxValue = "no"
	}
	
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Role = %s\n", role)
	fmt.Fprintf(w, "Email = %s\n", email)
	fmt.Fprintf(w, "Gender = %s\n", gender)
	fmt.Fprintf(w, "Eligible = %s\n", checkboxValue)
}
