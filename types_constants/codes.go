package types

// Error codes
const (
	SPICE_LINK_ERR_OK                    uint32 = 0 // Success
	SPICE_LINK_ERR_ERROR                 uint32 = 1
	SPICE_LINK_ERR_INVALID_MAGIC         uint32 = 2
	SPICE_LINK_ERR_INVALID_DATA          uint32 = 3
	SPICE_LINK_ERR_VERSION_MISMATCH      uint32 = 4
	SPICE_LINK_ERR_NEED_SECURED          uint32 = 5
	SPICE_LINK_ERR_NEED_UNSECURED        uint32 = 6
	SPICE_LINK_ERR_PERMISSION_DENIED     uint32 = 7
	SPICE_LINK_ERR_BAD_CONNECTION_ID     uint32 = 8
	SPICE_LINK_ERR_CHANNEL_NOT_AVAILABLE uint32 = 9
)

// Warning codes
const (
	SPICE_WARN_GENERAL uint32 = 0
)

// Information codes
const (
	SPICE_INFO_GENERAL uint32 = 0
)
