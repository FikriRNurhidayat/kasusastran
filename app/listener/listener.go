package listener

import "encoding/json"

type Listener interface {
	Parse(message []byte, out interface{}) (err error)
}

type listener struct{}

func NewListener() Listener {
	return &listener{}
}

func (l *listener) Parse(message []byte, out interface{}) (err error) {
	return json.Unmarshal(message, out)
}
