package xconv

import (
	"errors"
	"math"
	"strconv"
)

func Int8Int16(i int8) int16 {
	return int16(i)
}
func Int16Int8(i int16) (int8, error) {
	if i < math.MinInt8 {
		return 0, errors.New("goclub/conv: int16 to int8 fail, int must greater than or equal to math.MinInt8")
	}
	if i> math.MaxInt8  {
		return 0, errors.New("goclub/conv: int16 to int8 fail, int must less than or equal to math.MaxInt8")
	}
	return int8(i),nil
}
func Int8Int32(i int8) int32 {
	return int32(i)
}
func Int32Int8(i int32) (int8, error) {
	if i < math.MinInt8 {
		return 0, errors.New("goclub/conv: int32 to int8 fail, int must greater than or equal to math.MinInt8")
	}
	if i> math.MaxInt8  {
		return 0, errors.New("goclub/conv: int32 to int8 fail, int must less than or equal to math.MaxInt8")
	}
	return int8(i),nil
}
func Int8Int64(i int8) int64 {
	return int64(i)
}
func Int64Int8(i int64) (int8, error) {
	if i < math.MinInt8 {
		return 0, errors.New("goclub/conv: int64 to int8 fail, int must greater than or equal to math.MinInt8")
	}
	if i> math.MaxInt8  {
		return 0, errors.New("goclub/conv: int64 to int8 fail, int must less than or equal to math.MaxInt8")
	}
	return int8(i),nil
}

func IntBool (i int) bool{
	if i == 0 { return false }
	return true
}
// int 类型转换为字符串（10进制）
func IntString(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
// int32 类型转换为字符串（10进制）
func Int32String(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}
// int64 类型转换为字符串（10进制）
func Int64String(i int64) string {
	return strconv.FormatInt(i, 10)
}
// int64 类型转换为字符串（自定义进制）2~36
func Int64StringWithBase(i int64, base int) string {
	return strconv.FormatInt(i, base)
}
