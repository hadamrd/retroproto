// File: msgcli/account_credential.go
package msgcli

import (
	"fmt"
	"strings"

	"github.com/hadamrd/retroproto"
)

type AccountCredential struct {
	Token string
}

func NewAccountCredential(extra string) (AccountCredential, error) {
	var m AccountCredential
	if err := m.Deserialize(extra); err != nil {
		return AccountCredential{}, fmt.Errorf("could not deserialize message: %w", err)
	}
	return m, nil
}

func (m AccountCredential) MessageId() retroproto.MsgCliId {
	return retroproto.AccountCredential
}

func (m AccountCredential) MessageName() string {
	return "AccountCredential"
}

func (m AccountCredential) Serialized() (string, error) {
	return fmt.Sprintf("#Z\n%s", m.Token), nil
}

func (m *AccountCredential) Deserialize(extra string) error {
	// Expected format: #Z\n{token}
	if !strings.HasPrefix(extra, "#Z\n") {
		return retroproto.ErrInvalidMsg
	}
	
	m.Token = strings.TrimPrefix(extra, "#Z\n")
	
	if m.Token == "" {
		return retroproto.ErrInvalidMsg
	}
	
	return nil
}