package xconv

import (
	"fmt"
	xerr "github.com/goclub/error"
	"math"
	"strconv"
)

func Int64PointerInt64(int642Pointer *int64) int64 {
	if int642Pointer == nil {
		return 0
	} else {
		return *int642Pointer
	}
}

func IntBool(i int) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("IntBool: invalid input %d", i))
	}
}

func IntString(i int) string {
	return strconv.Itoa(i)
}

func IntInt8(i int) (int8, error) {
	if i < math.MinInt8 || i > math.MaxInt8 {
		return 0, xerr.New(fmt.Sprintf("IntInt8: input out of range %d", i))
	}
	return int8(i), nil
}

func IntInt16(i int) (int16, error) {
	if i < math.MinInt16 || i > math.MaxInt16 {
		return 0, xerr.New(fmt.Sprintf("IntInt16: input out of range %d", i))
	}
	return int16(i), nil
}

func IntInt32(i int) (int32, error) {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0, xerr.New(fmt.Sprintf("IntInt32: input out of range %d", i))
	}
	return int32(i), nil
}

func IntInt64(i int) (int64, error) {
	return int64(i), nil
}

func IntUint(i int) (uint, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("IntUint: input out of range %d", i))
	}
	return uint(i), nil
}

func IntUint8(i int) (uint8, error) {
	if i < 0 || i > math.MaxUint8 {
		return 0, xerr.New(fmt.Sprintf("IntUint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func IntUint16(i int) (uint16, error) {
	if i < 0 || i > math.MaxUint16 {
		return 0, xerr.New(fmt.Sprintf("IntUint16: input out of range %d", i))
	}
	return uint16(i), nil
}

func IntUint32(i int) (uint32, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("IntUint32: input out of range %d", i))
	}
	return uint32(i), nil
}

func IntUint64(i int) (uint64, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("IntUint64: input out of range %d", i))
	}
	return uint64(i), nil
}

func IntFloat32(i int) (float32, error) {
	return float32(i), nil
}

func IntFloat64(i int) (float64, error) {
	return float64(i), nil
}

func Int8Bool(i int8) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("Int8Bool: invalid input %d", i))
	}
}

func Int8String(i int8) (string, error) {
	return strconv.Itoa(int(i)), nil
}

func Int8Int(i int8) (int, error) {
	return int(i), nil
}

func Int8Int16(i int8) (int16, error) {
	return int16(i), nil
}

func Int8Int32(i int8) (int32, error) {
	return int32(i), nil
}

func Int8Int64(i int8) (int64, error) {
	return int64(i), nil
}

func Int8Uint(i int8) (uint, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int8Uint: input out of range %d", i))
	}
	return uint(i), nil
}

func Int8Uint8(i int8) (uint8, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int8Uint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func Int8Uint16(i int8) (uint16, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int8Uint16: input out of range %d", i))
	}
	return uint16(i), nil
}

func Int8Uint32(i int8) (uint32, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int8Uint32: input out of range %d", i))
	}
	return uint32(i), nil
}

func Int8Uint64(i int8) (uint64, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int8Uint64: input out of range %d", i))
	}
	return uint64(i), nil
}

func Int8Float32(i int8) (float32, error) {
	return float32(i), nil
}

func Int8Float64(i int8) (float64, error) {
	return float64(i), nil
}

func Int16Bool(i int16) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("Int16Bool: invalid input %d", i))
	}
}

func Int16String(i int16) (string, error) {
	return strconv.Itoa(int(i)), nil
}

func Int16Int(i int16) (int, error) {
	return int(i), nil
}

func Int16Int8(i int16) (int8, error) {
	if i < math.MinInt8 || i > math.MaxInt8 {
		return 0, xerr.New(fmt.Sprintf("Int16Int8: input out of range %d", i))
	}
	return int8(i), nil
}

func Int16Int32(i int16) (int32, error) {
	return int32(i), nil
}

func Int16Int64(i int16) (int64, error) {
	return int64(i), nil
}

func Int16Uint(i int16) (uint, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int16Uint: input out of range %d", i))
	}
	return uint(i), nil
}

func Int16Uint8(i int16) (uint8, error) {
	if i < 0 || i > math.MaxUint8 {
		return 0, xerr.New(fmt.Sprintf("Int16Uint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func Int16Uint16(i int16) (uint16, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int16Uint16: input out of range %d", i))
	}
	return uint16(i), nil
}

func Int16Uint32(i int16) (uint32, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int16Uint32: input out of range %d", i))
	}
	return uint32(i), nil
}

