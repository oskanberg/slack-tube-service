package main

type SlackRequest struct {
	Token        string
	Team_id      string
	Team_domain  string
	Channel_id   string
	Channel_name string
	User_id      string
	User_name    string
	Command      string
	Text         []string
	Response_url string
}
