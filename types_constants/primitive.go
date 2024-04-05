package types

type SPICE_ADDRESS uint64

// 28 high bits are signed integer. Low 4 bits is unsigned integer numerator of a fraction with denominator 16.
type SPICE_FIXED28_4 uint32

type POINT struct {
	X int32
	Y int32
}

type POINT16 struct {
	X int16
	Y int16
}
type RECT struct {
	Top    int32
	Left   int32
	Bottom int32
	Right  int32
}

type POINTFIX struct {
	X SPICE_FIXED28_4
	Y SPICE_FIXED28_4
}

type SPICE_TICKET_PUBKEY_BYTES [162]uint8

func GetSigned_Fixed28_4(num SPICE_FIXED28_4) uint32 {
	return (uint32)(num >> 4)
}

func GetUnsigned_Fixed28_4(num SPICE_FIXED28_4) uint8 {
	return (uint8)(0x0000000F & num)
}

func GetFraction_Fixed28_4(num SPICE_FIXED28_4) float64 {
	return float64(GetUnsigned_Fixed28_4(num)) / 16.0
}