func Int16Uint64(i int16) (uint64, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int16Uint64: input out of range %d", i))
	}
	return uint64(i), nil
}

func Int16Float32(i int16) (float32, error) {
	return float32(i), nil
}

func Int16Float64(i int16) (float64, error) {
	return float64(i), nil
}

func Int32Bool(i int32) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("Int32Bool: invalid input %d", i))
	}
}

func Int32String(i int32) (string, error) {
	return strconv.Itoa(int(i)), nil
}

func Int32Int(i int32) (int, error) {
	return int(i), nil
}

func Int32Int8(i int32) (int8, error) {
	if i < math.MinInt8 || i > math.MaxInt8 {
		return 0, xerr.New(fmt.Sprintf("Int32Int8: input out of range %d", i))
	}
	return int8(i), nil
}

func Int32Int16(i int32) (int16, error) {
	if i < math.MinInt16 || i > math.MaxInt16 {
		return 0, xerr.New(fmt.Sprintf("Int32Int16: input out of range %d", i))
	}
	return int16(i), nil
}

func Int32Int64(i int32) (int64, error) {
	return int64(i), nil
}

func Int32Uint(i int32) (uint, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int32Uint: input out of range %d", i))
	}
	return uint(i), nil
}

func Int32Uint8(i int32) (uint8, error) {
	if i < 0 || i > math.MaxUint8 {
		return 0, xerr.New(fmt.Sprintf("Int32Uint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func Int32Uint16(i int32) (uint16, error) {
	if i < 0 || i > math.MaxUint16 {
		return 0, xerr.New(fmt.Sprintf("Int32Uint16: input out of range %d", i))
	}
	return uint16(i), nil
}

func Int32Uint32(i int32) (uint32, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int32Uint32: input out of range %d", i))
	}
	return uint32(i), nil
}

func Int32Uint64(i int32) (uint64, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int32Uint64: input out of range %d", i))
	}
	return uint64(i), nil
}

func Int32Float32(i int32) (float32, error) {
	return float32(i), nil
}

func Int32Float64(i int32) (float64, error) {
	return float64(i), nil
}

func Int64Bool(i int64) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("Int64Bool: invalid input %d", i))
	}
}

func Int64String(i int64) (string, error) {
	return strconv.FormatInt(i, 10), nil
}

func Int64Int(i int64) (int, error) {
	if i > math.MaxInt32 || i < math.MinInt32 {
		return 0, xerr.New(fmt.Sprintf("Int64Int: input out of range %d", i))
	}
	return int(i), nil
}

func Int64Int8(i int64) (int8, error) {
	if i < math.MinInt8 || i > math.MaxInt8 {
		return 0, xerr.New(fmt.Sprintf("Int64Int8: input out of range %d", i))
	}
	return int8(i), nil
}

func Int64Int16(i int64) (int16, error) {
	if i < math.MinInt16 || i > math.MaxInt16 {
		return 0, xerr.New(fmt.Sprintf("Int64Int16: input out of range %d", i))
	}
	return int16(i), nil
}

func Int64Int32(i int64) (int32, error) {
	if i < math.MinInt32 || i > math.MaxInt32 {
		return 0, xerr.New(fmt.Sprintf("Int64Int32: input out of range %d", i))
	}
	return int32(i), nil
}

func Int64Uint(i int64) (uint, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int64Uint: input out of range %d", i))
	}
	return uint(i), nil
}

func Int64Uint8(i int64) (uint8, error) {
	if i < 0 || i > math.MaxUint8 {
		return 0, xerr.New(fmt.Sprintf("Int64Uint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func Int64Uint16(i int64) (uint16, error) {
	if i < 0 || i > math.MaxUint16 {
		return 0, xerr.New(fmt.Sprintf("Int64Uint16: input out of range %d", i))
	}
	return uint16(i), nil
}

func Int64Uint32(i int64) (uint32, error) {
	if i < 0 || i > math.MaxUint32 {
		return 0, xerr.New(fmt.Sprintf("Int64Uint32: input out of range %d", i))
	}
	return uint32(i), nil
}

func Int64Uint64(i int64) (uint64, error) {
	if i < 0 {
		return 0, xerr.New(fmt.Sprintf("Int64Uint64: input out of range %d", i))
	}
	return uint64(i), nil
}

func Int64Float32(i int64) (float32, error) {
	if float64(i) > math.MaxFloat32 || float64(i) < -math.MaxFloat32 {
		return 0, xerr.New(fmt.Sprintf("Int64Float32: input out of range %d", i))
	}
	return float32(i), nil
}

func Int64Float64(i int64) (float64, error) {
	return float64(i), nil
}

func UintBool(i uint) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("UintBool: invalid input %d", i))
	}
}

