package listener_test

import (
	"testing"

	"github.com/fikrirnurhidayat/kasusastran/app/listener"
	"github.com/stretchr/testify/assert"
)

func TestListener_Parse(t *testing.T) {
	lis := listener.NewListener()
	t.Run("OK", func(t *testing.T) {
		var msg struct {
			Name string `json:"name"`
		}

		err := lis.Parse([]byte("{\"name\": \"Fikri\"}"), &msg)

		assert.Nil(t, err)
		assert.Equal(t, msg.Name, "Fikri")
	})
}
