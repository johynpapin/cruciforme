package handlers

import (
	"github.com/johynpapin/cruciforme/server/email"
	"github.com/johynpapin/cruciforme/server/store"
)

type Handlers struct {
	JwtSigningKey []byte
	Store         *store.Store
	EmailSender   *email.Sender
}
