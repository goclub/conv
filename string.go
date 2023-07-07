package xconv

import (
	"strconv"

	xerr "github.com/goclub/error"
)

func StringInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return i, nil
}

func StringInt8(s string) (int8, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return int8(i), nil
}

func StringInt16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return int16(i), nil
}

func StringInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return int32(i), nil
}

func StringInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return i, nil
}

func StringUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return uint(i), nil
}

func StringUint8(s string) (uint8, error) {
	i, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return uint8(i), nil
}

func StringUint16(s string) (uint16, error) {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return uint16(i), nil
}

func StringUint32(s string) (uint32, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return uint32(i), nil
}

func StringUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return i, nil
}

func StringFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return float32(f), nil
}

func StringFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, xerr.WithStack(err)
	}
	return f, nil
}

func StringBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, xerr.WithStack(err)
	}
	return b, nil
}

func StringBytes(s string) []byte {
	return []byte(s)
}
func StringPointerString(p *string) (s string) {
	if p == nil {
		return ""
	}
	return *p
}
