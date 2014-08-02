package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

type Attachment struct {
	Fallback string  `json:"fallback,omitempty"`
	PreText  string  `json:"pretext,omitempty"`
	Text     string  `json:"text,omitempty"`
	Color    string  `json:"color,omitempty"`
	Fields   []Field `json:"fields,omitempty"`
}

type Message struct {
	Channel     string       `json:"channel,omitempty"`
	BotName     string       `json:"username,omitempty"`
	Text        string       `json:"text,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

type Bot struct {
	url     string
	Name    string
	IconURL string
}

// Posts the provided text on the webhook integration's default channel.
func (this *Bot) PublishText(text string) {
	msg := new(Message)
	msg.Text = text
	this.Publish(*msg)
}

// Posts the provided message to the webhook integration. If a channel has been set in the message, the message will
// be posted to that channel. Otherwise, the message will be posted to the default channel configured for the integration.
func (this *Bot) Publish(message Message) {
	if len(message.BotName) == 0 && len(this.Name) != 0 {
		message.BotName = this.Name
	}

	if len(message.IconURL) == 0 && len(this.IconURL) != 0 {
		message.IconURL = this.IconURL
	}

	msgJSON, _ := json.Marshal(message)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", this.url, bytes.NewReader(msgJSON))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}

// Creates a new bot for the specified team.
//
// You can configure the bot's name and icon using the Name and IconURL properties respectively.
func NewBot(team string, token string) (bot *Bot) {
	bot = &Bot{
		url: fmt.Sprintf("https://%s.slack.com/services/hooks/incoming-webhook?token=%s", team, token),
	}
	return
}

// Creates a new Field with the title, value, short values provided. It is for use in conjunction with Attachments.
func NewField(title string, value string, short bool) (f Field) {
	f = Field{
		Title: title,
		Value: value,
		Short: short,
	}

	return
}

func NewAttachment(fallback string, pretext string, text string, color string, fields []Field) (a Attachment) {
	a = Attachment{
		Fallback: fallback,
		PreText:  pretext,
		Text:     text,
		Color:    color,
		Fields:   fields,
	}

	return
}

func NewMessage(channel string, unfurlLinks bool) (m Message) {
	m = Message{Channel: channel, UnfurlLinks: unfurlLinks}
	return
}
