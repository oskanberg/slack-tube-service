package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"strings"
	"time"
)

const minStatusPollPeriod = 2

var statuses []Report

func lineStatusHandler(w http.ResponseWriter, r *http.Request) {

	var response []Report

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if isUpdateNeeded() {
		if err := updateStatusInformation(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			if err := json.NewEncoder(w).Encode("There was an error getting information from TFL"); err != nil {
				log.Panic(err)
			}
		}
	}

	vars := mux.Vars(r)
	tubeLine, lineIsPresentInPath := vars["line"]

	if !lineIsPresentInPath {
		for _, line := range statuses {
			response = append(response, mapTflLineToResponse(line))
		}
	} else {
		for _, line := range statuses {
			if strings.ToLower(line.Name) == strings.ToLower(tubeLine) {
				response = append(response, mapTflLineToResponse(line))
			}
		}
		if len(response) == 0 {
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode("Not a recognised line."); err != nil {
				log.Panic(err)
			}
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Panic(err)
	}
}

func slackRequestHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Panic(err)
	}

	var slackRequest = new(SlackRequest)
	decoder := schema.NewDecoder()
	err = decoder.Decode(slackRequest, r.PostForm)
	if err != nil {
		log.Panic(err)
	}

	var slackResponse SlackResponse
	var attachments []Attachment

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if slackRequest.Token != "MGFfi8oiCo4EDoQNVb0YE2gl" {
		w.WriteHeader(http.StatusUnauthorized)
		slackResponse.Text = "Unauthorised"
	} else {

		if isUpdateNeeded() {
			if err := updateStatusInformation(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				slackResponse.Text = "There was an error getting information from TFL"
			}
		}

		tubeLine := slackRequest.Text
		fmt.Print(r.PostForm)

		w.WriteHeader(http.StatusOK)
		slackResponse.Text = "Slack Tube Service"

		if tubeLine == "" {
			for _, line := range statuses {
				attachments = append(attachments, mapTflLineToSlackAttachment(line))
			}
		} else {
			for _, line := range statuses {
				if strings.ToLower(line.Name) == strings.ToLower(tubeLine) {
					attachments = append(attachments, mapTflLineToSlackAttachment(line))
				}
			}
			if len(attachments) == 0 {
				w.WriteHeader(http.StatusNotFound)
				slackResponse.Text = "Not a recognised line."
			}
		}

		slackResponse.Attachments = attachments
	}

	if err := json.NewEncoder(w).Encode(slackResponse); err != nil {
		log.Panic(err)
	}
}

func isUpdateNeeded() bool {
	return time.Since(lastStatusCheck).Minutes() > minStatusPollPeriod
}

func updateStatusInformation() error {
	url := "https://api.tfl.gov.uk/line/mode/tube/status"

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	var data []Report
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	statuses = data
	return nil
}
