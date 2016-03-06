package main

type Attachment struct {
	Fallback string `json:"fallback"`
	Color    string `json:"color"`
	Pretext  string `json:"pretext"`
	Text     string `json:"text"`
}

type SlackResponse struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

func mapTflLineToSlackAttachment(report Report) Attachment {
	var slackAttachment Attachment
	status := report.LineStatuses[0]
	slackAttachment.Text = report.Name + " :: " + status.StatusSeverityDescription
	slackAttachment.Color = mapTflStatuServerityToSlackColor(status.StatusSeverity)
	return slackAttachment
}

var severity = map[int]string{
	0:  "danger",
	1:  "danger",
	2:  "danger",
	3:  "danger",
	4:  "danger",
	5:  "danger",
	6:  "danger",
	7:  "warning",
	8:  "warning",
	9:  "warning",
	10: "good",
	11: "danger",
	12: "warning",
	13: "warning",
	14: "warning",
	15: "danger",
	16: "danger",
	17: "danger",
	18: "good",
	19: "good",
	20: "danger",
}

func mapTflStatuServerityToSlackColor(statusSeverity int) string {
	return severity[statusSeverity]
}
