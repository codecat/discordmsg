package discordmsg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EmbedFooter struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Embed struct {
	Title       string    `json:"title,omitempty"`
	Type        string    `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Color       int       `json:"color,omitempty"`

	Footer    *EmbedFooter   `json:"footer,omitempty"`
	Image     *EmbedImage    `json:"image,omitempty"`
	Thumbnail *EmbedImage    `json:"thumbnail,omitempty"`
	Provider  *EmbedProvider `json:"provider,omitempty"`
	Author    *EmbedAuthor   `json:"author,omitempty"`
	Fields    []EmbedField   `json:"fields,omitempty"`
}

func (e *Embed) SetFooter(text, iconURL string) {
	e.Footer = &EmbedFooter{
		Text:    text,
		IconURL: iconURL,
	}
}

func (e *Embed) SetImage(url string) {
	e.Image = &EmbedImage{
		URL: url,
	}
}

func (e *Embed) SetThumbnail(url string) {
	e.Thumbnail = &EmbedImage{
		URL: url,
	}
}

func (e *Embed) SetProvider(name, url string) {
	e.Provider = &EmbedProvider{
		Name: name,
		URL:  url,
	}
}

func (e *Embed) SetAuthor(name, url, iconURL string) {
	e.Author = &EmbedAuthor{
		Name:    name,
		URL:     url,
		IconURL: iconURL,
	}
}

func (e *Embed) AddField(name, value string, inline bool) {
	if e.Fields == nil {
		e.Fields = make([]EmbedField, 0)
	}

	e.Fields = append(e.Fields, EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})
}

// Documentation: https://discord.com/developers/docs/resources/webhook#execute-webhook
type MessageData struct {
	Content   string   `json:"content,omitempty"`
	Username  string   `json:"username,omitempty"`
	AvatarURL string   `json:"avatar_url,omitempty"`
	TTS       bool     `json:"tts,omitempty"`
	Embeds    []*Embed `json:"embeds,omitempty"`

	//TODO: If we want to @ mention users or roles
	//AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`

	//TODO: Components (allows stuff like buttons)

	//TODO: More?
}

type Message struct {
	WebhookURL string
	Data       MessageData
}

func (m *Message) AddEmbed() *Embed {
	newEmbed := &Embed{}
	m.Data.Embeds = append(m.Data.Embeds, newEmbed)
	return newEmbed
}

func (m *Message) Send() error {
	res, _ := json.Marshal(m.Data)
	_, err := http.Post(m.WebhookURL, "application/json", bytes.NewReader(res))
	if err != nil {
		return fmt.Errorf("unable to run webhook: %s", err.Error())
	}
	return nil
}

func New(webhookURL string) *Message {
	return &Message{
		WebhookURL: webhookURL,

		Data: MessageData{
			Embeds: make([]*Embed, 0),
		},
	}
}
