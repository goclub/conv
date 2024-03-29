package xconv

import (
	"errors"
	"reflect"
	"strconv"
)


func StringInt(s string) (i int, err error) {
	return strconv.Atoi(s)
}
func StringInt64 (s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
func StringFloat64 (s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
func StringFloat32 (s string) (float32, error) {
	f64, err := strconv.ParseFloat(s, 32)
	return float32(f64), err
}
func StringBool(s string) (bool, error) {
	switch s {
	case "true",
		"True",
		"t",
		"T",
		"1":
		return true, nil
	case "false",
		"False",
		"f",
		"F",
		"0":
		return false, nil
	}
	return false, errors.New("goclub/conv: " + s + " can't conv to bool")
}
func ReflectToString(rValue reflect.Value) (value string, err error) {
	rType := rValue.Type()
	switch rType.Kind() {
	case reflect.String:
		return rValue.String(), nil
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		return strconv.FormatInt(rValue.Int(), 10), nil
	case reflect.Uint,reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		return strconv.FormatUint(rValue.Uint(), 10), nil
	case reflect.Bool:
		// 采用1 0 尽可能的让数据量小
		if rValue.Bool() {
			return "1", nil
		} else {
			return "0", nil
		}
	case reflect.Float32:
		return strconv.FormatFloat(rValue.Float(),'g', -1, 32), nil
	case reflect.Float64:
		return strconv.FormatFloat(rValue.Float(),'g', -1, 64), nil
	case reflect.Array, reflect.Slice:
		elemType := rType.Elem()
		switch elemType.Kind() {
		case reflect.Uint8: // []byte
			bytes := rValue.Bytes()
			return string(bytes), nil
		default:
			return "", errors.New("goclub/conv: field(name:" + rType.Name() + "kind:" +  rType.Kind().String() + ")" + "can not conv to string " + rType.String())
		}
	default:
		return "", errors.New("goclub/conv: field(name:" + rType.Name() + "kind:" +  rType.Kind().String() + ")" + "can not conv to string " + rType.String())
	}
	return
}
func StringToReflect(s string, rValue reflect.Value)  error {
	rType := rValue.Type()
	if !rValue.CanSet() {
		if rValue.Kind() == reflect.Ptr {
			rValue = rValue.Elem()
			rType = rType.Elem()
		}
		if !rValue.CanSet() {
			return errors.New("goclub/conv: StringReflect(s, rValue) rValue must can set, mu be you should use reflect.ValueOf(pointer)")
		}
	}
	return coreStringToReflect(s, rValue, rType)
}
func coreStringToReflect (s string, rValue reflect.Value, rType reflect.Type) error {
	switch rType.Kind() {
	case reflect.String:
		rValue.SetString(s)
	case reflect.Int:
		i, err := StringInt64(s) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Int8:
		i, err := strconv.ParseInt(s, 10, 8) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Int16:
		i, err := strconv.ParseInt(s, 10, 16) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Int32:
		i, err := strconv.ParseInt(s, 10, 32) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64) ; if err != nil { return err}
		rValue.SetInt(i)
	case reflect.Uint:
		i, err := strconv.ParseUint(s, 10, 64) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Uint8:
		i, err := strconv.ParseUint(s, 10, 8) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Uint16:
		i, err := strconv.ParseUint(s, 10, 16) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Uint32:
		i, err := strconv.ParseUint(s, 10, 32) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Uint64:
		i, err := strconv.ParseUint(s, 10, 64) ; if err != nil { return err}
		rValue.SetUint(i)
	case reflect.Bool:
		b, err := StringBool(s) ; if err != nil {panic(err)}
		rValue.SetBool(b)
	case reflect.Float32:
		f, err := StringFloat32(s) ; if err != nil {panic(err)}
		rValue.SetFloat(float64(f))
	case reflect.Float64:
		f, err := StringFloat64(s) ; if err != nil {panic(err)}
		rValue.SetFloat(f)
	case reflect.Array, reflect.Slice:
		elemType := rType.Elem()
		switch elemType.Kind() {
		case reflect.Uint8: // []byte
			bytes := []byte(s)
			rValue.SetBytes(bytes)
		default:
			return errors.New("goclub/conv: field(name:" + rType.Name() + "kind:" +  rType.Kind().String() + ")" + "can not string conv type " + rType.String())
		}
	default:
		return errors.New("goclub/conv: field(name:" + rType.Name() + "kind:" +  rType.Kind().String() + ")" + "can not string conv type " + rType.String())
	}
	return nil
}

func StringPointerString(stringPointer *string) string {
	if stringPointer == nil {
		return ""
	} else {
		return *stringPointer
	}
}