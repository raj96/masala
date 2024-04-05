package state_manager

import (
	"crypto/rsa"
	types "masala/types_constants"
	"time"
)

type Ticket struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  types.SPICE_TICKET_PUBKEY_BYTES
	expiry     time.Time
}
