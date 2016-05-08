package main

type Report struct {
	Name         string
	LineStatuses []Status
}

func mapTflLineToResponse(tflLine Report) Report {
	return tflLine
}