func UintString(i uint) (string, error) {
	return strconv.Itoa(int(i)), nil
}

func UintInt(i uint) (int, error) {
	if i > uint(math.MaxInt32) {
		return 0, xerr.New(fmt.Sprintf("UintInt: input out of range %d", i))
	}
	return int(i), nil
}

func UintInt8(i uint) (int8, error) {
	if i > uint(math.MaxInt8) {
		return 0, xerr.New(fmt.Sprintf("UintInt8: input out of range %d", i))
	}
	return int8(i), nil
}

func UintInt16(i uint) (int16, error) {
	if i > uint(math.MaxInt16) {
		return 0, xerr.New(fmt.Sprintf("UintInt16: input out of range %d", i))
	}
	return int16(i), nil
}

func UintInt32(i uint) (int32, error) {
	if i > uint(math.MaxInt32) {
		return 0, xerr.New(fmt.Sprintf("UintInt32: input out of range %d", i))
	}
	return int32(i), nil
}

func UintInt64(i uint) (int64, error) {
	return int64(i), nil
}

func UintUint8(i uint) (uint8, error) {
	if i > math.MaxUint8 {
		return 0, xerr.New(fmt.Sprintf("UintUint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func UintUint16(i uint) (uint16, error) {
	if i > math.MaxUint16 {
		return 0, xerr.New(fmt.Sprintf("UintUint16: input out of range %d", i))
	}
	return uint16(i), nil
}

func UintUint32(i uint) (uint32, error) {
	if i > math.MaxUint32 {
		return 0, xerr.New(fmt.Sprintf("UintUint32: input out of range %d", i))
	}
	return uint32(i), nil
}

func UintUint64(i uint) (uint64, error) {
	return uint64(i), nil
}

func UintFloat32(i uint) (float32, error) {
	return float32(i), nil
}

func UintFloat64(i uint) (float64, error) {
	return float64(i), nil
}

func Uint8Bool(i uint8) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("Uint8Bool: invalid input %d", i))
	}
}

func Uint8String(i uint8) (string, error) {
	return strconv.Itoa(int(i)), nil
}

func Uint8Int(i uint8) (int, error) {
	return int(i), nil
}

func Uint8Int8(i uint8) (int8, error) {
	return int8(i), nil
}

func Uint8Int16(i uint8) (int16, error) {
	return int16(i), nil
}

func Uint8Int32(i uint8) (int32, error) {
	return int32(i), nil
}

func Uint8Int64(i uint8) (int64, error) {
	return int64(i), nil
}

func Uint8Uint(i uint8) (uint, error) {
	return uint(i), nil
}

func Uint8Uint16(i uint8) (uint16, error) {
	return uint16(i), nil
}

func Uint8Uint32(i uint8) (uint32, error) {
	return uint32(i), nil
}

func Uint8Uint64(i uint8) (uint64, error) {
	return uint64(i), nil
}

func Uint8Float32(i uint8) (float32, error) {
	return float32(i), nil
}

func Uint8Float64(i uint8) (float64, error) {
	return float64(i), nil
}
func Uint16Bool(i uint16) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("Uint16Bool: invalid input %d", i))
	}
}

func Uint16String(i uint16) (string, error) {
	return strconv.Itoa(int(i)), nil
}

func Uint16Int(i uint16) (int, error) {
	return int(i), nil
}

func Uint16Int8(i uint16) (int8, error) {
	if i > uint16(math.MaxInt8) {
		return 0, xerr.New(fmt.Sprintf("Uint16Int8: input out of range %d", i))
	}
	return int8(i), nil
}

func Uint16Int16(i uint16) (int16, error) {
	return int16(i), nil
}

func Uint16Int32(i uint16) (int32, error) {
	return int32(i), nil
}

func Uint16Int64(i uint16) (int64, error) {
	return int64(i), nil
}

func Uint16Uint(i uint16) (uint, error) {
	return uint(i), nil
}

