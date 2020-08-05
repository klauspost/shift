package shift

import "math/bits"

const (
	// Masks for shifts with register sizes of the shift value.
	// This can be used to work around the x86 design of shifting by mod register size.
	// It can be used when a variable shift is always smaller than the register size.

	// mask(ValueType)(ShiftType)
	maskU8U8  = 7
	maskU16U8 = 15
	maskU32U8 = 31
	maskU64U8 = 63
	maskUU8   = bits.UintSize - 1

	maskU8U16  = maskU8U8
	maskU16U16 = maskU16U8
	maskU32U16 = maskU32U8
	maskU64U16 = maskU64U8
	maskUU16   = maskUU8

	maskU8U32  = maskU8U8
	maskU16U32 = maskU16U8
	maskU32U32 = maskU32U8
	maskU64U32 = maskU64U8
	maskUU32   = maskUU8

	maskU8U64  = maskU8U8
	maskU16U64 = maskU16U8
	maskU32U64 = maskU32U8
	maskU64U64 = maskU64U8
	maskUU64   = maskUU8

	maskU8U  = maskU8U8
	maskU16U = maskU16U8
	maskU32U = maskU32U8
	maskU64U = maskU64U8
	maskUU   = maskUU8

	maskI8U8  = 7
	maskI16U8 = 15
	maskI32U8 = 31
	maskI64U8 = 63
	maskIU8   = bits.UintSize - 1

	maskI8U16  = maskI8U8
	maskI16U16 = maskI16U8
	maskI32U16 = maskI32U8
	maskI64U16 = maskI64U8
	maskIU16   = maskIU8

	maskI8U32  = maskI8U8
	maskI16U32 = maskI16U8
	maskI32U32 = maskI32U8
	maskI64U32 = maskI64U8
	maskIU32   = maskIU8

	maskI8U64  = maskI8U8
	maskI16U64 = maskI16U8
	maskI32U64 = maskI32U8
	maskI64U64 = maskI64U8
	maskIU64   = maskIU8

	maskI8U  = maskI8U8
	maskI16U = maskI16U8
	maskI32U = maskI32U8
	maskI64U = maskI64U8
	maskIU   = maskIU8
)
