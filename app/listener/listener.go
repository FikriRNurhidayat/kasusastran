package listener

import "encoding/json"

// TODO: Abstract this
func Parse(message []byte, out interface{}) (err error) {
	return json.Unmarshal(message, out)
}
