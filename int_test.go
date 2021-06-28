package xconv

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestIntBool(t *testing.T) {
	assert.Equal(t, true, IntBool(int(1)))
	assert.Equal(t, false, IntBool(int(0)))
}

func TestIntString(t *testing.T) {
	assert.Equal(t, "123456", IntString(int(123456)))
	assert.Equal(t, "123456", Int32String(int32(123456)))
	assert.Equal(t, "123456", Int64String(int64(123456)))
	assert.Equal(t, "11110001001000000", Int64StringWithBase(int64(123456), 2))
}
func TestInt8Int16(t *testing.T) {
	assert.Equal(t, Int8Int16(int8(math.MinInt8)), int16(-128))
}

func TestInt16Int8(t *testing.T) {
	{
		v, err := Int16Int8(int16(math.MinInt8-1))
		assert.EqualError(t, err, "goclub/conv: int16 to int8 fail, int must greater than or equal to math.MinInt8")
		assert.Equal(t, v, int8(0))
	}
	{
		v, err := Int16Int8(int16(math.MinInt8))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(-128))
	}
	{
		v, err := Int16Int8(int16(math.MinInt8+1))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(-127))
	}
	{
		v, err := Int16Int8(int16(math.MaxInt8))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(127))
	}
	{
		v, err := Int16Int8(int16(math.MaxInt8-1))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(126))
	}
	{
		v, err := Int16Int8(int16(math.MaxInt8+1))
		assert.EqualError(t, err, "goclub/conv: int16 to int8 fail, int must less than or equal to math.MaxInt8")
		assert.Equal(t, v, int8(0))
	}
}

func TestInt8Int32(t *testing.T) {
	assert.Equal(t, Int8Int32(int8(math.MinInt8)), int32(-128))
}
func TestInt32Int8(t *testing.T) {
	{
		v, err := Int32Int8(int32(math.MinInt8-1))
		assert.EqualError(t, err, "goclub/conv: int32 to int8 fail, int must greater than or equal to math.MinInt8")
		assert.Equal(t, v, int8(0))
	}
	{
		v, err := Int32Int8(int32(math.MinInt8))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(-128))
	}
	{
		v, err := Int32Int8(int32(math.MinInt8+1))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(-127))
	}
	{
		v, err := Int32Int8(int32(math.MaxInt8))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(127))
	}
	{
		v, err := Int32Int8(int32(math.MaxInt8-1))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(126))
	}
	{
		v, err := Int32Int8(int32(math.MaxInt8+1))
		assert.EqualError(t, err, "goclub/conv: int32 to int8 fail, int must less than or equal to math.MaxInt8")
		assert.Equal(t, v, int8(0))
	}
}
func TestInt8Int64(t *testing.T) {
	assert.Equal(t, Int8Int64(int8(math.MinInt8)), int64(-128))
}
func TestInt64Int8(t *testing.T) {
	{
		v, err := Int64Int8(int64(math.MinInt8-1))
		assert.EqualError(t, err, "goclub/conv: int64 to int8 fail, int must greater than or equal to math.MinInt8")
		assert.Equal(t, v, int8(0))
	}
	{
		v, err := Int64Int8(int64(math.MinInt8))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(-128))
	}
	{
		v, err := Int64Int8(int64(math.MinInt8+1))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(-127))
	}
	{
		v, err := Int64Int8(int64(math.MaxInt8))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(127))
	}
	{
		v, err := Int64Int8(int64(math.MaxInt8-1))
		assert.NoError(t, err)
		assert.Equal(t, v, int8(126))
	}
	{
		v, err := Int64Int8(int64(math.MaxInt8+1))
		assert.EqualError(t, err, "goclub/conv: int64 to int8 fail, int must less than or equal to math.MaxInt8")
		assert.Equal(t, v, int8(0))
	}
}
func TestInt32Int64(t *testing.T) {
	assert.Equal(t, Int32Int64(math.MinInt32), int64(math.MinInt32))
	assert.Equal(t, Int32Int64(math.MaxInt32), int64(math.MaxInt32))
}

func TestInt64Int32(t *testing.T) {
	{
		i32, err := Int64Int32(math.MinInt64)
		assert.EqualError(t, err, "goclub/conv: int64 to int32 fail, int must greater than or equal to math.MinInt32")
		assert.Equal(t, i32, int32(0))
	}
	{
		i32, err := Int64Int32(math.MinInt32)
		assert.NoError(t, err)
		assert.Equal(t, i32, int32(-2147483648))
	}
	{
		i32, err := Int64Int32(math.MinInt32-1)
		assert.EqualError(t, err, "goclub/conv: int64 to int32 fail, int must greater than or equal to math.MinInt32")
		assert.Equal(t, i32, int32(0))
	}
	{
		i32, err := Int64Int32(math.MaxInt32)
		assert.NoError(t, err)
		assert.Equal(t, i32, int32(2147483647))
	}
	{
		i32, err := Int64Int32(math.MaxInt32+1)
		assert.EqualError(t, err, "goclub/conv: int64 to int32 fail, int must less than or equal to math.MaxInt32")
		assert.Equal(t, i32, int32(0))
	}
}

func TestUint32Uint64(t *testing.T) {
	assert.Equal(t, Uint32Uint64(0), uint64(0))
	assert.Equal(t, Uint32Uint64(math.MaxUint32), uint64(math.MaxUint32))
}

func TestUint64Uint32(t *testing.T) {
	{
		i32, err := Uint64Uint32(math.MaxUint32)
		assert.NoError(t, err)
		assert.Equal(t, i32, uint32(4294967295))
	}
	{
		i32, err := Uint64Uint32(math.MaxUint32+1)
		assert.EqualError(t, err, "goclub/conv: uint64 to uint32 fail, int must less than or equal to math.MaxUint32")
		assert.Equal(t, i32, uint32(0))
	}
}
