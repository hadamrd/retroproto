package msgcli

import (
	"fmt"
	"html"
	"strings"

	"github.com/hadamrd/retroproto"
	"github.com/hadamrd/retroproto/typ"
)

type ChatSend struct {
	ChatChannel     typ.ChatChannel
	PrivateReceiver string
	Message         string
	Params          string // TODO
}

func NewChatSend(extra string) (ChatSend, error) {
	var m ChatSend

	if err := m.Deserialize(extra); err != nil {
		return ChatSend{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ChatSend) MessageId() retroproto.MsgCliId {
	return retroproto.ChatSend
}

func (m ChatSend) MessageName() string {
	return "ChatSend"
}

func (m ChatSend) Serialized() (string, error) {
	dest := string(m.ChatChannel)
	if m.ChatChannel == typ.ChatChannelPrivate {
		dest = m.PrivateReceiver
	}

	return fmt.Sprintf("%s|%s|%s", dest, m.Message, m.Params), nil
}

func (m *ChatSend) Deserialize(extra string) error {
	sli := strings.SplitN(extra, "|", 3)
	if len(sli) != 3 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] == "" {
		return retroproto.ErrInvalidMsg
	}

	var r rune
	for _, v := range sli[0] {
		r = v
		break
	}
	chatChannel := typ.ChatChannel(r)

	switch chatChannel {
	case '¤':
		chatChannel = typ.ChatChannelNewbies
	case typ.ChatChannelPrivate:
		return retroproto.ErrInvalidMsg
	}

	if len(sli[0]) >= 2 {
		chatChannel = typ.ChatChannelPrivate
		m.PrivateReceiver = html.EscapeString(sli[0])
	}

	_, ok := typ.ChatChannels[chatChannel]
	if !ok {
		return retroproto.ErrInvalidMsg
	}

	m.ChatChannel = chatChannel

	message := html.EscapeString(sli[1])
	message = strings.TrimSpace(message)
	m.Message = message

	m.Params = html.EscapeString(sli[2])

	return nil
}
