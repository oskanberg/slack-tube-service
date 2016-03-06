package main

type SlackResponse struct {
	Text        string       `json:"text"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback  string   `json:"fallback"`
	Color     string   `json:"color"`
	Pretext   string   `json:"pretext"`
	Text      string   `json:"text"`
	Mrkdwn_in []string `json:"mrkdwn_in"`
}

func mapTflLineToSlackAttachment(report Report) Attachment {
	var slackAttachment Attachment
	slackAttachment.Text = createSlackText(report)
	slackAttachment.Color = mapLineNameToHexColor(report.Name)
	slackAttachment.Mrkdwn_in = []string{"text"}
	return slackAttachment
}

func createSlackText(report Report) string {
	slackText := ""
	slackSeverity := mapTflStatuServerityToSlackSeverity(report.LineStatuses[0].StatusSeverity)
	slackText = slackText + slackSeverity.Emoji
	slackText = slackText + "  *" + report.Name + "*"
	slackText = slackText + " :: " + report.LineStatuses[0].StatusSeverityDescription
	if slackSeverity == danger || slackSeverity == warning {
		slackText = slackText + "\n" + report.LineStatuses[0].Reason
	}
	return slackText
}

func mapLineNameToHexColor(lineName string) string {
	return lineColors[lineName]
}

func mapTflStatuServerityToSlackSeverity(statusSeverity int) SlackSeverity {
	return severity[statusSeverity]
}

var danger = SlackSeverity{"danger", ":rage:"}
var warning = SlackSeverity{"warning", ":warning:"}
var good = SlackSeverity{"good", ":grinning:"}

type SlackSeverity struct {
	Color string
	Emoji string
}

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
	20: warning,
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
