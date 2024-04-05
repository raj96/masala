package state_manager

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	constants "masala/types_constants"
	types "masala/types_constants"
	"net"
	"time"
)

type ConnectionState struct {
	Connection *net.TCPConn
	KeyTicket  *Ticket
	Password   string
}

func NewConnectionState(tcpConn *net.TCPConn) *ConnectionState {
	cs := &ConnectionState{
		Connection: tcpConn,
	}

	cs.instantiateNewTicket()

	return cs
}

func (cs *ConnectionState) instantiateNewTicket() error {
	const ticketExpiryInSeconds = 180

	privateKey, err := rsa.GenerateKey(rand.Reader, constants.SPICE_RSA_KEY_SIZE)
	if err != nil {
		return err
	}

	pubKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}

	cs.KeyTicket = &Ticket{
		PrivateKey: privateKey,
		PublicKey:  types.SPICE_TICKET_PUBKEY_BYTES(pubKey),
		expiry:     time.Now().Add(time.Second * ticketExpiryInSeconds),
	}

	return nil
}

func (cs *ConnectionState) IsTicketValid() bool {
	if cs.KeyTicket == nil {
		return false
	}

	if time.Now().After(cs.KeyTicket.expiry) {
		cs.KeyTicket = nil
		return false
	}

	return true
}
