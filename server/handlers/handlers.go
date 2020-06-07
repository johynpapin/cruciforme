package handlers

import (
	"github.com/johynpapin/cruciforme/server/store"
)

type Handlers struct {
	JwtSigningKey []byte
	Store         *store.Store
}
