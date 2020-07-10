package state

import (
	"fmt"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Message struct {
	Text  string
	Color rl.Color
}

type MessageLog struct {
	X        int
	Width    int
	Height   int
	Messages []Message
}

func (m *MessageLog) AddMessage(message Message) {
	var lines []string
	for len(message.Text) > 0 {
		var length int
		if len(message.Text) < m.Width {
			length = len(message.Text)
		} else {
			length = m.Width
		}
		line := strings.Trim(message.Text[:length], " ")
		lines = append(lines, line)
		message.Text = message.Text[length:]
	}
	for _, line := range lines {
		if len(m.Messages) == m.Height {
			m.Messages = m.Messages[1:]
		}
		m.Messages = append(m.Messages, Message{line, message.Color})
	}
	fmt.Println(len(m.Messages), m.Height)
}
