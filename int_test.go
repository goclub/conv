package xconv

import (
	"math"
	"testing"
)

func TestIntBool(t *testing.T) {
	Cases(t, IntBool, []Case{
		{0, false, ``},
		{1, true, ``},
		{2, false, `IntBool: invalid input 2`},
		{-1, false, `IntBool: invalid input -1`},
	})
}

func TestIntString(t *testing.T) {
	Cases(t, IntString, []Case{
		{0, "0", ``},
		{123, "123", ``},
		{-456, "-456", ``},
	})
}

func TestIntInt8(t *testing.T) {
	Cases(t, IntInt8, []Case{
		{0, int8(0), ``},
		{math.MaxInt8, int8(math.MaxInt8), ``},
		{math.MinInt8, int8(math.MinInt8), ``},
		{math.MaxInt8 + 1, int8(0), `IntInt8: input out of range 128`},
		{math.MinInt8 - 1, int8(0), `IntInt8: input out of range -129`},
	})
}

func TestIntInt16(t *testing.T) {
	Cases(t, IntInt16, []Case{
		{0, int16(0), ``},
		{math.MaxInt16, int16(math.MaxInt16), ``},
		{math.MinInt16, int16(math.MinInt16), ``},
		{math.MaxInt16 + 1, int16(0), `IntInt16: input out of range 32768`},
		{math.MinInt16 - 1, int16(0), `IntInt16: input out of range -32769`},
	})
}

func TestIntInt32(t *testing.T) {
	Cases(t, IntInt32, []Case{
		{int(0), int32(0), ``},
		{math.MaxInt32, int32(math.MaxInt32), ``},
		{math.MinInt32, int32(math.MinInt32), ``},
	})
}

func TestIntInt64(t *testing.T) {
	Cases(t, IntInt64, []Case{
		{0, int64(0), ``},
		{math.MaxInt64, int64(math.MaxInt64), ``},
		{math.MinInt64, int64(math.MinInt64), ``},
	})
}

func TestIntUint(t *testing.T) {
	Cases(t, IntUint, []Case{
		{0, uint(0), ``},
		{123, uint(123), ``},
		{-456, uint(0), `IntUint: input out of range -456`},
	})
}

func TestIntUint8(t *testing.T) {
	Cases(t, IntUint8, []Case{
		{0, uint8(0), ``},
		{math.MaxUint8, uint8(math.MaxUint8), ``},
		{-1, uint8(0), `IntUint8: input out of range -1`},
		{math.MaxUint8 + 1, uint8(0), `IntUint8: input out of range 256`},
	})
}

func TestIntUint16(t *testing.T) {
	Cases(t, IntUint16, []Case{
		{0, uint16(0), ``},
		{math.MaxUint16, uint16(math.MaxUint16), ``},
		{-1, uint16(0), `IntUint16: input out of range -1`},
		{math.MaxUint16 + 1, uint16(0), `IntUint16: input out of range 65536`},
	})
}

func TestIntUint32(t *testing.T) {
	Cases(t, IntUint32, []Case{
		{0, uint32(0), ``},
		{123, uint32(123), ``},
		{-456, uint32(0), `IntUint32: input out of range -456`},
	})
}

func TestIntUint64(t *testing.T) {
	Cases(t, IntUint64, []Case{
		{0, uint64(0), ``},
		{123, uint64(123), ``},
		{-456, uint64(0), `IntUint64: input out of range -456`},
	})
}

func TestIntFloat32(t *testing.T) {
	Cases(t, IntFloat32, []Case{
		{0, float32(0), ``},
		{123, float32(123), ``},
		{-456, float32(-456), ``},
	})
}

func TestIntFloat64(t *testing.T) {
	Cases(t, IntFloat64, []Case{
		{0, float64(0), ``},
		{123, float64(123), ``},
		{-456, float64(-456), ``},
	})
}

func TestInt8Bool(t *testing.T) {
	Cases(t, Int8Bool, []Case{
		{int8(0), false, ``},
		{int8(1), true, ``},
		{int8(2), false, `Int8Bool: invalid input 2`},
	})
}

func TestInt8String(t *testing.T) {
	Cases(t, Int8String, []Case{
		{int8(0), "0", ``},
		{int8(123), "123", ``},
		{int8(-123), "-123", ``},
	})
}

func TestInt8Int(t *testing.T) {
	Cases(t, Int8Int, []Case{
		{int8(0), 0, ``},
		{int8(math.MaxInt8), int(math.MaxInt8), ``},
		{int8(math.MinInt8), int(math.MinInt8), ``},
	})
}

func TestInt8Int16(t *testing.T) {
	Cases(t, Int8Int16, []Case{
		{int8(0), int16(0), ``},
		{int8(math.MaxInt8), int16(math.MaxInt8), ``},
		{int8(math.MinInt8), int16(math.MinInt8), ``},
	})
}

func TestInt8Int32(t *testing.T) {
	Cases(t, Int8Int32, []Case{
		{int8(0), int32(0), ``},
		{int8(math.MaxInt8), int32(math.MaxInt8), ``},
		{int8(math.MinInt8), int32(math.MinInt8), ``},
	})
}

