package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	return nil
}

func main() {
	err := updateStatusInformation()
	if err != nil {
		return
	}

	for _, line := range statuses {
		if len(line.LineStatuses) == 1 {
			fmt.Println(line.Name, line.LineStatuses)
		}
	}
}
