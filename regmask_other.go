//+build !amd64

package shift

const (
	// Masks for shifts with register sizes of the shift value.
	// This can be used to work around the x86 design of shifting by mod register size.
	// It can be used when a variable shift is always smaller than the register size.

	maskU8U8  = 0xff
	maskU16U8 = 0xff
	maskU32U8 = 0xff
	maskU64U8 = 0xff
	maskUU8   = 0xff

	maskU8U16  = 0xffff
	maskU16U16 = 0xffff
	maskU32U16 = 0xffff
	maskU64U16 = 0xffff
	maskUU16   = 0xffff

	maskU8U32  = 0xffffffff
	maskU16U32 = 0xffffffff
	maskU32U32 = 0xffffffff
	maskU64U32 = 0xffffffff
	maskUU32   = 0xffffffff

	maskU8U64  = 0xffffffffffffffff
	maskU16U64 = 0xffffffffffffffff
	maskU32U64 = 0xffffffffffffffff
	maskU64U64 = 0xffffffffffffffff
	maskUU64   = 0xffffffffffffffff

	maskU8U  = ^uint(0)
	maskU16U = ^uint(0)
	maskU32U = ^uint(0)
	maskU64U = ^uint(0)
	maskUU   = ^uint(0)

	maskI8U8  = -1
	maskI16U8 = -1
	maskI32U8 = -1
	maskI64U8 = -1
	maskIU8   = -1

	maskI8U16  = -1
	maskI16U16 = -1
	maskI32U16 = -1
	maskI64U16 = -1
	maskIU16   = -1

	maskI8U32  = -1
	maskI16U32 = -1
	maskI32U32 = -1
	maskI64U32 = -1
	maskIU32   = -1

	maskI8U64  = -1
	maskI16U64 = -1
	maskI32U64 = -1
	maskI64U64 = -1
	maskIU64   = -1

	maskI8U  = -1
	maskI16U = -1
	maskI32U = -1
	maskI64U = -1
	maskIU   = -1
)