func TestInt8Int64(t *testing.T) {
	Cases(t, Int8Int64, []Case{
		{int8(0), int64(0), ``},
		{int8(math.MaxInt8), int64(math.MaxInt8), ``},
		{int8(math.MinInt8), int64(math.MinInt8), ``},
	})
}

func TestInt8Uint(t *testing.T) {
	Cases(t, Int8Uint, []Case{
		{int8(0), uint(0), ``},
		{int8(math.MaxInt8), uint(math.MaxInt8), ``},
		{int8(-123), uint(0), `Int8Uint: input out of range -123`},
	})
}

func TestInt8Uint8(t *testing.T) {
	Cases(t, Int8Uint8, []Case{
		{int8(0), uint8(0), ``},
		{int8(math.MaxInt8), uint8(math.MaxInt8), ``},
		{int8(-123), uint8(0), `Int8Uint8: input out of range -123`},
	})
}

func TestInt8Uint16(t *testing.T) {
	Cases(t, Int8Uint16, []Case{
		{int8(0), uint16(0), ``},
		{int8(math.MaxInt8), uint16(math.MaxInt8), ``},
		{int8(-123), uint16(0), `Int8Uint16: input out of range -123`},
	})
}

func TestInt8Uint32(t *testing.T) {
	Cases(t, Int8Uint32, []Case{
		{int8(0), uint32(0), ``},
		{int8(math.MaxInt8), uint32(math.MaxInt8), ``},
		{int8(-123), uint32(0), `Int8Uint32: input out of range -123`},
	})
}

func TestInt8Uint64(t *testing.T) {
	Cases(t, Int8Uint64, []Case{
		{int8(0), uint64(0), ``},
		{int8(math.MaxInt8), uint64(math.MaxInt8), ``},
		{int8(-123), uint64(0), `Int8Uint64: input out of range -123`},
	})
}

func TestInt8Float32(t *testing.T) {
	Cases(t, Int8Float32, []Case{
		{int8(0), float32(0), ``},
		{int8(math.MaxInt8), float32(math.MaxInt8), ``},
		{int8(-123), float32(-123), ``},
	})
}

func TestInt8Float64(t *testing.T) {
	Cases(t, Int8Float64, []Case{
		{int8(0), float64(0), ``},
		{int8(math.MaxInt8), float64(math.MaxInt8), ``},
		{int8(-123), float64(-123), ``},
	})
}

func TestInt16Bool(t *testing.T) {
	Cases(t, Int16Bool, []Case{
		{int16(0), false, ``},
		{int16(1), true, ``},
		{int16(2), false, `Int16Bool: invalid input 2`},
	})
}

func TestInt16String(t *testing.T) {
	Cases(t, Int16String, []Case{
		{int16(0), "0", ``},
		{int16(123), "123", ``},
		{int16(-456), "-456", ``},
	})
}
func TestInt16Int(t *testing.T) {
	cases := []Case{
		{int16(0), 0, ""},
		{int16(1), 1, ""},
		{int16(-1), -1, ""},
	}
	Cases(t, Int16Int, cases)
}

func TestInt16Int8(t *testing.T) {
	cases := []Case{
		{int16(0), int8(0), ""},
		{int16(1), int8(1), ""},
		{int16(math.MaxInt8), int8(math.MaxInt8), ""},
		{int16(math.MinInt8), int8(math.MinInt8), ""},
		{int16(math.MaxInt8 + 1), int8(0), "Int16Int8: input out of range 128"},
		{int16(math.MinInt8 - 1), int8(0), "Int16Int8: input out of range -129"},
	}
	Cases(t, Int16Int8, cases)
}

func TestInt16Int32(t *testing.T) {
	cases := []Case{
		{int16(0), int32(0), ""},
		{int16(1), int32(1), ""},
		{int16(-1), int32(-1), ""},
	}
	Cases(t, Int16Int32, cases)
}

func TestInt16Int64(t *testing.T) {
	cases := []Case{
		{int16(0), int64(0), ""},
		{int16(1), int64(1), ""},
		{int16(-1), int64(-1), ""},
	}
	Cases(t, Int16Int64, cases)
}

func TestInt16Uint(t *testing.T) {
	cases := []Case{
		{int16(0), uint(0), ""},
		{int16(1), uint(1), ""},
		{int16(-1), uint(0), "Int16Uint: input out of range -1"},
	}
	Cases(t, Int16Uint, cases)
}

func TestInt16Uint8(t *testing.T) {
	cases := []Case{
		{int16(0), uint8(0), ""},
		{int16(1), uint8(1), ""},
		{int16(math.MaxUint8), uint8(math.MaxUint8), ""},
		{int16(-1), uint8(0), "Int16Uint8: input out of range -1"},
	}
	Cases(t, Int16Uint8, cases)
}

func TestInt16Uint16(t *testing.T) {
	cases := []Case{
		{int16(0), uint16(0), ""},
		{int16(1), uint16(1), ""},
		{int16(-1), uint16(0), "Int16Uint16: input out of range -1"},
	}
	Cases(t, Int16Uint16, cases)
}

