package main

type Attachment struct {
	Fallback  string   `json:"fallback"`
	Color     string   `json:"color"`
	Pretext   string   `json:"pretext"`
	Text      string   `json:"text"`
	Mrkdwn_in []string `json:"mrkdwn_in"`
}

type SlackResponse struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

func mapTflLineToSlackAttachment(report Report) Attachment {
	var slackAttachment Attachment
	status := report.LineStatuses[0]
	slackAttachment.Text = mapTflStatuServerityToSlackSeverity(status.StatusSeverity).Emoji + "  *" + report.Name + "* :: " + status.StatusSeverityDescription
	slackAttachment.Color = mapLineNameToHexColor(report.Name)
	slackAttachment.Mrkdwn_in = []string{"text"}
	return slackAttachment
}

type SlackSeverity struct {
	Color string
	Emoji string
}

var danger = SlackSeverity{"danger", ":rage:"}
var warning = SlackSeverity{"warning", ":warning:"}
var good = SlackSeverity{"good", ":grinning:"}

var severity = map[int]SlackSeverity{
	0:  danger,
	1:  danger,
	2:  danger,
	3:  danger,
	4:  danger,
	5:  danger,
	6:  danger,
	7:  warning,
	8:  warning,
	9:  warning,
	10: good,
	11: danger,
	12: warning,
	13: warning,
	14: warning,
	15: danger,
	16: danger,
	17: danger,
	18: good,
	19: good,
	20: danger,
}

var lineColors = map[string]string{
	"Bakerloo":           "#B36305",
	"Central":            "#E32017",
	"Circle":             "#FFD300",
	"District":           "#00782A",
	"Hammersmith & City": "#F3A9BB",
	"Jubilee":            "#A0A5A9",
	"Metropolitan":       "#9B0056",
	"Northern":           "#000000",
	"Piccadilly":         "#003688",
	"Victoria":           "#0098D4",
	"Waterloo & City":    "#95CDBA",
}

func mapTflStatuServerityToSlackSeverity(statusSeverity int) SlackSeverity {
	return severity[statusSeverity]
}

func mapLineNameToHexColor(lineName string) string {
	return lineColors[lineName]
}
