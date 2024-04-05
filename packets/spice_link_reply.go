package packets

import (
	"bytes"
	"encoding/binary"
	"masala/state_manager"
	constants "masala/types_constants"
	types "masala/types_constants"
)

type SpiceLinkReply struct {
	Magic          uint32
	MajorVersion   uint32
	MinorVersion   uint32
	Size           uint32
	Error          uint32
	PublicKey      types.SPICE_TICKET_PUBKEY_BYTES
	NumCommonCaps  uint32
	NumChannelCaps uint32
	CapsOffset     uint32
}

func NewSpiceLinkReply() *SpiceLinkReply {
	state := state_manager.GetState()

	linkReply := &SpiceLinkReply{
		Magic:        constants.SPICE_MAGIC,
		MajorVersion: constants.SPICE_VERSION_MAJOR,
		MinorVersion: constants.SPICE_VERSION_MINOR,
		Error:        constants.SPICE_LINK_ERR_OK,
		PublicKey:    state.PublicKey,
	}

	if state.Error != constants.SPICE_LINK_ERR_OK {
		linkReply.Error = state.Error
		return linkReply
	}

	return linkReply
}

func (linkReply *SpiceLinkReply) ToBytes() ([]byte, error) {
	var b []byte
	err := binary.Write(bytes.NewBuffer(b), binary.LittleEndian, linkReply)
	if err != nil {
		return b, err
	}

	return b, nil
}
