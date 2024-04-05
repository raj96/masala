package packets

import (
	"bytes"
	"encoding/binary"
	"masala/state_manager"
	constants "masala/types_constants"
)

type SpiceLinkMess struct {
	Magic          uint32
	MajorVersion   uint32
	MinorVersion   uint32
	Size           uint32
	ConnectionId   uint32
	ChannelType    uint8
	ChannelId      uint8
	NumCommonCaps  uint32
	NumChannelCaps uint32
	CapsOffset     uint32
}

type SpiceLinkCapability uint32
type SpiceLinkCapabilities []SpiceLinkCapability

// Returns the spice link mess packet or nil on error
func ParseSpiceLinkMess(buffer []byte) (*SpiceLinkMess, *SpiceLinkCapabilities) {
	state := state_manager.GetState()

	linkMess := new(SpiceLinkMess)
	linkCaps := new(SpiceLinkCapabilities)

	err := binary.Read(bytes.NewBuffer(buffer), binary.LittleEndian, linkMess)

	if err != nil {
		state.SetState(func() {
			state.Error = constants.SPICE_LINK_ERR_INVALID_DATA
		})
		return nil, nil
	}

	if linkMess.CapsOffset != 0 {
		extensionBuffer := buffer[16+linkMess.CapsOffset : 16+linkMess.Size]
		for i := 0; i < len(extensionBuffer); i += 4 {
			*linkCaps = append(*linkCaps, (SpiceLinkCapability)(binary.LittleEndian.Uint32(extensionBuffer[i:i+4])))
		}
	}

	if linkMess.Magic != constants.SPICE_MAGIC {
		// Invalid magic TODO
		state.SetState(func() {
			state.Error = constants.SPICE_LINK_ERR_INVALID_MAGIC
		})
	}

	if linkMess.MajorVersion != constants.SPICE_VERSION_MAJOR || linkMess.MinorVersion != constants.SPICE_VERSION_MINOR {
		// Invalid version TODO
		state.SetState(func() {
			state.Error = constants.SPICE_LINK_ERR_VERSION_MISMATCH
		})
	}

	return linkMess, linkCaps
}

func (spLinkMess *SpiceLinkMess) Reply() error {

	return nil
}
