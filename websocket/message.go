package websocket

import (
	"time"
	
	"github.com/duke-git/lancet/v2/random"
	jsoniter "github.com/json-iterator/go"
)

type Message struct {
	MessageID string `json:"messageID"`
	Username  string `json:"username"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func NewMessage(username, title, message, messageType string) *Message {
	messageID, _ := random.UUIdV4()
	return &Message{
		MessageID: messageID,
		Username:  username,
		Title:     title,
		Type:      messageType,
		Message:   message,
		Timestamp: time.Now().Unix(),
	}
}

func (m *Message) ToJson() ([]byte, error) {
	return jsoniter.Marshal(m)
}

func (m *Message) FromJson(data []byte) error {
	return jsoniter.Unmarshal(data, m)
}
