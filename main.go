package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Status struct {
	StatusSeverity            int
	StatusSeverityDescription string
}

type Report struct {
	Name         string
	LineStatuses []Status
}

var statuses []Report
var lastStatusCheck time.Time

const MIN_STATUS_POLL_PERIOD = 2

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

func getStatusString(line Report) string {
	// simplest case, this is just a single 'segment'
	if len(line.LineStatuses) == 1 {
		return line.Name + " :: " + line.LineStatuses[0].StatusSeverityDescription + "\n"
	}

	//TODO: what if there is more than one segment?
	return ""
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if time.Since(lastStatusCheck).Minutes() > MIN_STATUS_POLL_PERIOD {
		err := updateStatusInformation()
		if err != nil {
			io.WriteString(w, "There was an error getting information from TFL")
		}
	}

	queryLine := r.URL.Query().Get("line")
	if len(queryLine) == 0 {
		w.WriteHeader(200)
		for _, line := range statuses {
			io.WriteString(w, getStatusString(line))
		}
		return
	}

	for _, line := range statuses {
		if strings.ToLower(line.Name) == strings.ToLower(queryLine) {
			io.WriteString(w, getStatusString(line))
			return
		}
	}

	io.WriteString(w, "Not a recognised line")
}

func main() {
	var err error
	lastStatusCheck, err = time.Parse(time.RFC3339, "1970-01-01T00:00:00+00:00")
	if err != nil {
		return
	}

	http.HandleFunc("/", handleRequest)
	fmt.Println("Serving on port 1123")
	http.ListenAndServe(":1123", nil)
}
