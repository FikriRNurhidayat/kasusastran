package message

import (
	"time"

	"github.com/google/uuid"
)

type Message[T any] struct {
	ID        uuid.UUID `json:"id"`
	Kind      string    `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
	Actor     *Actor    `json:"actor"`
	Payload   T         `json:"payload"`
}