func TestInt16Uint32(t *testing.T) {
	cases := []Case{
		{int16(0), uint32(0), ""},
		{int16(1), uint32(1), ""},
		{int16(-1), uint32(0), "Int16Uint32: input out of range -1"},
	}
	Cases(t, Int16Uint32, cases)
}

func TestInt16Uint64(t *testing.T) {
	cases := []Case{
		{int16(0), uint64(0), ""},
		{int16(1), uint64(1), ""},
		{int16(-1), uint64(0), "Int16Uint64: input out of range -1"},
	}
	Cases(t, Int16Uint64, cases)
}

func TestInt16Float32(t *testing.T) {
	cases := []Case{
		{int16(0), float32(0), ""},
		{int16(1), float32(1), ""},
		{int16(-1), float32(-1), ""},
	}
	Cases(t, Int16Float32, cases)
}

func TestInt16Float64(t *testing.T) {
	cases := []Case{
		{int16(0), float64(0), ""},
		{int16(1), float64(1), ""},
		{int16(-1), float64(-1), ""},
	}
	Cases(t, Int16Float64, cases)
}

func TestInt32Bool(t *testing.T) {
	cases := []Case{
		{int32(0), false, ""},
		{int32(1), true, ""},
		{int32(2), false, "Int32Bool: invalid input 2"},
	}
	Cases(t, Int32Bool, cases)
}

func TestInt32String(t *testing.T) {
	cases := []Case{
		{int32(0), "0", ""},
		{int32(1), "1", ""},
		{int32(-1), "-1", ""},
	}
	Cases(t, Int32String, cases)
}

func TestInt32Int(t *testing.T) {
	cases := []Case{
		{int32(0), 0, ""},
		{int32(1), 1, ""},
		{int32(-1), -1, ""},
	}
	Cases(t, Int32Int, cases)
}

func TestInt32Int8(t *testing.T) {
	cases := []Case{
		{int32(0), int8(0), ""},
		{int32(1), int8(1), ""},
		{int32(math.MaxInt8), int8(math.MaxInt8), ""},
		{int32(math.MinInt8), int8(math.MinInt8), ""},
		{int32(math.MaxInt8 + 1), int8(0), "Int32Int8: input out of range 128"},
		{int32(math.MinInt8 - 1), int8(0), "Int32Int8: input out of range -129"},
	}
	Cases(t, Int32Int8, cases)
}

func TestInt32Int16(t *testing.T) {
	Cases(t, Int32Int16, []Case{
		{int32(math.MinInt32), int16(0), `Int32Int16: input out of range -2147483648`},
		{int32(math.MaxInt32), int16(0), `Int32Int16: input out of range 2147483647`},
	})
}

func TestInt32Int64(t *testing.T) {
	Cases(t, Int32Int64, []Case{
		{int32(math.MinInt32), int64(math.MinInt32), ``},
		{int32(math.MaxInt32), int64(math.MaxInt32), ``},
	})
}

func TestInt32Uint(t *testing.T) {
	Cases(t, Int32Uint, []Case{
		{int32(0), uint(0), ``},
		{int32(math.MaxInt32), uint(math.MaxInt32), ``},
		{int32(-1), uint(0), `Int32Uint: input out of range -1`},
	})
}

func TestInt32Uint8(t *testing.T) {
	Cases(t, Int32Uint8, []Case{
		{int32(0), uint8(0), ``},
		{int32(math.MaxInt32), uint8(0), `Int32Uint8: input out of range 2147483647`},
		{int32(math.MaxUint8), uint8(math.MaxUint8), ``},
		{int32(-1), uint8(0), `Int32Uint8: input out of range -1`},
	})
}

func TestInt32Uint16(t *testing.T) {
	Cases(t, Int32Uint16, []Case{
		{int32(0), uint16(0), ``},
		{int32(math.MaxInt32), uint16(0), `Int32Uint16: input out of range 2147483647`},
		{int32(math.MaxUint16), uint16(math.MaxUint16), ``},
		{int32(-1), uint16(0), `Int32Uint16: input out of range -1`},
	})
}

func TestInt32Uint32(t *testing.T) {
	Cases(t, Int32Uint32, []Case{
		{int32(0), uint32(0), ``},
		{int32(math.MaxInt32), uint32(math.MaxInt32), ``},
		{int32(-1), uint32(0), `Int32Uint32: input out of range -1`},
	})
}

func TestInt32Uint64(t *testing.T) {
	Cases(t, Int32Uint64, []Case{
		{int32(0), uint64(0), ``},
		{int32(math.MaxInt32), uint64(math.MaxInt32), ``},
		{int32(-1), uint64(0), `Int32Uint64: input out of range -1`},
	})
}

func TestInt32Float32(t *testing.T) {
	Cases(t, Int32Float32, []Case{
		{int32(0), float32(0), ``},
		{int32(math.MaxInt32), float32(math.MaxInt32), ``},
	})
}

func TestInt32Float64(t *testing.T) {
	Cases(t, Int32Float64, []Case{
		{int32(0), float64(0), ``},
		{int32(math.MaxInt32), float64(math.MaxInt32), ``},
	})
}
