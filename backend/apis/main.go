package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

type WebsiteStatusRequest struct {
	Url string `json:"url"`
}

type WebsiteStatusResponse struct {
	Status string `json:"status"`
}

func isWebsiteUp(url string) string {
	_, err := http.Get(url)
	if err != nil {
		return "up"
	} else {
		return "down"
	}
}

func IsValidUrl(requestData WebsiteStatusRequest) bool {
	regex := regexp.MustCompile(`^((http|https):\/\/)?([\w-]+\.)+([\w-]+)(:\d+)?(\/|\/([\w-#!:.?+=&%! \-\/]))?`)
	match := regex.MatchString(requestData.Url)
	return match
}

func websiteStatus(w http.ResponseWriter, r *http.Request) {
	var requestData WebsiteStatusRequest
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	match := IsValidUrl(requestData)
	if !match {
		http.Error(w, "Not a valid url", http.StatusBadRequest)
	}
	var responseData WebsiteStatusResponse
	var isUp = isWebsiteUp(requestData.Url)
	responseData.Status = isUp
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to this website!")
	})
	r.HandleFunc("/websiteStatus", websiteStatus).Methods("POST")

	http.ListenAndServe(":80", r)
}
