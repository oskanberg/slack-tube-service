package main

type Attachment struct {
	Fallback	string
	Color		string
	Pretext		string
	Text		string
}

type SlackResponse struct {
	Text		string
	Attachments	[]Attachment
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