package state_manager

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"log"
	constants "masala/types_constants"
	types "masala/types_constants"
	"sync"
)

type State struct {
	Error       uint32
	PrivateKey  *rsa.PrivateKey
	PublicKey   types.SPICE_TICKET_PUBKEY_BYTES
	Connections map[uint32]ConnectionState
}

var globalState *State
var writeLock *sync.Mutex

func GetState() *State {
	var lock sync.Mutex
	if globalState == nil {
		lock.Lock()
		defer lock.Unlock()
		if globalState == nil {
			globalState = new(State)
			globalState.Intialize()
		}
	}

	return globalState
}

func (s *State) Intialize() {
	privateKey, err := rsa.GenerateKey(rand.Reader, constants.SPICE_RSA_KEY_SIZE)

	if err != nil {
		log.Fatalln("Could not generate key, please restart service")
	}

	publicKey, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalln("Could not generate key, please restart service")
	}

	s.PrivateKey = privateKey
	s.PublicKey = constants.SPICE_TICKET_PUBKEY_BYTES(publicKey)
	s.Connections = make(map[uint32]ConnectionState)
}

func (s *State) SetState(stateWriter func()) {
	writeLock.Lock()
	defer writeLock.Unlock()

	stateWriter()
}
