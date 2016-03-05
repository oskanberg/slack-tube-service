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

func mapTflStatuServerityToSlackColor(statusSeverity int) string {
	switch {
	case statusSeverity < 5:
		return "danger"
	case statusSeverity < 9:
		return "warning"
	default:
		return "good"
	}
}