func Uint16Uint8(i uint16) (uint8, error) {
	if i > math.MaxUint8 {
		return 0, xerr.New(fmt.Sprintf("Uint16Uint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func Uint16Uint32(i uint16) (uint32, error) {
	return uint32(i), nil
}

func Uint16Uint64(i uint16) (uint64, error) {
	return uint64(i), nil
}

func Uint16Float32(i uint16) (float32, error) {
	return float32(i), nil
}

func Uint16Float64(i uint16) (float64, error) {
	return float64(i), nil
}

func Uint32Bool(i uint32) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("Uint32Bool: invalid input %d", i))
	}
}

func Uint32String(i uint32) (string, error) {
	return strconv.Itoa(int(i)), nil
}

func Uint32Int(i uint32) (int, error) {
	if i > uint32(math.MaxInt32) {
		return 0, xerr.New(fmt.Sprintf("Uint32Int: input out of range %d", i))
	}
	return int(i), nil
}

func Uint32Int8(i uint32) (int8, error) {
	if i > uint32(math.MaxInt8) {
		return 0, xerr.New(fmt.Sprintf("Uint32Int8: input out of range %d", i))
	}
	return int8(i), nil
}

func Uint32Int16(i uint32) (int16, error) {
	if i > uint32(math.MaxInt16) {
		return 0, xerr.New(fmt.Sprintf("Uint32Int16: input out of range %d", i))
	}
	return int16(i), nil
}

func Uint32Int32(i uint32) (int32, error) {
	return int32(i), nil
}

func Uint32Int64(i uint32) (int64, error) {
	return int64(i), nil
}

func Uint32Uint(i uint32) (uint, error) {
	return uint(i), nil
}

func Uint32Uint8(i uint32) (uint8, error) {
	if i > math.MaxUint8 {
		return 0, xerr.New(fmt.Sprintf("Uint32Uint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func Uint32Uint16(i uint32) (uint16, error) {
	if i > math.MaxUint16 {
		return 0, xerr.New(fmt.Sprintf("Uint32Uint16: input out of range %d", i))
	}
	return uint16(i), nil
}

func Uint32Uint64(i uint32) (uint64, error) {
	return uint64(i), nil
}

func Uint32Float32(i uint32) (float32, error) {
	return float32(i), nil
}

func Uint32Float64(i uint32) (float64, error) {
	return float64(i), nil
}

func Uint64Bool(i uint64) (bool, error) {
	if i == 0 {
		return false, nil
	} else if i == 1 {
		return true, nil
	} else {
		return false, xerr.New(fmt.Sprintf("Uint64Bool: invalid input %d", i))
	}
}

func Uint64String(i uint64) (string, error) {
	return strconv.FormatUint(i, 10), nil
}

func Uint64Int(i uint64) (int, error) {
	if i > uint64(math.MaxInt32) {
		return 0, xerr.New(fmt.Sprintf("Uint64Int: input out of range %d", i))
	}
	return int(i), nil
}

func Uint64Int8(i uint64) (int8, error) {
	if i > uint64(math.MaxInt8) {
		return 0, xerr.New(fmt.Sprintf("Uint64Int8: input out of range %d", i))
	}
	return int8(i), nil
}

func Uint64Int16(i uint64) (int16, error) {
	if i > uint64(math.MaxInt16) {
		return 0, xerr.New(fmt.Sprintf("Uint64Int16: input out of range %d", i))
	}
	return int16(i), nil
}

func Uint64Int32(i uint64) (int32, error) {
	if i > uint64(math.MaxInt32) {
		return 0, xerr.New(fmt.Sprintf("Uint64Int32: input out of range %d", i))
	}
	return int32(i), nil
}

func Uint64Int64(i uint64) (int64, error) {
	return int64(i), nil
}

func Uint64Uint(i uint64) (uint, error) {
	return uint(i), nil
}

func Uint64Uint8(i uint64) (uint8, error) {
	if i > math.MaxUint8 {
		return 0, xerr.New(fmt.Sprintf("Uint64Uint8: input out of range %d", i))
	}
	return uint8(i), nil
}

func Uint64Uint16(i uint64) (uint16, error) {
	if i > math.MaxUint16 {
		return 0, xerr.New(fmt.Sprintf("Uint64Uint16: input out of range %d", i))
	}
	return uint16(i), nil
}

func Uint64Uint32(i uint64) (uint32, error) {
	if i > math.MaxUint32 {
		return 0, xerr.New(fmt.Sprintf("Uint64Uint32: input out of range %d", i))
	}
	return uint32(i), nil
}

func Uint64Float32(i uint64) (float32, error) {
	return float32(i), nil
}

func Uint64Float64(i uint64) (float64, error) {
	return float64(i), nil
}
